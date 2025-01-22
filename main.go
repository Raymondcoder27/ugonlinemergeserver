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
	// r.GET("/logs", controllers.AutodocsLogs)
	// r.GET("/daterange-metrics", controllers.GetRangeMetrics)
	// r.GET("failed-generations", controllers.GetFailedGenerations)

	r.GET("/templates/preview/:refNumber", controllers.PreviewTemplate)
	r.GET("/documents/preview/:refNumber", controllers.PreviewDocument)

	r.DELETE("/templates/:refNumber", controllers.DeleteTemplate)
	r.DELETE("/documents/:refNumber", controllers.DeleteDocument)
	r.DELETE("/clear-logs", controllers.DeleteAllLogs)

	//endpoint to log the html before it turns to pdf
	r.POST("/htmlbeforepdf", controllers.HtmlBeforePDF)
	r.Run()
}
