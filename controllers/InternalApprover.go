package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/palladiumkenya/individual-data-request-backend/internal/db"
	"github.com/palladiumkenya/individual-data-request-backend/internal/models"
	"gorm.io/gorm"

	"log"
	"net/http"
)

func GetRequests(c *gin.Context) {
	//cfg := config.LoadConfig()
	DB, err := db.Connect()

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

func Get_Internal_approval(c *gin.Context) {
	ID := c.Param("id")
	DB, err := db.Connect()

	// Retrieve a requests
	requests, err := models.GetRequestByID(DB, uuid.MustParse(ID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Return 404 if not found
			c.JSON(http.StatusNotFound, gin.H{"error": "Request not found"})
			return
		}
		// Handle other potential errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	fmt.Printf("Retrieved requests: %+v\n", requests)

	// Set the Content-Type header and write the JSON response
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   requests,
	})

}
