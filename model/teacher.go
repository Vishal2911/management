package model

import (
	"time"

	"github.com/google/uuid"
)

// Teacher struct
type Teacher struct {
	ID          uuid.UUID `gorm:"primarykey" json:"id"`
	FirstName   string    `json:"first_name" binding:"required" example:"John"`
	MiddleName  string    `json:"middle_name" binding:"required"  example:"Doe"`
	LastName    string    `json:"last_name" binding:"required" example:"Smith"`
	Email       string    `json:"email" binding:"required" gorm:"unique;not null" example:"vishal"`
	Password    string    `json:"password" binding:"required" gorm:"not null"  example:"password"`
	Lane        string    `json:"lane" example:"1234 Elm St"`
	Village     string    `json:"village" example:"Springfield"`
	City        string    `json:"city" binding:"required" example:"Metropolis"`
	District    string    `json:"district" binding:"required" example:"Central"`
	Pincode     int       `json:"pincode" binding:"required" example:"123456"`
	State       string    `json:"state" binding:"required" example:"NY"`
	SchoolID    uuid.UUID `binding:"required" json:"school_id"`
	Active      bool      `json:"active"  binding:"required" example:"true"`
	Salary      int       `json:"salary" `
	CreatedBy   string    `json:"created_by" binding:"required"  example:"vishal"`
	DeletedBy   string    `json:"deleted_by" `
	UpdatedBy   string    `json:"updated_by" `
	CreatedAt   time.Time `json:"created_at" `
	UpdatedAt   time.Time `json:"updated_at" `
	DeletedAt   time.Time `json:"deleted_at" `
	ClassesID   string    `json:"classes_id" `
	JoiningDate time.Time `json:"joining_date" `
	SubjectsID  string    `json:"subjects_id"`
	Skills      string    `json:"skills"`
}
