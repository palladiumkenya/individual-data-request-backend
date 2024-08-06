package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/palladiumkenya/individual-data-request-backend/controllers"
)

func Handlers(router *gin.Engine) {
	// Define routes here
	router.GET("/api_health", controllers.GetApiHealth) // Test endpoint

}
