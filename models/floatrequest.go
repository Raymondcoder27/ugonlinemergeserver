package models

import "time"

type FloatRequest struct {
	ID uint `gorm:"primaryKey"`
	// Amount    float64
	Amount    int
	Status    string
	CreatedAt time.Time
}
