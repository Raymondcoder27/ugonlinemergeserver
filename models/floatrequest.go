package models

import "time"

type FloatRequest struct {
	ID        uint `gorm:"primaryKey"`
	Amount    float64
	Status    string
	CreatedAt time.Time
}
