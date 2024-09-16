package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/palladiumkenya/individual-data-request-backend/internal/db"
	"github.com/palladiumkenya/individual-data-request-backend/internal/models"
	"gorm.io/gorm"

	"log"
	"net/http"
)

func GetRequests(c *gin.Context) {
	DB, err := db.Connect()

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
	DB, err := db.Connect()
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

	DB, err := db.Connect()

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

	DB, err := db.Connect()

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

	DB, err := db.Connect()

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
	DB, err := db.Connect()

	var newApproval *models.Approvals

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

	log.Printf("Return approval results")
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   approval,
	})

}
