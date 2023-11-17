package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/charmbracelet/log"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/pewpewnor/portorico/server/src/handlers"
	"github.com/pewpewnor/portorico/server/src/model"
)

func cleanAllSoftDelete(ctx context.Context, wg *sync.WaitGroup, db *sqlx.DB) {
	defer wg.Done()
	ticker := time.NewTicker(10 * time.Minute)

	for {
		select {
		case <-ticker.C:
			for _, table := range model.Tables {
				if _, err := db.Exec("DELETE FROM $1 WHERE deleted_at IS NOT NULL", table); err != nil {
					log.Errorf("cannot clean soft deleted data in table %v: %v", table, err)
				}
			}
			log.Info("all soft deleted data has been cleaned")
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}

func shutdownServerWhenInterrupt(osChan chan os.Signal, app *fiber.App, db *sqlx.DB, cancel context.CancelFunc, wg *sync.WaitGroup) {
	_ = <-osChan

	err := app.Shutdown()
	if err != nil {
		log.Errorf("error while shutting down fiber: %v", err)
	} else {
		log.Info("server has killed fiber")
	}

	cancel()
	wg.Wait()
	log.Info("server has killed all background routines")

	if err := db.Close(); err != nil {
		log.Errorf("error while closing database connection: %v", err)
	}
	log.Info("server has closed database connection")
}

func main() {
	if err := godotenv.Load(".env.local"); err != nil {
		log.Fatalf("cannot load .env.local file: %v\n", err)
	}

	URI := os.Getenv("DB_URI")
	if URI == "" {
		log.Fatal("environment variable has no 'DB_URI'")
	}
	PORT := os.Getenv("SERVER_PORT")
	if PORT == "" {
		PORT = "8000"
	}

	db, err := sqlx.Connect("postgres", URI)
	if err != nil {
		log.Fatal("sqlx cannot connect to database: %v", err)
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		log.Fatalf("error while creating driver from sqlx for migrator: %v", err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://server/migrations", "postgres", driver)
	if err != nil {
		log.Fatalf("error while creating migrator: %v", err)
	}
	if err := m.Force(1); err != nil {
		log.Fatalf("error while forcing version to 1: %v", err)
	}
	if err := m.Down(); err != nil {
		log.Fatalf("error while migrating down: %v", err)
	}
	if err := m.Up(); err != nil {
		log.Fatalf("error while migrating up: %v", err)
	}

	app := fiber.New(fiber.Config{
		Immutable: true,
	})
	app.Use(cors.New())
	app.Use(helmet.New())

	h := handlers.NewHandler(db, validator.New())

	app.Get("/metrics", monitor.New())
	app.Get("/statusz", h.ServerStatus)
	app.Get("/users", h.GetAllUsers)
	app.Post("/user", h.CreateUser)

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go cleanAllSoftDelete(ctx, wg, db)

	osChan := make(chan os.Signal)
	signal.Notify(osChan, os.Interrupt)
	go shutdownServerWhenInterrupt(osChan, app, db, cancel, wg)

	log.Infof("starting server on PORT %v...", PORT)
	if err := app.ListenTLS(":"+PORT, "server.crt", "server.key"); err != nil {
		log.Fatalf("cannot start server on PORT %v: %v\n", PORT, err)
	}
}
