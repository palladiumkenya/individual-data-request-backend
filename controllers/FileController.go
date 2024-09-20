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
	"path/filepath"
)

func UploadFile(c *gin.Context) {
	// unique id for folder name
	folderId := uuid.New().String()

	// Get destination folder
	destinationFolder := c.PostForm("destination") // either "files" or "supporting-documents"
	request := c.PostForm("request")               // request that this file is associated with
	var requestUuid *uuid.UUID
	if request != "" {
		parsedUuid, _ := uuid.Parse(request)
		requestUuid = &parsedUuid
	}

	// Get the file from the POST form
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the file locally
	localFilePath := filepath.Join("uploads", folderId, file.Filename)
	if err := c.SaveUploadedFile(file, localFilePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Upload to Nextcloud and get the file URL
	remoteFilePath := fmt.Sprintf("idr/%s/%s/%s", destinationFolder, folderId, file.Filename)
	fileURL, err := services.UploadFileToNextcloud(localFilePath, remoteFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file to Nextcloud"})
		return
	}

	// Save the file URL to the database
	// TODO:: Save And User ID
	DB, err := db.Connect()
	requestFile := models.Files{
		FileName:  file.Filename,
		FileURL:   fileURL,
		RequestId: requestUuid,
	}

	// Save the file details to the database
	if err := models.UploadFiles(DB, &requestFile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file details to the database"})
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
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "DB connection issue. Record not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	pdfFile, err := models.FetchFile(DB, FileType, uuid.MustParse(RequestId))
	//fileFound := func() string {
	//	if pdfFile != nil {
	//		return pdfFile.FileURL
	//	}
	//	return ""
	//}()

	log.Printf("Return file url results")

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   pdfFile,
	})
}

func FetchFiles(c *gin.Context) {
	RequestId := c.Param("request_id")

	DB, err := db.Connect()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "DB connection issue. Record not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	pdfFile, err := models.FetchFiles(DB, uuid.MustParse(RequestId))

	log.Printf("Return file url results")

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   pdfFile,
	})
}
