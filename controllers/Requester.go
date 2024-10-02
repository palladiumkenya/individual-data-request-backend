package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
