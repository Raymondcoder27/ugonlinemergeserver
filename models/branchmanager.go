package models

import "time"

// TillOperator represents a user managing float at specific tills.
type TillOperator struct {
	ID string `json:"id" gorm:"primaryKey"`
	// Username string `json:"username" gorm:"unique;not null"`
	FirstName string `json:"firstName" gorm:"not null"`
	LastName  string `json:"lastName" gorm:"not null"`
	// FullName  string `json:"fullName" gorm:"not null"`
	Email  string `json:"email" gorm:"not null"`
	Phone  string `json:"phone" gorm:"not null"`
	Role   string `json:"role" gorm:"not null"`   // e.g., "Manager"
	Till   string `json:"till" gorm:"not null"`   // e.g., "Till 1"
	Status string `json:"status" gorm:"not null"` // e.g., "Active", "Inactive"
}

type BranchManagerFloatLedger struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Description string    `json:"description" gorm:"not null"`
	Amount      float64   `json:"amount" gorm:"not null"` // The provider of the service
	Balance     float64   `json:"balance" gorm:"not null"`
	Status      string    `json:"status" gorm:"not null"` // e.g., "Active", "Inactive"
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

// BackofficeAccount represents a user who manages back-office operations.
type BackofficeAccount struct {
	ID        string `json:"id" gorm:"primaryKey"`
	FirstName string `json:"firstName" gorm:"unique;"`
	// MiddleName  string `json:"middleName" gorm:""`
	LastName string `json:"lastName" gorm:""`
	Phone    string `json:"phone" gorm:"unique;"`
	Email    string `json:"email" gorm:"unique;"`
	Role     string `json:"role" gorm:""`   // e.g., "Administrator", "Manager"
	Till     string `json:"till" gorm:""`   // e.g., "Till 1"
	Status   string `json:"status" gorm:""` // e.g., "Active", "Inactive"
}

// BackofficeAccount represents a user who manages back-office operations.
type BranchBackofficeAccount struct {
	ID        string `json:"id" gorm:"primaryKey"`
	FirstName string `json:"firstName" gorm:"unique;"`
	// MiddleName  string `json:"middleName" gorm:""`
	LastName string `json:"lastName" gorm:""`
	Phone    string `json:"phone" gorm:"unique;"`
	Email    string `json:"email" gorm:"unique;"`
	Role     string `json:"role" gorm:""`   // e.g., "Administrator", "Manager"
	Branch   string `json:"branch" gorm:""` // e.g., "Till 1"
	Status   string `json:"status" gorm:""` // e.g., "Active", "Inactive"
}

type AllocateBranchManager struct {
	BranchID  string `json:"branchId" gorm:""`
	ManagerID string `json:"managerId" gorm:""`
}

type BranchManagers struct {
	ID string `json:"id" gorm:"primaryKey"`
	// BranchID  string `json:"branchId" gorm:""`
	// ManagerID string `json:"managerId" gorm:""`
	// Branch    Branch            `json:"branch" gorm:"foreignKey:BranchID"`
	// Manager   BackofficeAccount `json:"manager" gorm:"foreignKey:ManagerID"`
	FirstName string `json:"firstName" gorm:""`
	LastName  string `json:"lastName" gorm:""`
	Email     string `json:"email" gorm:""`
	Phone     string `json:"phone" gorm:""`
	Branch    string `json:"branch" gorm:""`
}
