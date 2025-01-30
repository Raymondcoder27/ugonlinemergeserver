package models

// BackofficeAccount represents a user who manages back-office operations.
type BackofficeAccount struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	FirstName   string `json:"firstname" gorm:"unique;not null"`
	MiddleName  string `json:"middleName" gorm:"not null"`
	LastName    string `json:"lastName" gorm:"not null"`
	PhoneNumber string `json:"phoneNumber" gorm:"unique;not null"`
	Email       string `json:"email" gorm:"unique;not null"`
	Role        string `json:"role" gorm:"not null"`   // e.g., "Administrator", "Manager"
	Till        string `json:"till" gorm:"not null"`   // e.g., "Till 1"
	Status      string `json:"status" gorm:"not null"` // e.g., "Active", "Inactive"
}
