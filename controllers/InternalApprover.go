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

func GetInternalApproval(c *gin.Context) {
	ID := c.Param("id")

	DB, err := db.Connect()

	// Retrieve a request
	requests, err := models.GetApprovalByID(DB, uuid.MustParse(ID))
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

func GetApproval(c *gin.Context) {
	ID := c.Param("id")
	approvalType := c.Param("type")

	DB, err := db.Connect()

	// Retrieve a request
	requests, err := models.GetApprovalByIDAndType(DB, uuid.MustParse(ID), approvalType)
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

func ApproverAction(c *gin.Context) {
	DB, err := db.Connect()

	// create approval
	var newApproval *models.Approvals

	// Call BindJSON to bind the received JSON to approval
	if err := c.BindJSON(&newApproval); err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, newApproval)

	approval, err := models.CreateApproval(DB, newApproval)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Return 404 if not found
			c.JSON(http.StatusNotFound, gin.H{"error": "Approval failed to be created"})
			return
		}
		// Handle other potential errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Set the Content-Type header and write the JSON response
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   approval,
	})

}
