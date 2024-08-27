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

func GetApprovedTasks(c *gin.Context) {
	DB, err := db.Connect()
	assigneeUuidStr := c.Query("assignee")
	assigneeUuid, err := uuid.Parse(assigneeUuidStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignee UUID"})
		log.Fatalf("Error invalid UUID: %v\n", err)
		return
	}

	// Retrieve a requests
	requests, err := models.GetAssigneeTasks(DB, assigneeUuid)
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
