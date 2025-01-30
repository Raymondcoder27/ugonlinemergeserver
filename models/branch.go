package models

// TillOperator represents a user managing float at specific tills.
type Branch struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique;not null"` // e.g., "Till 1"
	// Description string `json:"description" gorm:"not null"` // e.g., "Active", "Inactive"
}

//id should be
