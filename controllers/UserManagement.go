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

func GetUserRole(c *gin.Context) {
	DB, err := db.Connect()
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	emailStr := c.Query("email")
	if emailStr == "" || emailStr == "null" {
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
	DB, err := db.Connect()
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	var requester models.Requesters
	if err := c.BindJSON(&requester); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}
	fmt.Printf("requestStatus: %+v\n", requester)

	id, err := models.CreateRequester(DB, requester)
	if err != nil {
		log.Printf("Error checking if the user is a requester: %v\n", err)
		c.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Requester created successfully",
		"data": gin.H{
			"role": "requester",
			"id":   id,
		},
	})
	return

}

func CreateNewApprover(c *gin.Context) {
	DB, err := db.Connect()
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	var approver models.Approvers
	if err := c.BindJSON(&approver); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	id, err := models.CreateApprover(DB, approver)
	if err != nil {
		log.Printf("Error checking if the user is a requester: %v\n", err)
		c.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Approver created successfully",
		"id":      id,
	})
	return
}

func DeleteApprover(c *gin.Context) {
	DB, err := db.Connect()
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}
	idStr := c.Query("id")
	if idStr == "" || idStr == "null" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Approver UUID"})
		log.Fatalf("Error invalid UUID: %v\n", err)
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Approver UUID"})
		log.Fatalf("Error invalid UUID: %v\n", err)
		return
	}
	err = models.DeleteApprover(DB, id)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Approver deleted successfully",
	})
	return
}
