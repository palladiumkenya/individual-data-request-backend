package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/palladiumkenya/individual-data-request-backend/internal/db"
	"github.com/palladiumkenya/individual-data-request-backend/internal/models"
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

	// Check if the user is a requester
	requester, err := models.CheckUserRequester(DB, emailStr)
	if err != nil {
		log.Printf("Error checking if the user is a requester: %v\n", err)
		return
	}
	if requester.Email == emailStr {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data": gin.H{
				"role": "requester",
				"id":   requester.ID,
			},
		})
		return
	}

	// Check if the user is an approver
	approver, err := models.CheckUserApprover(DB, emailStr)
	if err != nil {
		log.Printf("Error checking if the user is an approver: %v\n", err)
	}
	if approver.Email == emailStr {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data": gin.H{
				"role": "approver",
				"id":   approver.ID,
				"type": approver.Approver_Type,
			},
		})
		return
	}

	// Check if the user is an analyst
	analyst, err := models.CheckUserAnalyst(DB, emailStr)
	if err != nil {
		log.Printf("Error checking if the user is an analyst: %v\n", err)
	}
	if analyst.Email == emailStr {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data": gin.H{
				"role": "analyst",
				"id":   analyst.ID,
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"role": nil,
			"id":   nil,
		},
	})
}

func CreateNewRequester(c *gin.Context) {
	_, err := db.Connect()
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

}
