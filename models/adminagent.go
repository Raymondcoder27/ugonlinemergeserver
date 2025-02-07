package models

import "time"

// BranchManager represents a manager assigned to a specific branch.
type BranchManager struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null"`
	FullName string `json:"fullName" gorm:"not null"`
	Role     string `json:"role" gorm:"not null"`   // e.g., "Manager"
	Branch   string `json:"branch" gorm:"not null"` // e.g., "Branch 1"
	Status   string `json:"status" gorm:"not null"` // e.g., "Active", "Inactive"
}

type AgentAdminFloatLedger struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Description string    `json:"description" gorm:"not null"`
	Amount      float64   `json:"amount" gorm:"not null"` // The provider of the service
	Balance     float64   `json:"balance" gorm:"not null"`
	Status      string    `json:"status" gorm:"not null"` // e.g., "Active", "Inactive"
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

// FloatAllocation represents the allocation of float to branches.
// type FloatAllocation struct {
// 	ID           string      `json:"id" gorm:"primaryKey"`
// 	DateAssigned time.Time `json:"dateAssigned" gorm:"not null"`
// 	Amount       float64   `json:"amount" gorm:"not null"`
// 	Status       string    `json:"status" gorm:"not null"` // e.g., "Allocated", "Pending", "Failed"
// 	Branch       string    `json:"branch" gorm:"not null"` // e.g., "Branch 1"
// }
