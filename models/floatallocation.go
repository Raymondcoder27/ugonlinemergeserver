package models

import "time"

// FloatAllocation represents the allocation of float to tills.
type FloatAllocation struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	DateAssigned time.Time `json:"dateAssigned" gorm:"not null"`
	Amount       float64   `json:"amount" gorm:"not null"`
	Status       string    `json:"status" gorm:"not null"` // e.g., "Allocated", "Pending", "Failed"
	Till         string    `json:"till" gorm:"not null"`   // e.g., "Till 1"
}
