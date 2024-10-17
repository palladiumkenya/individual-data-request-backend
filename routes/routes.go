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

	router.GET("/analysts/jobs", controllers.GetApprovedTasks)        // analysts get all requests
	router.GET("/analysts/job", controllers.GetApprovedTask)          // analysts get one request
	router.PUT("/analysts/job/:id", controllers.UpdateAnalystRequest) // analysts update request status

	router.POST("/upload", controllers.UploadFile) // upload to nextcloud
	router.GET("/fetch_file/:file_type/:request_id", controllers.FetchFile)
	router.GET("/fetch_request_files/:request_id", controllers.FetchFiles)

	router.POST("/request/create", controllers.NewRequest)                 // New request
	router.GET("/request/requester/get", controllers.GetRequesterRequests) // get requests

	router.GET("/user/role", controllers.GetUserRole)                  // get user role
	router.POST("/user/new_requester", controllers.CreateNewRequester) // create new requester

	router.GET("/approvals/:type", controllers.GetAllApprovals)                             // get all approvals
	router.POST("/approval/action", controllers.ApproverAction)                             // approve or reject requests
	router.GET("/approval/:type/:id", controllers.GetApproval)                              // get approval page data
	router.GET("/approvals/count/:type", controllers.GetApprovalsCount)                     // get all approvals
	router.GET("/request/:id", controllers.GetRequestForApproval)                           // get requests
	router.POST("/assign/action/:requestid/:assigneeId", controllers.AssignToAnalystAction) //  assing analyst

	router.POST("/new_review_thread", controllers.CreateReviewThread)      // create review thread
	router.POST("/add_review", controllers.AddReview)                      // add review
	router.GET("/get_reviews/:thread_id", controllers.GetReviewsForThread) // get reviews

}
