package controllers

import (
	"github.com/gin-gonic/gin"
	uuid2 "github.com/gofrs/uuid"
	"github.com/google/uuid"
	"github.com/palladiumkenya/individual-data-request-backend/internal/db"
	"github.com/palladiumkenya/individual-data-request-backend/internal/models"
	"github.com/palladiumkenya/individual-data-request-backend/services"
	"gorm.io/gorm"
	"log"
	"net/http"
	"path/filepath"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the file locally
	localFilePath := filepath.Join("uploads", file.Filename)
	if err := c.SaveUploadedFile(file, localFilePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Upload to Nextcloud and get the file URL
	remoteFilePath := "/idr/files/" + file.Filename
	fileURL, err := services.UploadFileToNextcloud(localFilePath, remoteFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file to Nextcloud"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "File uploaded successfully",
		"file_url": fileURL,
	})
}

func FetchFile(c *gin.Context) {
	RequestId := c.Param("request_id")
	FileType := c.Param("file_type")

	DB, err := db.Connect()

	pdfFile, err := models.FetchFiles(DB, FileType, uuid2.UUID(uuid.MustParse(RequestId)))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Request File record not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	log.Printf("Return file url results")
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   pdfFile.FileURL,
	})
}
