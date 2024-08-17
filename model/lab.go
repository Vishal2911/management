package model

import (
	"github.com/google/uuid"
	"time"
)

// Lab struct
type Lab struct {
	ID                 uuid.UUID `gorm:"primarykey" json:"id"`
	LabName            string    `binding:"required" json:"lab_name"`
	SchoolID           uuid.UUID `binding:"required" json:"school_id"`
	RoomID             uuid.UUID `binding:"required" json:"room_id"`
	Active             bool      `json:"active"  binding:"required" example:"true"`
	NumberOfEquipments int       `binding:"required" json:"number_of_equipments" `
	CreatedBy          string    `json:"created_by" binding:"required"  example:"vishal"`
	DeletedBy          string    `json:"deleted_by" `
	UpdatedBy          string    `json:"updated_by" `
	CreatedAt          time.Time `json:"created_at" `
	UpdatedAt          time.Time `json:"updated_at" `
	DeletedAt          time.Time `json:"deleted_at" `
	EquipmentType      string    `json:"equipment_type"  `
	LabAssistantID     uuid.UUID `json:"lab_assistant_id"`
}
