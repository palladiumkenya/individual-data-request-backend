package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Create new Gin router
	router := gin.Default()

	// Test endpoint
	router.GET("/api_health", func(ctx *gin.Context) {
		ctx.String(200, "Server is up and running!")
	})

	// Start server on port 8080
	router.Run(":8080")
}
