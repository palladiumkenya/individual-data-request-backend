package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/palladiumkenya/individual-data-request-backend/controllers"
)

func Handlers(router *gin.Engine) {
	// Define routes here
	router.GET("/api_health", controllers.GetApiHealth) // Test endpoint
	router.POST("/send_mail", controllers.SendMail)     // Send test email

	router.GET("/requests", controllers.GetRequests) // get pdf

}
