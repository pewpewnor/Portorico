package main

import (
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/pewpewnor/portorico/server/handlers"
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

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Cannot connect to database: %v\n", err)
	}

	db.AutoMigrate()

	e := echo.New()
	h := handlers.Handler{DB: db}

	e.GET("/statusz", h.ServerStatus)

	if err := e.StartTLS(":8000", "server.crt", "server.key"); err != http.ErrServerClosed {
		log.Fatalf("Cannot start server on port 8000: %v\n", err)
	}
}
