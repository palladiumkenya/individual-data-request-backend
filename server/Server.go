package server

import (
	"github.com/gin-gonic/gin"
	"github.com/palladiumkenya/individual-data-request-backend/routes"
)

var router = gin.Default()

func Run() {
	routes.Handlers(router)
	// Start server on port 8080
	router.Run(":8000")
}
