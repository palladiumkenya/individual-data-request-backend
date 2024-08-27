package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatalf("DATABASE_URL is not set in the environment variables")
	}
	fmt.Printf("Database URL: %s\n", databaseURL)
	return Config{
		DatabaseURL: databaseURL,
	}
}
