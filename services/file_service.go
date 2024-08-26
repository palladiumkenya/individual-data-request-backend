package services

import (
	"fmt"
	"github.com/studio-b12/gowebdav"
	"os"
)

func UploadFileToNextcloud(localFilePath, remoteFilePath string) (string, error) {
	fmt.Printf("Uploading file %s to Nextcloud\n", localFilePath)
	nextcloudURL := os.Getenv("NEXTCLOUD_URL")
	username := os.Getenv("NEXTCLOUD_USERNAME")
	password := os.Getenv("NEXTCLOUD_PASSWORD")

	client := gowebdav.NewClient(nextcloudURL, username, password)

	file, err := os.Open(localFilePath)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return "", err
	}
	defer file.Close()

	data, err := os.ReadFile(localFilePath)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return "", err
	}

	// Upload the file
	err = client.Write(remoteFilePath, data, 0644)
	if err != nil {
		fmt.Printf("Error uploading file: %s\n", err)
		return "", err
	}

	fmt.Printf("File uploaded successfully to %s\n", remoteFilePath)

	// Generate the file URL
	fileURL := fmt.Sprintf("%s/%s", nextcloudURL, remoteFilePath)

	return fileURL, nil
}

