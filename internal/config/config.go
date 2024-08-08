package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DatabaseURL string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Read environment variables
	database_url := os.Getenv("DATABASE_URL")

	fmt.Println(database_url)

	return Config{
		DatabaseURL: database_url,
	}
}
