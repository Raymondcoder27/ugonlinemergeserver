package models

import "time"

// Transaction represents a financial transaction in the branch management system.
type Transaction struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	TrackingNumber string    `json:"trackingNumber" gorm:"not null"`
	Service        string    `json:"service" gorm:"not null"`
	Provider       string    `json:"provider" gorm:"not null"`
	Till           string    `json:"till" gorm:"not null"`
	Fee            float64   `json:"fee" gorm:"not null"`
	Date           time.Time `json:"date" gorm:"not null"`
	Status         string    `json:"status" gorm:"not null"` // e.g., "success", "failed", "pending"
}

// FloatRequest represents a request for float allocation in the branch.
type FloatRequest struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	RequestDate time.Time `json:"requestDate" gorm:"not null"`
	Amount      float64   `json:"amount" gorm:"not null"`
	Status      string    `json:"status" gorm:"not null"` // e.g., "pending", "approved", "rejected"
	Till        string    `json:"till" gorm:"not null"`
	ApprovedBy  *string   `json:"approvedBy"`                // Null if not approved
	Requester   string    `json:"requester" gorm:"not null"` // Name of the person requesting the float
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
}

// FloatLedger represents the ledger for tracking float allocations and usage.
type FloatLedger struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Date        time.Time `json:"date" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"` // e.g., "Recharge", "Till X"
	Amount      float64   `json:"amount" gorm:"not null"`
	Balance     float64   `json:"balance" gorm:"not null"`
}

// BackofficeUser represents a user who manages back-office operations.
type BackofficeUser struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null"`
	FullName string `json:"fullName" gorm:"not null"`
	Role     string `json:"role" gorm:"not null"`   // e.g., "Administrator", "Manager"
	Till     string `json:"till" gorm:"not null"`   // e.g., "Till 1"
	Status   string `json:"status" gorm:"not null"` // e.g., "Active", "Inactive"
}

// TillOperator represents a user managing float at specific tills.
type TillOperator struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null"`
	FullName string `json:"fullName" gorm:"not null"`
	Role     string `json:"role" gorm:"not null"`   // e.g., "Manager"
	Till     string `json:"till" gorm:"not null"`   // e.g., "Till 1"
	Status   string `json:"status" gorm:"not null"` // e.g., "Active", "Inactive"
}

// FloatAllocation represents the allocation of float to tills.
type FloatAllocation struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	DateAssigned time.Time `json:"dateAssigned" gorm:"not null"`
	Amount       float64   `json:"amount" gorm:"not null"`
	Status       string    `json:"status" gorm:"not null"` // e.g., "Allocated", "Pending", "Failed"
	Till         string    `json:"till" gorm:"not null"`   // e.g., "Till 1"
}
