package models

import "time"

// Transaction represents a financial transaction.
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

// FloatRequest represents a request for float allocation.
type FloatRequest struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	RequestDate time.Time `json:"requestDate" gorm:"not null"`
	Amount      float64   `json:"amount" gorm:"not null"`
	Status      string    `json:"status" gorm:"not null"` // e.g., "pending", "approved", "rejected"
	Branch      string    `json:"branch" gorm:"not null"`
	ApprovedBy  *string   `json:"approvedBy"` // Null if not approved
}

// FloatLedger represents the ledger for float allocations and usage.
type FloatLedger struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Date        time.Time `json:"date" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"` // e.g., "Recharge", "Branch X"
	Amount      float64   `json:"amount" gorm:"not null"`
	Balance     float64   `json:"balance" gorm:"not null"`
}

// BackofficeUser represents a user in the back office system.
type BackofficeUser struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null"`
	FullName string `json:"fullName" gorm:"not null"`
	Role     string `json:"role" gorm:"not null"`   // e.g., "Administrator", "Manager"
	Branch   string `json:"branch" gorm:"not null"` // e.g., "Branch 1"
	Status   string `json:"status" gorm:"not null"` // e.g., "Active", "Inactive"
}

// BranchManager represents a manager assigned to a specific branch.
type BranchManager struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null"`
	FullName string `json:"fullName" gorm:"not null"`
	Role     string `json:"role" gorm:"not null"`   // e.g., "Manager"
	Branch   string `json:"branch" gorm:"not null"` // e.g., "Branch 1"
	Status   string `json:"status" gorm:"not null"` // e.g., "Active", "Inactive"
}

// FloatAllocation represents the allocation of float to branches.
type FloatAllocation struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	DateAssigned time.Time `json:"dateAssigned" gorm:"not null"`
	Amount       float64   `json:"amount" gorm:"not null"`
	Status       string    `json:"status" gorm:"not null"` // e.g., "Allocated", "Pending", "Failed"
	Branch       string    `json:"branch" gorm:"not null"` // e.g., "Branch 1"
}
