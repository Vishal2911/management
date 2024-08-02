package model

import (
	"time"

	"github.com/google/uuid"
)

// School struct
type School struct {
	ID          uuid.UUID `gorm:"primarykey" json:"id"`
	Active      bool      `json:"active"  binding:"required" example:"true"`
	CreatedBy   string    `json:"created_by" binding:"required"  example:"vishal"`
	Domain      string    `json:"domain" binding:"required" gorm:"not null" example:"slrtce"`
	CreatedAt   time.Time `json:"created_at" `
	UpdatedBy   string    `json:"updated_by" `
	UpdatedAt   time.Time `json:"updated_at" `
	DeletedBy   string    `json:"deleted_by" `
	DeletedAt   time.Time `json:"deleted_at" `
	PrincipleID string    `json:"principle_id"  `
	AdminID     string    `json:"admin_id"   `
	DirectorID  uuid.UUID `json:"director_id"  binding:"required"  example:"e9b2a0e3-e86c-4386-8162-ae441ffa28c8"`
	HostelID    string    `json:"hostel_id"  `
	Lane        string    `json:"lane" example:"1234 Elm St"`
	Village     string    `json:"village" example:"Springfield"`
	City        string    `json:"city" binding:"required" example:"Metropolis"`
	District    string    `json:"district" binding:"required" example:"Central"`
	Pincode     int       `json:"pincode" binding:"required" example:"123456"`
	State       string    `json:"state" binding:"required" example:"NY"`
	BoardType   string    `json:"board_type" gorm:"not null"`
	ClassUpTo   int       `json:"class_up_to" gorm:"not null"`
}
