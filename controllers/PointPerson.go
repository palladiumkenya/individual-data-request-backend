package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/palladiumkenya/individual-data-request-backend/internal/db"
	"github.com/palladiumkenya/individual-data-request-backend/internal/models"
	"github.com/palladiumkenya/individual-data-request-backend/services"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

//var DB, err = db.Connect()

func AssignToAnalystAction(c *gin.Context) {
	DB, err := db.Connect()
	Id := c.Param("requestid")
	assigneeId := c.Param("assigneeId")
	analystID, err := uuid.Parse(assigneeId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignee UUID"})
		log.Fatalf("Error invalid UUID: %v\n", err)
		return
	}

	//assign, err := models.AssignRequestToAnalyst(DB, uuid.MustParse(Id), uuid.Parse(analystID))
	assign, err := models.AssignRequestToAnalyst(DB, uuid.MustParse(Id), analystID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Approval failed to be created"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	sendAlert, err := SendAssigneeEmailNotifications(analystID, c)
	if err != nil {
		log.Fatalf(sendAlert, " : Error sending emails: %v\n", err)
	}

	log.Printf("Return approval results")
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   assign,
	})
}

func GetPointPersonByEmail(c *gin.Context) {
	email := c.Param("email")

	//DB, err := db.Connect()

	approvals, err := models.GetPointPersonByEmail(DB, email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Pointperson not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	log.Printf("Return pointperson results")
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   approvals,
	})

}

func SendAssigneeEmailNotifications(analystID uuid.UUID, c *gin.Context) (string, error) {
	// send acknowledgement email to analyst
	template := "email_templates/request_reviewer_notifications.html"
	frontendUrl := os.Getenv("FRONTEND_URL")

	assignee, _ := models.GetAssigneeByID(DB, analystID)
	email := assignee.Email
	subject := "Request Assigned to You!"
	message := "A request has been assigned to you. Please login to IDR Platform and navigate to your dashboard to view the request assigned to you."

	body := map[string]interface{}{
		"request_id":    "",
		"request_url":   frontendUrl + "/assignee/dashboard",
		"frontend_url":  frontendUrl,
		"message_title": "Status progress: " + subject,
		"message":       message,
	}

	emailId, err := services.SendEmailAlerts(subject, body, email, template, c)
	if err != nil {
		log.Fatalf("Error sending email: %v\n", err)
	} else {
		fmt.Printf("Email sent successfully. Email ID: %s\n", emailId)
	}

	return emailId, err

}
