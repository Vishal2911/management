package store

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgress struct {
	DB *gorm.DB
}

func (store *Postgress) NewStore() error {
	dsn := "host=localhost user=vishal password=password dbname=manage port=5432 sslmode=disable"

	util.Log(model.LogLevelInfo, model.StorePackage, model.NewStore, "creating new store", nil)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, "error while creating store", err)
		return err
	} else {
		store.DB = db
	}

	err = db.AutoMigrate(
		model.User{},
	)
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, "error while running automigration", err)
		return err
	}

	fmt.Printf("db = %v\n", db)
	return nil
}

type SoteOperations interface {
	NewStore() error
	CreateUser(user *model.User) error
	GetUsers() ([]model.User, error)
	GetUser(uuid.UUID) (model.User, error)
}
