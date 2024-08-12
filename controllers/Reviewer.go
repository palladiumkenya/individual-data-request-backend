package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palladiumkenya/individual-data-request-backend/internal/models"
	"gorm.io/gorm"
)

func CreateReviewThread(c *gin.Context, db *gorm.DB) {
	var thread models.ReviewThread
	if err := c.ShouldBindJSON(&thread); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&thread)
	c.JSON(http.StatusOK, thread)
}

func AddReview(c *gin.Context, db *gorm.DB) {
	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&review)
	c.JSON(http.StatusOK, review)
}

func GetReviewsForThread(c *gin.Context, db *gorm.DB) {
	threadID := c.Param("thread_id")
	var reviews []models.Review
	db.Where("review_thread_id = ?", threadID).Find(&reviews)
	c.JSON(http.StatusOK, reviews)
}
