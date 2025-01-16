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
	"strconv"
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

func GetApprovedTask(c *gin.Context) {
	DB, err := db.Connect()

	idUuidStr := c.Query("id")
	idUuid, err := uuid.Parse(idUuidStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request UUID"})
		log.Fatalf("Error invalid UUID: %v\n", err)
		return
	}

	// Retrieve a requests
	request, err := models.GetAssigneeTask(DB, idUuid)
	if err != nil {
		log.Fatalf("Error retrieving requests: %v\n", err)
	}

	fmt.Printf("Retrieved requests: %+v\n", request)

	// Set the Content-Type header and write the JSON response
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   request,
	})

}

func UpdateAnalystRequest(c *gin.Context) {
	DB, err := db.Connect()

	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request id"})
		log.Fatalf("Error invalid int: %v\n", err)
		return
	}

	var requestStatus models.UpdateStatusRequest
	if err := c.BindJSON(&requestStatus); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}
	fmt.Printf("requestStatus: %+v\n", requestStatus)

	// Updated Status a requests
	err = models.UpdateRequestStatus(DB, idInt, requestStatus.Status)
	if err != nil {
		log.Fatalf("Error retrieving requests: %v\n", err)
	}

	fmt.Printf("Updated requests: \n")

	// Set the Content-Type header and write the JSON response
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   "Updated request",
	})

}

func GetAssignedAnalyst(c *gin.Context) {
	request_id := c.Param("request_id")

	//DB, err := db.Connect()

	approvals, err := models.GetAssignedAnalyst(DB, uuid.MustParse(request_id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Assignee not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	log.Printf("Return assignee results")
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   approvals,
	})

}
