package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/maxwelbm/alkemy-g6/internal/application"
)

func main() {
	// Initializes env variables from .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Starts web server with config from .env file
	cfg := &application.ConfigServerChi{
		ServerAddress:  os.Getenv("SERVER_ADDRESS"),
		LoaderFilePath: os.Getenv("DB_PATH"),
	}
	app := application.NewServerChi(cfg)

	if err := app.Run(); err != nil {
		log.Fatalf("Error initializint server %s", err)
		return
	}
}
