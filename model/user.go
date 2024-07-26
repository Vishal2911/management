package model

import (
	"time"

	"github.com/google/uuid"
)

// User struct
type User struct {
	ID        uuid.UUID `gorm:"primarykey" json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Name      Name      `gorm:"embedded" binding:"required" json:"name"`
	Address   Address   `gorm:"embedded" binding:"required" json:"address"`
	Active    bool      `json:"active" example:"true"`
	CreatedBy string    `json:"created_by" binding:"required"  example:"admin"`
	Email     string    `json:"email" binding:"required" gorm:"unique;not null" example:"vishal"`
	Password  string    `json:"password" binding:"required" gorm:"not null"  example:"password"`
	CreatedAt time.Time `json:"created_at" example:"2024-07-27T00:00:00Z"`
	UpdatedBy string    `json:"updated_by"  example:"admin"`
	UpdatedAt time.Time `json:"updated_at" example:"2024-07-27T00:00:00Z"`
	DeletedBy string    `json:"deleted_by" example:"admin"`
	DeletedAt time.Time `json:"deleted_at" example:"2024-07-27T00:00:00Z"`
}

type Name struct {
	FirstName  string `json:"first_name" binding:"required" example:"John"`
	MiddleName string `json:"middle_name" binding:"required"  example:"Doe"`
	LastName   string `json:"last_name" binding:"required" example:"Smith"`
}

type Address struct {
	Lane     string `json:"lane" example:"1234 Elm St"`
	Village  string `json:"village" example:"Springfield"`
	City     string `json:"city" binding:"required" example:"Metropolis"`
	District string `json:"district" binding:"required" example:"Central"`
	Pincode  int    `json:"pincode" binding:"required" example:"123456"`
	State    string `json:"state" binding:"required" example:"NY"`
}
