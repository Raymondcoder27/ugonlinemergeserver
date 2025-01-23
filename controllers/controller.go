package controllers

import (
	"example/pdfgenerator/models" // Replace with your actual models package
	"net/http"

	"github.com/gin-gonic/gin"
)

// TillOperatorRequestFloat handles the request for float allocation by Till Operator.
func TillOperatorRequestFloat(c *gin.Context) {
	var request models.FloatRequest

	// Bind JSON request to the FloatRequest model
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set default status to "Pending"
	request.Status = "pending"

	// Save request to database
	if err := db.Create(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create float request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Float request created successfully", "data": request})
}

// TillOperatorServiceRequest handles the service request by the Till Operator.
func TillOperatorServiceRequest(c *gin.Context) {
	var request models.ServiceRequest

	// Bind JSON request to the ServiceRequest model
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the service request
	if err := db.Create(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create service request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Service request created successfully", "data": request})
}

// BranchManagerRequestFloat handles the request for float allocation by Branch Manager.
func BranchManagerRequestFloat(c *gin.Context) {
	var request models.FloatRequest

	// Bind JSON request to the FloatRequest model
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set default status to "pending"
	request.Status = "pending"

	// Save request to database
	if err := db.Create(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create float request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Float request created successfully", "data": request})
}

// BranchManagerApproveFloat handles the approval of a float request by Branch Manager.
func BranchManagerApproveFloat(c *gin.Context) {
	refNumber := c.Param("refNumber")

	var request models.FloatRequest

	// Find the float request by refNumber
	if err := db.Where("ref_number = ?", refNumber).First(&request).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Float request not found"})
		return
	}

	// Approve the float request
	request.Status = "approved"
	if err := db.Save(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to approve float request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Float request approved", "data": request})
}

// GetBranchManagerFloatRequests fetches all float requests for the Branch Manager.
func GetBranchManagerFloatRequests(c *gin.Context) {
	var requests []models.FloatRequest

	// Fetch all float requests for the branch manager
	if err := db.Where("status = ?", "pending").Find(&requests).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch float requests"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": requests})
}

// GetBranchManagerFloatRequest fetches a specific float request by reference number.
func GetBranchManagerFloatRequest(c *gin.Context) {
	refNumber := c.Param("refNumber")

	var request models.FloatRequest

	// Fetch the specific float request by refNumber
	if err := db.Where("ref_number = ?", refNumber).First(&request).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Float request not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": request})
}

// AgentAdminApproveFloat handles the approval of a float request by the Agent Admin.
func AgentAdminApproveFloat(c *gin.Context) {
	refNumber := c.Param("refNumber")

	var request models.FloatRequest

	// Find the float request by refNumber
	if err := db.Where("ref_number = ?", refNumber).First(&request).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Float request not found"})
		return
	}

	// Approve the float request
	request.Status = "approved"
	if err := db.Save(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to approve float request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Float request approved", "data": request})
}

// GetAgentAdminFloatRequests fetches all float requests for the Agent Admin.
func GetAgentAdminFloatRequests(c *gin.Context) {
	var requests []models.FloatRequest

	// Fetch all float requests for the agent admin
	if err := db.Where("status = ?", "pending").Find(&requests).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch float requests"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": requests})
}

// GetAgentAdminFloatRequest fetches a specific float request by reference number.
func GetAgentAdminFloatRequest(c *gin.Context) {
	refNumber := c.Param("refNumber")

	var request models.FloatRequest

	// Fetch the specific float request by refNumber
	if err := db.Where("ref_number = ?", refNumber).First(&request).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Float request not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": request})
}
