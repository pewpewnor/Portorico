package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/joho/godotenv"
	"github.com/pewpewnor/portorico/server/handlers"
	"github.com/pewpewnor/portorico/server/model"
	"github.com/pewpewnor/portorico/server/validator"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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

	h := handlers.Handler{DB: db}
	validator.Init()

	// app.Get("/metrics", monitor.New())
	app.Get("/statusz", h.ServerStatus)
	app.Get("/users", h.GetAllUsers)
	app.Post("/user", h.CreateUser)

	log.Infof("Starting server on port %v...", port)
	if err := app.ListenTLS(":"+port, "server.crt", "server.key"); err != nil {
		log.Fatalf("Cannot start server on port %v: %v\n", port, err)
	}
}
