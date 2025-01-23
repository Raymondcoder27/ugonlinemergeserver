package models

import "time"

type FloatLedger struct {
	ID        uint `gorm:"primaryKey"`
	Amount    float64
	CreatedAt time.Time
}
