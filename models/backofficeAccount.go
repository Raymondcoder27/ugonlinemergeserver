package models

// BackofficeAccount represents a user who manages back-office operations.
type BackofficeAccount struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"firstName" gorm:"unique;"`
	// MiddleName  string `json:"middleName" gorm:""`
	LastName string `json:"lastName" gorm:""`
	Phone    string `json:"phone" gorm:"unique;"`
	Email    string `json:"email" gorm:"unique;"`
	Role     string `json:"role" gorm:""`   // e.g., "Administrator", "Manager"
	Till     string `json:"till" gorm:""`   // e.g., "Till 1"
	Status   string `json:"status" gorm:""` // e.g., "Active", "Inactive"
}

type AllocateBranchManager struct {
	BranchID  uint `json:"branchId" gorm:""`
	ManagerID uint `json:"managerId" gorm:""`
}

type BranchManagers struct {
	BranchID  uint              `json:"branchId" gorm:""`
	ManagerID uint              `json:"managerId" gorm:""`
	Branch    Branch            `json:"branch" gorm:"foreignKey:BranchID"`
	Manager   BackofficeAccount `json:"manager" gorm:"foreignKey:ManagerID"`
	firstName string            `json:"firstName" gorm:""`
	lastName  string            `json:"lastName" gorm:""`
	Email     string            `json:"email" gorm:""`
	Phone     string            `json:"phone" gorm:""`
}
