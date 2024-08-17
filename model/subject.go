package model

import (
	"time"

	"github.com/google/uuid"
)

// Subject struct
type Subject struct {
	ID        uuid.UUID `gorm:"primarykey" json:"id"`
	SchoolID  uuid.UUID `binding:"required" json:"school_id"`
	BookID    uuid.UUID `binding:"required" json:"book_id"`
	Name      string    `json:"name" `
	Active    bool      `json:"active"  binding:"required" example:"true"`
	CreatedBy string    `json:"created_by" binding:"required"  example:"vishal"`
	DeletedBy string    `json:"deleted_by" `
	UpdatedBy string    `json:"updated_by" `
	CreatedAt time.Time `json:"created_at" `
	UpdatedAt time.Time `json:"updated_at" `
	DeletedAt time.Time `json:"deleted_at" `
}
