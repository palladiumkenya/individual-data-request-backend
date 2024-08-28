package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/palladiumkenya/individual-data-request-backend/controllers"
)

func Handlers(router *gin.Engine) {
	// Define routes here
	router.GET("/api_health", controllers.GetApiHealth) // Test endpoint
	router.POST("/send_mail", controllers.SendMail)     // Send test email

	router.GET("/requests", controllers.GetRequests) // get requests

	router.GET("/analysts/jobs", controllers.GetApprovedTasks)

	router.POST("/upload", controllers.UploadFile) // upload to nextcloud
	//router.GET("/get_upload", controllers.GetFile) // nextcloud download
	router.GET("/fetch_file/:file_type/:request_id", controllers.FetchFile)

	router.GET("/approvals/:type", controllers.GetAllApprovals)          // get all approvals
	router.POST("/internal_approval/action", controllers.ApproverAction) // approve or reject requests
	router.GET("/approval/:type/:id", controllers.GetApproval)           // get approval page data
	router.GET("/approvals/count/:type", controllers.GetApprovalsCount)  // get all approvals

	router.POST("/new_review_thread", controllers.CreateReviewThread)      // create review thread
	router.POST("/add_review", controllers.AddReview)                      // add review
	router.GET("/get_reviews/:thread_id", controllers.GetReviewsForThread) // get reviews

}
