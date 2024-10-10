package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/palladiumkenya/individual-data-request-backend/internal/db"
	"log"
	"net/http"
)

func GetUserRole(c *gin.Context) {
	DB, err := db.Connect()
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	emailStr := c.Query("email")
	if emailStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Email Provided"})
		log.Fatalf("Invalid Email Provided")
		return
	}

	var count int64

	// Check if the user is a requester
	result := DB.Table("requesters").Where("email =?", emailStr).Count(&count)
	if result.Error != nil {
		log.Printf("Error checking if the user is a requester: %v\n", result.Error)
	}
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data": gin.H{
				"role": "requester",
			},
		})
		return
	}

	// Check if the user is an approver
	result = DB.Table("approvers").Where("email =?", emailStr).Count(&count)
	if result.Error != nil {
		log.Printf("Error checking if the user is an approver: %v\n", result.Error)
	}
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data": gin.H{
				"role": "approver",
			},
		})
		return
	}

	// Check if the user is an analyst
	result = DB.Table("assignees").Where("email =?", emailStr).Count(&count)
	if result.Error != nil {
		log.Printf("Error checking if the user is an analyst: %v\n", result.Error)
	}
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data": gin.H{
				"role": "analyst",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"role": nil,
		},
	})
}

