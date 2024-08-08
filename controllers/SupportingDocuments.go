package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/palladiumkenya/individual-data-request-backend/internal/config"
	"github.com/palladiumkenya/individual-data-request-backend/internal/db"
	"io"
	"log"
	"net/http"
	"os"
)

func pdfHandler(w http.ResponseWriter, r *http.Request) {
	cfg := config.LoadConfig()
	db.Connect(cfg.DatabaseURL)
	// Define the path to the local PDF file
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current Working Directory:", cwd)

	pdfPath := cwd + "/supporting-documents/ethics-approval.pdf"

	// Open the PDF file
	file, err := os.Open(pdfPath)
	if err != nil {
		http.Error(w, "Unable to open PDF file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Set the content type to PDF
	w.Header().Set("Content-Type", "application/pdf")

	// Serve the file content
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Unable to serve PDF file", http.StatusInternalServerError)
	}
}

func getpdf(c *gin.Context) {

	http.HandleFunc("/pdf", pdfHandler)
	fmt.Println("Serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
