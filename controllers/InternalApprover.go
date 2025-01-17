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

var DB, err = db.Connect()

func GetRequests(c *gin.Context) {
	//DB, err := db.Connect()

	approvals, err := models.GetRequests(DB)
	if err != nil {
		log.Fatalf("Error retrieving approvals: %v\n", err)
	}

	log.Printf("Return approval results")
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   approvals,
	})
}

func GetRequestForApproval(c *gin.Context) {
	//DB, err := db.Connect()
	ID := c.Param("id")

	approvals, err := models.GetRequestByID(DB, uuid.MustParse(ID))
	if err != nil {
		log.Fatalf("Error retrieving requests for approval page: %v\n", err)
	}

	log.Printf("Return approval results")
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   approvals,
	})
}

func GetAllApprovals(c *gin.Context) {
	approvalType := c.Param("type")

	//DB, err := db.Connect()

	approvals, err := models.GetApprovalsByType(DB, approvalType)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Approvals not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	log.Printf("Return approval results")
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   approvals,
	})

}

func GetApprovalsCount(c *gin.Context) {
	approvalType := c.Param("type")

	//DB, err := db.Connect()

	approvals, err := models.GetApprovalsCounts(DB, approvalType)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Approvals not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	log.Printf("Return approval results")
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   approvals,
	})

}

func GetApproval(c *gin.Context) {
	ID := c.Param("id")
	approvalType := c.Param("type")

	//DB, err := db.Connect()

	approvals, err := models.GetApprovalByIDAndType(DB, uuid.MustParse(ID), approvalType)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Approval not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	log.Printf("Return approval results")
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   approvals,
	})

}

func ApproverAction(c *gin.Context) {
	//DB, err := db.Connect()

	var newApproval *models.CreateApprovalsStruct

	if err := c.BindJSON(&newApproval); err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, newApproval)

	approval, err := models.CreateApproval(DB, newApproval)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Approval failed to be created"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	sendAlert, err := SendEmailNotifications(newApproval.Request_id.String(), newApproval.Approver_type, newApproval.Approved, c)
	if err != nil {
		log.Fatalf(sendAlert, " : Error sending emails: %v\n", err)
	}

	log.Printf("Return approval results")
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   approval,
	})

}

func GetAnalysts(c *gin.Context) {
	//DB, err := db.Connect()

	approvals, err := models.GetAnalysts(DB)
	if err != nil {
		log.Fatalf("Error retrieving approvals: %v\n", err)
	}

	log.Printf("Return approval results")
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   approvals,
	})
}

func SendEmailNotifications(request_id string, approver_type string, approved *bool, c *gin.Context) (string, error) {
	// Launch background job to send email alert
	// send acknowledgement email to requester
	template := "email_templates/request_reviewer_notifications.html"
	frontendUrl := os.Getenv("FRONTEND_URL")

	//reviewer, _ := models.GetApproversByType(DB, approver_type)
	//externalReviewersEmails, _ := models.GetAllExternalApprovers(DB)
	requester, _ := models.GetRequesterByRequestID(DB, uuid.MustParse(request_id))
	email := requester.Email
	//subject := approved ? "Request is Aprroved!" :"Request is Rejected!"
	subject := "Request is Rejected!"
	message := "After review we regret to inform you that your data request was declined. Please login to view comments of the review."
	if *approved {
		subject = "Request is Approved!"
		message = "After review we're happy to inform you that your data request was approved " + approver_type + "ly. It will now move to the next stage of review or assignment. "
	}
	body := map[string]interface{}{
		"request_id":    request_id,
		"request_url":   frontendUrl + "/requester/request-details?id=" + request_id,
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

	if *approved {
		getRequest, _ := models.GetRequestByID(DB, uuid.MustParse(request_id))
		if approver_type == "internal" {
			// send review email to reviewer
			externalReviewersEmails, _ := models.GetAllExternalApprovers(DB)

			for _, recipient := range externalReviewersEmails {
				email = recipient
				subject = "External Review stage"
				template = "email_templates/request_reviewer_notifications.html"
				body = map[string]interface{}{
					"request_id":    request_id,
					"request_url":   frontendUrl + "/external/action/" + string(rune(getRequest.ReqId)) + "?type=external&id=" + request_id,
					"message_title": "Your Review is Needed",
					"message":       "A request has been approved  and transitioned to the External Review stage. Please navigate to your dashboard to review",
				}

				emailId, err = services.SendEmailAlerts(subject, body, email, template, c)
				if err != nil {
					log.Fatalf("Error sending email: %v\n", err)
				} else {
					fmt.Printf("Email sent successfully. Email ID: %s\n", emailId)
				}
			}
		} else {
			// send review email to point person
			pointpersonsEmails, _ := models.GetPointPersonsEmails(DB)
			for _, recipient := range pointpersonsEmails {

				email = recipient
				subject = "Request Assignment stage"
				template = "email_templates/request_reviewer_notifications.html"
				body = map[string]interface{}{
					"request_id":    request_id,
					"request_url":   frontendUrl + "/assign/action/" + string(rune(getRequest.ReqId)) + "?type=external&id=" + request_id,
					"message_title": "Assign to Analysts!",
					"message":       "A request has been approved by both the Internal and External Reviewers. Please navigate to your dashboard to assign it to an analyst.",
				}
				emailId, err = services.SendEmailAlerts(subject, body, email, template, c)
				if err != nil {
					log.Fatalf("Error sending email: %v\n", err)
				} else {
					fmt.Printf("Email sent successfully. Email ID: %s\n", emailId)
				}
			}
		}
		//emailId, err = services.SendEmailAlerts(subject, body, email, template, c)
		//if err != nil {
		//	log.Fatalf("Error sending email: %v\n", err)
		//} else {
		//	fmt.Printf("Email sent successfully. Email ID: %s\n", emailId)
		//}
	}

	return emailId, err
}

func GetApprovers(c *gin.Context) {
	DB, err := db.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}
	approvers, err := models.GetApprovers(DB)
	if err != nil {
		log.Fatalf("Error retrieving approvers: %v\n", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   approvers,
	})
	return
}

func GetApproversByEmails(c *gin.Context) {
	email := c.Param("email")

	//DB, err := db.Connect()

	approvals, err := models.GetApproversByEmail(DB, email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Approvers not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	log.Printf("Return approvers results")
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   approvals,
	})

}

func GetRejectedApproval(c *gin.Context) {
	request_id := c.Param("request_id")

	//DB, err := db.Connect()

	approvals, err := models.GetRejectedApproval(DB, uuid.MustParse(request_id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "No reject request reviews not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	log.Printf("Return rejected approval results")
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   approvals,
	})

}

func GetAllExternalApprovers(c *gin.Context) {
	DB, err := db.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}
	approvers, err := models.GetAllExternalApprovers(DB)
	if err != nil {
		log.Fatalf("Error retrieving external approvers: %v\n", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   approvers,
	})
	return
}
