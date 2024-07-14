package store

import (
	"fmt"

	"github.com/vishal2911/management/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgress struct {
	DB *gorm.DB
}

func (store *Postgress) NewStore() error {
	dsn := "host=localhost user=vishal password=password dbname=manage port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	} else {
		store.DB = db
	}

	db.AutoMigrate(&model.School{})
	fmt.Printf("db = %v\n", db)
	return nil
}

type SoteOperations interface {
	NewStore() error
	CreateSchool(school model.School) *gorm.DB
}
