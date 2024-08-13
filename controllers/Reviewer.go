package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palladiumkenya/individual-data-request-backend/internal/models"
	"github.com/palladiumkenya/individual-data-request-backend/internal/db"
)

func CreateReviewThread(c *gin.Context) {
	var thread models.ReviewThread
	if err := c.ShouldBindJSON(&thread); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Create(&thread)
	c.JSON(http.StatusOK, thread)
}

func AddReview(c *gin.Context) {
	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Create(&review)
	c.JSON(http.StatusOK, review)
}

func GetReviewsForThread(c *gin.Context) {
	threadID := c.Param("thread_id")
	var reviews []models.Review
	db.DB.Where("review_thread_id = ?", threadID).Find(&reviews)
	c.JSON(http.StatusOK, reviews)
}
