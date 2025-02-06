package models

import "time"

type FloatLedger struct {
	ID        string `gorm:"primaryKey"`
	Amount    float64
	CreatedAt time.Time
}
