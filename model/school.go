package model

import (
	"time"

	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type School struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Name      string
	Address   Address `gorm:"embedded;embeddedPrefix:school_"`
	Director  string
	Principal string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
