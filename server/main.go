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
	"github.com/joho/godotenv"
	"github.com/pewpewnor/portorico/server/handlers"
	"github.com/pewpewnor/portorico/server/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func cleanAllSoftDelete(ctx context.Context, wg *sync.WaitGroup, db *gorm.DB) {
	defer wg.Done()
	ticker := time.NewTicker(10 * time.Minute)

	for {
		select {
		case <-ticker.C:
			for _, model := range model.Models {
				if err := db.Unscoped().Where("deleted_at IS NOT NULL").Delete(model).Error; err != nil {
					log.Errorf("cannot clean soft deleted data from model %v: %v", model, err)
				}
			}
			log.Info("all soft deleted data has been cleaned")
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}

func shutdownServerWhenInterrupt(osChan chan os.Signal, app *fiber.App, cancel context.CancelFunc, wg *sync.WaitGroup) {
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
}

func main() {
	if err := godotenv.Load(".env.local"); err != nil {
		log.Fatalf("cannot load .env.local file: %v\n", err)
	}

	dsn := os.Getenv("DB_URI")
	if dsn == "" {
		log.Fatal("environment variable has no 'DB_URI'")
	}
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8000"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("cannot connect to database: %v\n", err)
	}

	db.Migrator().DropTable(model.Models...)
	db.AutoMigrate(model.Models...)

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
	go shutdownServerWhenInterrupt(osChan, app, cancel, wg)

	log.Infof("starting server on port %v...", port)
	if err := app.ListenTLS(":"+port, "server.crt", "server.key"); err != nil {
		log.Fatalf("cannot start server on port %v: %v\n", port, err)
	}
}
