package main

import (
	"context"
	"os"
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

func cleanSoftDeleted(ctx context.Context, db *gorm.DB) {
	ticker := time.NewTicker(10 * time.Minute)

	for {
		select {
		case <-ticker.C:
			for _, model := range model.Models {
				if err := db.Unscoped().Where("deleted_at IS NOT NULL").Delete(model).Error; err != nil {
					log.Errorf("Cannot clean up soft deleted from a model: %v", err)
				}
			}
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}

func main() {
	if err := godotenv.Load(".env.local"); err != nil {
		log.Fatalf("Cannot load .env.local file: %v\n", err)
	}

	dsn := os.Getenv("DB_URI")
	if dsn == "" {
		log.Fatal("Environment variable has no 'DB_URI'")
	}
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8000"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Cannot connect to database: %v\n", err)
	}

	db.AutoMigrate(model.Models...)

	app := fiber.New(fiber.Config{
		Immutable: true,
	})
	app.Use(cors.New())
	app.Use(helmet.New())

	h := handlers.Handler{DB: db, Validator: validator.New()}

	app.Get("/metrics", monitor.New())
	app.Get("/statusz", h.ServerStatus)
	app.Get("/users", h.GetAllUsers)
	app.Post("/user", h.CreateUser)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go cleanSoftDeleted(ctx, db)

	log.Infof("Starting server on port %v...", port)
	if err := app.ListenTLS(":"+port, "server.crt", "server.key"); err != nil {
		log.Fatalf("Cannot start server on port %v: %v\n", port, err)
	}

}
