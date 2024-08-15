package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palladiumkenya/individual-data-request-backend/internal/db"
	"github.com/palladiumkenya/individual-data-request-backend/internal/models"
)

func CreateReviewThread(c *gin.Context) {
	var err error
	DB, err := db.Connect()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	var thread models.ReviewThread
	if err := c.ShouldBindJSON(&thread); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Create(&thread)
	c.JSON(http.StatusOK, thread)
}

func AddReview(c *gin.Context) {
	var err error
	DB, err := db.Connect()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Create(&review)
	c.JSON(http.StatusOK, review)
}

func GetReviewsForThread(c *gin.Context) {
	var err error
	DB, err := db.Connect()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	threadID := c.Param("thread_id")
	var reviews []models.Review
	DB.Where("review_thread_id = ?", threadID).Find(&reviews)
	c.JSON(http.StatusOK, reviews)
}
