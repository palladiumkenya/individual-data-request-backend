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
