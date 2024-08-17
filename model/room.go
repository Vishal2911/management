package model

import (
	"time"

	"github.com/google/uuid"
)

// Room struct
type Room struct {
	ID          uuid.UUID `gorm:"primarykey" json:"id"`
	SchoolID    uuid.UUID `binding:"required" json:"school_id"`
	Active      bool      `json:"active"  binding:"required" example:"true"`
	Width       int       ` binding:"required" json:"width" `
	Hight       int       ` binding:"required" json:"hight" `
	Lenght      int       ` binding:"required" json:"length" `
	RoomNumber  int       `json:"room_number"  `
	FloorNumber int       `json:"floor_number"  `
	RoomType    string    `json:"room_type"  `
	Name        string    `json:"name"  `
	AC          bool      `json:"ac"  `
	Stared      bool      `json:"stared"  `
	CreatedBy   string    `json:"created_by" binding:"required"  example:"vishal"`
	DeletedBy   string    `json:"deleted_by" `
	UpdatedBy   string    `json:"updated_by" `
	CreatedAt   time.Time `json:"created_at" `
	UpdatedAt   time.Time `json:"updated_at" `
	DeletedAt   time.Time `json:"deleted_at" `
}
