package main

import (
	"fmt"
	"individual-data-request-backend/internal/config"
	"individual-data-request-backend/internal/db"
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

func main() {
	//cwd, err := os.Getwd()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Current Working Directory:", cwd)
	//// Read the entire file content
	//content, err := ioutil.ReadFile(cwd + "/supporting-documents/ethics-approval.pdf")
	//if err != nil {
	//	log.Fatal(err)
	//}

	// Print the content as a string
	//fmt.Println(string(content))
	http.HandleFunc("/pdf", pdfHandler)
	fmt.Println("Serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
