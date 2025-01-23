package models

import "time"

type FloatRequest struct {
	ID uint `gorm:"primaryKey"`
	// Amount    float64
	Amount    int       `json:"amount" gorm:"not null"`
	Status    string    `json:"status" gorm:"not null"` // e.g., "Allocated", "Pending", "Failed"
	CreatedAt time.Time `json:"createdAt"`
}
