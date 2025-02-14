package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/maxwelbm/alkemy-g6/internal/application"
	"github.com/maxwelbm/alkemy-g6/internal/application/sqlconfig"
)

// @title FrescosAPI
// @version 1.0
// @description This is the FrescosAPI documentation.
// @host localhost:8080
// @BasePath /
func main() {
	// Initializes env variables from .env file
	env := os.Getenv("GO_ENVIRONMENT")
	if env != "production" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	dbConfig := sqlconfig.NewConn(env)
	// - config
	cfg := &application.ConfigServerChi{
		DB:   dbConfig,
		Addr: ":8080",
	}

	// Starts web server with config from .env file
	app := application.NewServerChi(cfg)

	if err := app.Run(); err != nil {
		log.Fatalf("Error initializing server %s", err)
		return
	}
}
