package controllers

import (
	"net/http"
	"strconv"

	"example.com/ugonlinemergeserver/initializers"
	"example.com/ugonlinemergeserver/models" // Replace with your actual models package

	"github.com/gin-gonic/gin"
)

// TillOperatorRequestFloat handles the request for float allocation by Till Operator.
func TillOperatorRequestFloat(c *gin.Context) {
	var request models.TillOperatorFloatRequest

	// Bind JSON request to the FloatRequest model
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set default status to "Pending"
	request.Status = "pending"
	request.Till = "Till 1"

	// Save request to database
	if err := initializers.DB.Create(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create float request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Float request created successfully", "data": request})
}

// TillOperatorServiceRequest handles the service request by the Till Operator.
func TillOperatorServiceRequest(c *gin.Context) {
	var request models.CreateServiceRequest

	// Bind JSON request to the ServiceRequest model
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the service request
	if err := initializers.DB.Create(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create service request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Service request created successfully", "data": request})
}

// tillOperator.GET("/float-requests", controllers.GetTillOperatorFloatRequests)

// GetTillOperatorFloatRequests fetches all float requests for the Till Operator.
// func GetTillOperatorFloatRequests(c *gin.Context) {
// 	var requests []models.FloatRequest

// 	// Fetch all float requests for the till operator
// 	if err := initializers.DB.Where("status = ?", "pending").Find(&requests).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch float requests"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": requests})
// }

// BranchManagerRequestFloat handles the request for float allocation by Branch Manager.
func BranchManagerRequestFloat(c *gin.Context) {
	var request models.BranchManagerFloatRequest

	// Bind JSON request to the FloatRequest model
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set default status to "pending"
	request.Status = "pending"

	// Save request to database
	if err := initializers.DB.Create(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create float request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Float request created successfully", "data": request})
}

// BranchManagerApproveFloat handles the approval of a float request by Branch Manager.
// func BranchManagerApproveFloatRequest(c *gin.Context) {
// 	refNumber := c.Param("id")

// 	var request models.TillOperatorFloatRequest

// 	// Find the float request by refNumber
// 	if err := initializers.DB.Where("ID = ?", refNumber).First(&request).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Float request not found"})
// 		return
// 	}

// 	// Approve the float request
// 	request.Status = "approved"
// 	if err := initializers.DB.Save(&request).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to approve float request"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Float request approved", "data": request})
// }

// BranchManagerApproveFloatRequest handles the approval of a float request by Branch Manager.
func BranchManagerApproveFloatRequest(c *gin.Context) {
	// Extract the "id" parameter from the URL
	requestId := c.Param("id")

	// Validate and convert the ID to an integer
	id, err := strconv.ParseInt(requestId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format. ID must be a number"})
		return
	}

	var request models.TillOperatorFloatRequest

	// Find the float request by ID
	if err := initializers.DB.Where("ID = ?", id).First(&request).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Float request not found"})
		return
	}

	// Approve the float request
	request.Status = "approved"
	if err := initializers.DB.Save(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to approve float request"})
		return
	}

	//save updated record to floatrequests db
	//inserting delete request into logs table
	// if err := initializers.DB.Create(&models.TillOperatorFloatRequest{
	// 	Amount:    request.Amount,
	// 	CreatedAt: request.CreatedAt,
	// 	Till:      request.Till,
	// 	Status:    request.Status,
	// }).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving document metadata in database: " + err.Error()})
	// 	return
	// }

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "Float request approved", "data": request})
}

func CreateBranch(c *gin.Context) {
	var request models.CreateBranch

	// Bind JSON request to the FloatRequest model
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set default status to "Pending"
	// request.Status = "pending"
	// request.Till = "Till 1"

	// Save request to database
	if err := initializers.DB.Create(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create float request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Branch created successfully", "data": request})
}

func GetBackOfficeAccounts(c *gin.Context) {

}

// agentAdmin.POST("/create-branch", controllers.CreateBranch)

// GetBranchManagerFloatRequests fetches all float requests for the Branch Manager.
func GetBranchManagerFloatRequests(c *gin.Context) {
	var requests []models.TillOperatorFloatRequest

	// Fetch all float requests for the branch manager
	if err := initializers.DB.Where("status = ?", "pending").Find(&requests).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch float requests"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": requests})
}

func GetTillOperatorFloatRequests(c *gin.Context) {
	var requests []models.TillOperatorFloatRequest

	// Fetch all float requests for the branch manager
	if err := initializers.DB.Where("status = ?", "pending").Find(&requests).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch float requests"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": requests})
}

// GetBranchManagerFloatRequest fetches a specific float request by reference number.
func GetBranchManagerFloatRequest(c *gin.Context) {
	refNumber := c.Param("refNumber")

	var request models.BranchManagerFloatRequest

	// Fetch the specific float request by refNumber
	if err := initializers.DB.Where("ID = ?", refNumber).First(&request).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Float request not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": request})
}

// AgentAdminApproveFloat handles the approval of a float request by the Agent Admin.
func AgentAdminApproveFloat(c *gin.Context) {
	refNumber := c.Param("refNumber")

	var request models.BranchManagerFloatRequest

	// Find the float request by refNumber
	if err := initializers.DB.Where("ID = ?", refNumber).First(&request).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Float request not found"})
		return
	}

	// Approve the float request
	request.Status = "approved"
	if err := initializers.DB.Save(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to approve float request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Float request approved", "data": request})
}

// GetAgentAdminFloatRequests fetches all float requests for the Agent Admin.
func GetAgentAdminFloatRequests(c *gin.Context) {
	var requests []models.AdminAgentFloatRequest

	// Fetch all float requests for the agent admin
	if err := initializers.DB.Where("status = ?", "pending").Find(&requests).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch float requests"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": requests})
}

// GetAgentAdminFloatRequest fetches a specific float request by reference number.
func GetAgentAdminFloatRequest(c *gin.Context) {
	refNumber := c.Param("refNumber")

	var request models.AdminAgentFloatRequest

	// Fetch the specific float request by refNumber
	if err := initializers.DB.Where("ID = ?", refNumber).First(&request).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Float request not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": request})
}
