package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/palladiumkenya/individual-data-request-backend/internal/config"
	"github.com/palladiumkenya/individual-data-request-backend/internal/db"
	"github.com/palladiumkenya/individual-data-request-backend/internal/models"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

func GetRequests(c *gin.Context) {
	cfg := config.LoadConfig()
	pool, err := db.Connect(cfg.DatabaseURL)

	// Retrieve a requests
	requests, err := models.GetRequests(context.Background(), pool)
	if err != nil {
		log.Fatalf("Error retrieving requests: %v\n", err)
	}

	fmt.Printf("Retrieved requests: %+v\n", requests)

	// Convert the data to JSON
	//jsonData, err := json.Marshal(requests)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	// Set the Content-Type header and write the JSON response
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   requests,
	})

}
