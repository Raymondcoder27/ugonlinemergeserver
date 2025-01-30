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
			tillOperator.GET("/float-requests", controllers.GetTillOperatorFloatRequests)
			// Add more Till Operator-specific routes here as needed
		}

		// Branch Manager Dashboard
		branchManager := authGroup.Group("/branch-manager")
		{
			branchManager.POST("/request-float", controllers.BranchManagerRequestFloat)
			// branchManager.POST("/approve-float-request", controllers.BranchManagerApproveFloatRequest)
			branchManager.POST("/approve-float-request/:id", controllers.BranchManagerApproveFloatRequest)
			branchManager.GET("/float-requests", controllers.GetBranchManagerFloatRequests)
			// branchManager.GET("/float-requests", controllers.GetTillOperatorFloatRequests)
			branchManager.GET("/float-requests/:refNumber", controllers.GetBranchManagerFloatRequest)
		}

		// Agent Admin Dashboard
		agentAdmin := authGroup.Group("/agent-admin")
		{
			agentAdmin.GET("/services", controllers.GetAgentAdminFloatRequests)
			agentAdmin.POST("/create-branch", controllers.CreateBranch)
			agentAdmin.GET("/back-office-accounts", controllers.GetBackOfficeAccounts)
			agentAdmin.POST("/create-back-office-account", controllers.CreateBackOfficeAccount)
			agentAdmin.POST("/approve-float", controllers.AgentAdminApproveFloat)
			agentAdmin.GET("/float-requests", controllers.GetAgentAdminFloatRequests)
			agentAdmin.GET("/float-requests/:refNumber", controllers.GetAgentAdminFloatRequest)
			agentAdmin.GET("/branches", controllers.GetBranches)
			agentAdmin.GET("/branch-manager-accounts", controllers.GetBranchManagerAccounts)
		}
	}

	// Start the server
	r.Run(":8080") // You can customize the port as needed
}
