package models

// BackofficeAccount represents a user who manages back-office operations.
type BackofficeAccount struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"firstname" gorm:"unique;"`
	// MiddleName  string `json:"middleName" gorm:""`
	LastName    string `json:"lastName" gorm:""`
	PhoneNumber string `json:"phone" gorm:"unique;"`
	Email       string `json:"email" gorm:"unique;"`
	Role        string `json:"role" gorm:""`   // e.g., "Administrator", "Manager"
	Till        string `json:"till" gorm:""`   // e.g., "Till 1"
	Status      string `json:"status" gorm:""` // e.g., "Active", "Inactive"
}
