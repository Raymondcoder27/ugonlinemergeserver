package models

import "time"

// TillOperator represents a user managing float at specific tills.
type TillOperator struct {
	ID string `json:"id" gorm:"primaryKey"`
	// Username string `json:"username" gorm:"unique;not null"`
	FirstName string `json:"firstName" gorm:"not null"`
	LastName  string `json:"lastName" gorm:"not null"`
	FullName  string `json:"fullName" gorm:"not null"`
	Role      string `json:"role" gorm:"not null"`   // e.g., "Manager"
	Till      string `json:"till" gorm:"not null"`   // e.g., "Till 1"
	Status    string `json:"status" gorm:"not null"` // e.g., "Active", "Inactive"
}

type BranchManagerFloatLedger struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Description string    `json:"description" gorm:"not null"`
	Amount      float64   `json:"amount" gorm:"not null"` // The provider of the service
	Balance     float64   `json:"balance" gorm:"not null"`
	Status      string    `json:"status" gorm:"not null"` // e.g., "Active", "Inactive"
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
