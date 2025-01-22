package main

import (
	"example/pdfgenerator/controllers"
	"example/pdfgenerator/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.MigrateDB()
	initializers.InitMinioClient()
}

func main() {

	r := gin.Default()

	// Set up CORS middleware
	config := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}
	r.Use(cors.New(config))

	// r.POST("/upload-template", controllers.UploadTemplate)
	r.POST("/till-operator-request-float", controllers.TillOperatorRequestFloat)
	r.POST("/branch-manager-request-float", controllers.BranchManagerRequestFloat)
	r.POST("/branch-manager-approve-float", controllers.BranchManagerApproveFloat)
	r.POST("agent-admin-approve-float", controllers.AgentAdminApproveFloat)

	//get endpoints
	r.GET("/branch-manager-float-requests", controllers.GetBranchManagerFloatRequests)
	r.GET("/agent-admin-float-requests", controllers.GetAgentAdminFloatRequests)
	// r.GET("/till-operator-float-requests", controllers.GetTillOperatorFloatRequests)
	r.GET("/branch-manager-float-requests/:refNumber", controllers.GetBranchManagerFloatRequest)
	r.GET("/agent-admin-float-requests/:refNumber", controllers.GetAgentAdminFloatRequest)

	r.GET("/templates/preview/:refNumber", controllers.PreviewTemplate)
	r.GET("/documents/preview/:refNumber", controllers.PreviewDocument)

	r.DELETE("/templates/:refNumber", controllers.DeleteTemplate)
	r.DELETE("/documents/:refNumber", controllers.DeleteDocument)
	r.DELETE("/clear-logs", controllers.DeleteAllLogs)

	//endpoint to log the html before it turns to pdf
	r.POST("/htmlbeforepdf", controllers.HtmlBeforePDF)
	r.Run()
}
