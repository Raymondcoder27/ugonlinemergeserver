package models

import "time"

// Service represents a service offered by the Till Operator system.
type Service struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Provider    string    `json:"provider" gorm:"not null"` // The provider of the service
	Status      string    `json:"status" gorm:"not null"`   // e.g., "Active", "Inactive"
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

// ServiceSpecification represents detailed specifications of a service.
type ServiceSpecification struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	ServiceID string    `json:"serviceId" gorm:"not null"` // Foreign key to Service
	Key       string    `json:"key" gorm:"not null"`
	Value     string    `json:"value" gorm:"not null"`
	Status    string    `json:"status" gorm:"not null"` // e.g., "Enabled", "Disabled"
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

// ServiceResponse represents a response from service-related operations.
type ServiceResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

// CreateServiceRequest represents the payload for creating a service.
type CreateServiceRequest struct {
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Provider    string `json:"provider" gorm:"not null"`
}

// UpdateServiceRequest represents the payload for updating a service.
type UpdateServiceRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// CreateServiceSpecRequest represents the payload for creating a service specification.
type CreateServiceSpecRequest struct {
	ServiceID string `json:"serviceId" gorm:"not null"`
	Key       string `json:"key" gorm:"not null"`
	Value     string `json:"value" gorm:"not null"`
}

// UpdateServiceSpecRequest represents the payload for updating a service specification.
type UpdateServiceSpecRequest struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Status string `json:"status"`
}

type TillOperatorFloatLedger struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Description string    `json:"description" gorm:"not null"`
	Amount      float64   `json:"amount" gorm:"not null"` // The provider of the service
	Balance     float64   `json:"balance" gorm:"not null"`
	Status      string    `json:"status" gorm:"not null"` // e.g., "Active", "Inactive"
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

type AssignTillOperator struct {
	TIllID         string `json:"tillId" gorm:""`
	TillOperatorID string `json:"tillOperatorId" gorm:""`
}
