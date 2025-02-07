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
			tillOperator.GET("/float-ledgers", controllers.GetTillOperatorFloatLedger)
			tillOperator.PUT("/float-requests/:requestId", controllers.UpdateTillOperatorFloatRequest)
			tillOperator.PUT("/float-ledgers/:requestId", controllers.UpdateTillOperatorFloatLedger)
			// Add more Till Operator-specific routes here as needed
			//{{host}}/till-operator2-float-ledgers
			// tillOperator.GET("/float-requests/:refNumber", controllers.GetTillOperatorFloatRequest)
			// tillOperator.GET("/float-ledgers", controllers.GetTillOperatorFloatLedgers)
			tillOperator.POST("/add-float-ledger-record", controllers.AddTillOperatorFloatLeger)
			// tillOperator.GET("/service-requests", controllers.GetTillOperatorServiceRequests)
			// tillOperator.GET("/service-requests/:refNumber", controllers.GetTillOperatorServiceRequest)
			// //get till operator float ledgers
			// tillOperator.GET("/float-ledgers/:refNumber", controllers.GetTillOperatorFloatLedger)
			// tillOperator.GET("/float-ledgers", controllers.GetTillOperatorFloatLedgers)
			// tillOperator.PUT("/float-ledgers/:refNumber", controllers.UpdateTillOperatorFloatLedger)
		}

		// Branch Manager Dashboard
		branchManager := authGroup.Group("/branch-manager")
		{
			branchManager.POST("/request-float", controllers.BranchManagerRequestFloat)
			branchManager.POST("/add-float-ledger-record", controllers.AddBranchManagerFloatLedger)
			// branchManager.POST("/approve-float-request", controllers.BranchManagerApproveFloatRequest)
			branchManager.PUT("/update-float-request/:id", controllers.BranchManagerUpdateFloatRequest)
			// branchManager.PUT("/approve-float-request/{id}", controllers.BranchManagerApproveFloatRequest)
			branchManager.PUT("/update-float-ledger/:id", controllers.BranchManagerUpdateFloatLedger)
			branchManager.GET("/float-requests", controllers.GetBranchManagerFloatRequests)
			// branchManager.GET("/float-requests", controllers.GetTillOperatorFloatRequests)
			branchManager.GET("/float-requests/:refNumber", controllers.GetBranchManagerFloatRequest)
			// branchManager.GET("/float-ledgers", controllers.GetBranchManagerFloatLedgers)
			// branchManager.GET("/float-ledgers/:refNumber", controllers.GetBranchManagerFloatLedger)
			// branchManager.PUT("/float-ledgers/:refNumber", controllers.UpdateBranchManagerFloatLedger)
			// branchManager.GET("/service-requests", controllers.GetBranchManagerServiceRequests)
			// branchManager.GET("/service-requests/:refNumber", controllers.GetBranchManagerServiceRequest)
			// branchManager.DELETE("/till/:refNumber", controllers.DeleteTill)

		}

		// Agent Admin Dashboard
		agentAdmin := authGroup.Group("/agent-admin")
		{
			agentAdmin.GET("/services", controllers.GetAgentAdminFloatRequests)
			agentAdmin.POST("/create-branch", controllers.CreateBranch)
			agentAdmin.GET("/back-office-accounts", controllers.GetBackOfficeAccounts)
			agentAdmin.POST("/assign-branch-manager", controllers.AllocateBranchManager)
			agentAdmin.POST("/create-branch-manager-account", controllers.CreateBranchManagerAccount)
			agentAdmin.POST("/create-back-office-account", controllers.CreateBackOfficeAccount)
			agentAdmin.POST("/approve-float", controllers.AgentAdminApproveFloat)
			agentAdmin.GET("/float-requests", controllers.GetAgentAdminFloatRequests)
			agentAdmin.GET("/float-requests/:refNumber", controllers.GetAgentAdminFloatRequest)
			agentAdmin.GET("/branches", controllers.GetBranches)
			agentAdmin.GET("/branch-manager-accounts", controllers.GetBranchManagerAccounts)
			agentAdmin.PUT("/update-float-request/:id", controllers.AgentAdminUpdateFloatRequest)
			// branchManager.PUT("/approve-float-request/{id}", controllers.BranchManagerApproveFloatRequest)
			agentAdmin.PUT("/update-float-ledger/:id", controllers.AgentAdminUpdateFloatLedger)
		}
	}

	// Start the server
	r.Run(":8080") // You can customize the port as needed
}
