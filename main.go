package main

import (
	"example.com/ugonlinemergeserver/controllers"
	"example.com/ugonlinemergeserver/initializers"

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

	// Define route groups for the 3 dashboards
	// Auth group for protected routes
	// authGroup := r.Group("/", middleware.RequireAuth)
	authGroup := r.Group("/")

	{
		// Till Operator Dashboard
		tillOperator := authGroup.Group("/till-operator")
		{
			tillOperator.POST("/request-float", controllers.TillOperatorRequestFloat)
			tillOperator.POST("/service-request", controllers.TillOperatorServiceRequest)
			// Add more Till Operator-specific routes here as needed
		}

		// Branch Manager Dashboard
		branchManager := authGroup.Group("/branch-manager")
		{
			branchManager.POST("/request-float", controllers.BranchManagerRequestFloat)
			branchManager.POST("/approve-float", controllers.BranchManagerApproveFloat)
			branchManager.GET("/float-requests", controllers.GetBranchManagerFloatRequests)
			branchManager.GET("/float-requests/:refNumber", controllers.GetBranchManagerFloatRequest)
		}

		// Agent Admin Dashboard
		agentAdmin := authGroup.Group("/agent-admin")
		{
			agentAdmin.POST("/approve-float", controllers.AgentAdminApproveFloat)
			agentAdmin.GET("/float-requests", controllers.GetAgentAdminFloatRequests)
			agentAdmin.GET("/float-requests/:refNumber", controllers.GetAgentAdminFloatRequest)
		}
	}

	// Start the server
	r.Run(":8080") // You can customize the port as needed
}
