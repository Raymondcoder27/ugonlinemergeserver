package models

// BackofficeUser represents a user who manages back-office operations.
type BackofficeUser struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null"`
	FullName string `json:"fullName" gorm:"not null"`
	Role     string `json:"role" gorm:"not null"`   // e.g., "Administrator", "Manager"
	Till     string `json:"till" gorm:"not null"`   // e.g., "Till 1"
	Status   string `json:"status" gorm:"not null"` // e.g., "Active", "Inactive"
}
