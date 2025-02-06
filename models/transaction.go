package models

import "time"

type Transaction struct {
	ID        string `gorm:"primaryKey"`
	Amount    float64
	CreatedAt time.Time
}
