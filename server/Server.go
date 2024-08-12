package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/palladiumkenya/individual-data-request-backend/routes"
)

var router = gin.Default()

func Run() {
	routes.Handlers(router)

	// Load .env file
	godotenv.Load(".env")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Start server on port 8080
	router.Run(`:8081`)
}
