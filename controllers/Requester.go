package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/palladiumkenya/individual-data-request-backend/internal/db"
	"github.com/palladiumkenya/individual-data-request-backend/internal/models"
	"github.com/palladiumkenya/individual-data-request-backend/services"
	"log"
	"net/http"
	"os"
	"time"
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
	savedRequest, err := models.CreateRequest(DB, request)
	if err != nil {
		log.Fatalf("Error retrieving requests: %v\n", err)
	}

	fmt.Printf("Created request: \n")

	// Set the Content-Type header and write the JSON response
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"id":      savedRequest.ID,
			"message": "Request created successfully",
		},
	})

	// Launch background job to send email alert
	go func() {
		// send acknowledgement email to requester
		template := "email_templates/requester_new_request_alert.html"
		frontendUrl := os.Getenv("FRONTEND_URL")
		body := map[string]interface{}{
			"request_id":   savedRequest.ID,
			"request_url":  frontendUrl + "/requester/request-details?id=" + savedRequest.ID.String(),
			"frontend_url": frontendUrl,
		}

		requester, _ := models.GetRequesterByID(DB, request.Requestor_id)
		email := requester.Email
		subject := "Request Created Successfully"

		emailId, err := services.SendEmailAlerts(subject, body, email, template, c)
		if err != nil {
			log.Fatalf("Error sending email: %v\n", err)
		} else {
			fmt.Printf("Email sent successfully. Email ID: %s\n", emailId)
		}

		// send review email to reviewer
		reviewer, err := models.GetRandomApprover(DB, "InternalApprover")
		if err != nil {
			log.Fatalf("Error getting approver: %v\n", err)
			return
		}
		email = reviewer.Email
		subject = "New Request Needs Review"
		template = "email_templates/reviewer_new_request_alert.html"
		body = map[string]interface{}{
			"request_id":   savedRequest.ID,
			"request_url":  frontendUrl + "/internal/action/" + string(rune(savedRequest.ReqId)) + "?type=internal&id=" + savedRequest.ID.String(),
			"due_date":     savedRequest.Date_Due.Format(time.ANSIC),
			"date_created": savedRequest.Created_Date.Format(time.ANSIC),
		}
		emailId, err = services.SendEmailAlerts(subject, body, email, template, c)
		if err != nil {
			log.Fatalf("Error sending email: %v\n", err)
		} else {
			fmt.Printf("Email sent successfully. Email ID: %s\n", emailId)
		}
	}()

}

func GetRequesterRequests(c *gin.Context) {
	DB, err := db.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		log.Fatalf("Database connection failed: %v\n", err)
		return
	}

	requesterUuidStr := c.Query("requester")
	if requesterUuidStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Requester UUID is required"})
		return
	}

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

func GetRequestDetails(c *gin.Context) {
	DB, err := db.Connect()
	requestUuidStr := c.Query("request_id")
	requestUuid, err := uuid.Parse(requestUuidStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request UUID"})
		log.Fatalf("Error invalid UUID: %v\n", err)
		return
	}
	// Retrieve a requests
	request, err := models.GetRequesterRequestDetails(DB, requestUuid)
	if err != nil {
		log.Fatalf("Error retrieving requests: %v\n", err)
	}
	// Get files attached to the request
	files, err := models.FetchFiles(DB, requestUuid)
	if err != nil {
		log.Printf("Error retrieving files: %v\n", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"request": request,
			"files":   files,
		},
	})

}
