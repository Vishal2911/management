package model

import (
	"time"
	"github.com/google/uuid"
)

// Class struct
type Class struct {
	ID              uuid.UUID `gorm:"primarykey" json:"id"`
	SchoolID        uuid.UUID `binding:"required" json:"school_id"`
	Active          bool      `json:"active"  binding:"required" example:"true"`
	NumberOfStudent int       ` binding:"required" json:"number_of_student" `
	CreatedBy       string    `json:"created_by" binding:"required"  example:"vishal"`
	DeletedBy       string    `json:"deleted_by" `
	UpdatedBy       string    `json:"updated_by" `
	CreatedAt       time.Time `json:"created_at" `
	UpdatedAt       time.Time `json:"updated_at" `
	DeletedAt       time.Time `json:"deleted_at" `
	ClassTeacherID  string    `json:"class_tracher_id"  `
	TeachersID      string    `json:"trachers_id"`
	SubjectsID      string    `json:"subjects_id"`
	FloorNumber     int       `json:"floor_number"  `
	ClassName       string    `json:"class_name"  `
}
