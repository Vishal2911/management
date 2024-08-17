package model

import (
	"github.com/google/uuid"
	"time"
)

// Publisher struct
type Publisher struct {
	ID        uuid.UUID `gorm:"primarykey" json:"id"`
	Name      string    `json:"name" binding:"required" example:"some_publisher"`
	Active    bool      `json:"active"  binding:"required" example:"true"`
	CreatedBy string    `json:"created_by" binding:"required"  example:"vishal"`
	DeletedBy string    `json:"deleted_by" `
	UpdatedBy string    `json:"updated_by" `
	CreatedAt time.Time `json:"created_at" `
	UpdatedAt time.Time `json:"updated_at" `
	DeletedAt time.Time `json:"deleted_at" `
}
