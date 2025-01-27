package models

import "time"

type TillOperatorFloatRequest struct {
	// ID uint `gorm:"primaryKey"`
	// Amount    float64
	Amount    int       `json:"amount" gorm:"not null"`
	Status    string    `json:"status" gorm:"not null"` // e.g., "Allocated", "Pending", "Failed"
	CreatedAt time.Time `json:"createdAt"`
	Till      string    `json:"till" gorm:"not null"` // e.g., "Till 1"
}

type BranchManagerFloatRequest struct {
	// ID uint `gorm:"primaryKey"`
	// Amount    float64
	Amount    int       `json:"amount" gorm:"not null"`
	Status    string    `json:"status" gorm:"not null"` // e.g., "Allocated", "Pending", "Failed"
	CreatedAt time.Time `json:"createdAt"`
	Branch    string    `json:"branch" gorm:"not null"` // e.g., "Branch 1"
}

type AdminAgentFloatRequest struct {
	// ID uint `gorm:"primaryKey"`
	// Amount    float64
	Amount    int       `json:"amount" gorm:"not null"`
	Status    string    `json:"status" gorm:"not null"` // e.g., "Allocated", "Pending", "Failed"
	CreatedAt time.Time `json:"createdAt"`
	Agent     string    `json:"agent" gorm:"not null"` // e.g., "Agent 1"
}
