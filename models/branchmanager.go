package models

// TillOperator represents a user managing float at specific tills.
type TillOperator struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null"`
	FullName string `json:"fullName" gorm:"not null"`
	Role     string `json:"role" gorm:"not null"`   // e.g., "Manager"
	Till     string `json:"till" gorm:"not null"`   // e.g., "Till 1"
	Status   string `json:"status" gorm:"not null"` // e.g., "Active", "Inactive"
}
