package main

import (
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/maxwelbm/alkemy-g6/internal/application"
)

func main() {
	// Initializes env variables from .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// - config
	cfg := &application.ConfigServerChi{
		DB: &mysql.Config{
			User:      os.Getenv("DB_USER"),
			Passwd:    os.Getenv("DB_PASS"),
			Net:       "tcp",
			Addr:      os.Getenv("DB_ADDR"),
			DBName:    os.Getenv("DB_NAME"),
			ParseTime: true,
		},
		Addr:           os.Getenv("SVR_PORT"),
		LoaderFilePath: os.Getenv("DB_PATH"),
	}

	// Starts web server with config from .env file
	app := application.NewServerChi(cfg)

	if err := app.Run(); err != nil {
		log.Fatalf("Error initializing server %s", err)
		return
	}
}
