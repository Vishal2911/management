package model

import (
	"github.com/google/uuid"
	"time"
)

// Author struct
type Author struct {
	ID         uuid.UUID `gorm:"primarykey" json:"id"`
	FirstName  string    `json:"first_name" binding:"required" example:"John"`
	MiddleName string    `json:"middle_name" binding:"required"  example:"Doe"`
	LastName   string    `json:"last_name" binding:"required" example:"Smith"`
	Lane       string    `json:"lane" example:"1234 Elm St"`
	Village    string    `json:"village" example:"Springfield"`
	City       string    `json:"city" binding:"required" example:"Metropolis"`
	District   string    `json:"district" binding:"required" example:"Central"`
	Pincode    int       `json:"pincode" binding:"required" example:"123456"`
	State      string    `json:"state" binding:"required" example:"NY"`
	Active     bool      `json:"active"  binding:"required" example:"true"`
	CreatedBy  string    `json:"created_by" binding:"required"  example:"vishal"`
	DeletedBy  string    `json:"deleted_by" `
	UpdatedBy  string    `json:"updated_by" `
	CreatedAt  time.Time `json:"created_at" `
	UpdatedAt  time.Time `json:"updated_at" `
	DeletedAt  time.Time `json:"deleted_at" `
}
