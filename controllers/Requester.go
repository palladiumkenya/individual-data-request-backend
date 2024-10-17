package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/palladiumkenya/individual-data-request-backend/internal/db"
	"github.com/palladiumkenya/individual-data-request-backend/internal/models"
	"log"
	"net/http"
)

func NewRequest(c *gin.Context) {
	DB, err := db.Connect()

	var request models.NewRequest
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}
	fmt.Printf("requestStatus: %+v\n", request)

	// Updated Status a requests
	id, err := models.CreateRequest(DB, request)
	if err != nil {
		log.Fatalf("Error retrieving requests: %v\n", err)
	}

	fmt.Printf("Created request: \n")

	// Set the Content-Type header and write the JSON response
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"id":      id,
			"message": "Request created successfully",
		},
	})

}

func GetRequesterRequests(c *gin.Context) {
	DB, err := db.Connect()
	requesterUuidStr := c.Query("requester")
	requesterUuid, err := uuid.Parse(requesterUuidStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Requester UUID"})
		log.Fatalf("Error invalid UUID: %v\n", err)
		return
	}

	// Retrieve a requests
	requests, err := models.GetRequesterRequests(DB, requesterUuid)
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
