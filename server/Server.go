package server

import (
	"github.com/gin-contrib/cors"
	"github.com/palladiumkenya/individual-data-request-backend/internal/db"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/palladiumkenya/individual-data-request-backend/routes"
)

var router = gin.Default()

func Run() {
	// CORS configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// migrate models
	db.MigrateDB()

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
