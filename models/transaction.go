package models

import "time"

type Transaction struct {
	ID        uint `gorm:"primaryKey"`
	Amount    float64
	CreatedAt time.Time
}
