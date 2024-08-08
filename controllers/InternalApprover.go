package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/palladiumkenya/individual-data-request-backend/internal/config"
	"github.com/palladiumkenya/individual-data-request-backend/internal/db"
	"github.com/palladiumkenya/individual-data-request-backend/internal/models"

	"log"
	"net/http"
)

func GetRequests(c *gin.Context) {
	cfg := config.LoadConfig()
	DB, err := db.Connect(cfg.DatabaseURL)

	// Retrieve a requests
	requests, err := models.GetRequests(DB)
	if err != nil {
		log.Fatalf("Error retrieving requests: %v\n", err)
	}

	fmt.Printf("Retrieved requests: %+v\n", requests)

	// Set the Content-Type header and write the JSON response
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   requests,
	})

}
