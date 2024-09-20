package services

import (
	"fmt"
	"github.com/studio-b12/gowebdav"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// UploadFileToNextcloud uploads a file to Nextcloud and creates necessary folders.
func UploadFileToNextcloud(localFilePath, remoteFilePath string) (string, error) {
	fmt.Printf("Uploading file %s to Nextcloud\n", localFilePath)
	nextcloudURL := os.Getenv("NEXTCLOUD_URL")
	username := os.Getenv("NEXTCLOUD_USERNAME")
	password := os.Getenv("NEXTCLOUD_PASSWORD")

	client := gowebdav.NewClient(nextcloudURL, username, password)
	folderPath := strings.TrimSuffix(remoteFilePath, "/"+getFileName(remoteFilePath))

	// Ensure all parent folders exist
	err := createFoldersIfNeeded(nextcloudURL, username, password, folderPath)
	if err != nil {
		return "", err
	}

	// Read the file
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

// createFoldersIfNeeded ensures all required folders in the path exist.
func createFoldersIfNeeded(baseURL, username, password, folderPath string) error {
	folderParts := strings.Split(folderPath, "/")
	currentPath := ""

	for _, part := range folderParts {
		if part == "" {
			continue
		}

		currentPath = strings.TrimRight(currentPath+"/"+part, "/")
		if !folderExists(baseURL, username, password, currentPath) {
			fmt.Printf("Creating folder: %s\n", currentPath)
			err := createFolder(baseURL, username, password, currentPath)
			if err != nil {
				return fmt.Errorf("error creating folder %s: %w", currentPath, err)
			}
		}
	}

	return nil
}

// folderExists checks if a folder exists using a PROPFIND request.
func folderExists(baseURL, username, password, folderURL string) bool {
	req, err := http.NewRequest("PROPFIND", baseURL+folderURL, nil)
	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		return false
	}
	req.SetBasicAuth(username, password)
	req.Header.Set("Depth", "1") // Check the depth of the folder

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		return false
	}
	defer resp.Body.Close()

	// Return true if we get a 207 Multi-Status or 200 OK, indicating the folder exists
	return resp.StatusCode == http.StatusMultiStatus || resp.StatusCode == http.StatusOK
}

// createFolder creates a folder using a MKCOL request.
func createFolder(baseURL, username, password, folderURL string) error {
	fmt.Printf("Creating folder:%s %s\n", baseURL, folderURL)
	req, err := http.NewRequest("MKCOL", baseURL+folderURL, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	req.SetBasicAuth(username, password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to create folder. Status code: %d, Response: %s", resp.StatusCode, body)
	}

	return nil
}

// Helper function to get the file name from a path
func getFileName(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}
