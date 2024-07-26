package store

import (
	"github.com/google/uuid"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
	"gorm.io/gorm"
)

func (store Postgress) CreateUser(user *model.User) error {

	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateUser, "creating new user", nil)
	response := store.DB.Create(user)
	if response.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, "error while creating new user", response.Error)
		return response.Error
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateUser, "Created new user", user)
	return nil
}

func (store Postgress) GetUsers() ([]model.User, error) {

	users := []model.User{}
	util.Log(model.LogLevelInfo, model.StorePackage, model.GetUsers, "fetching records of user from db", nil)
	if err := store.DB.Find(&users).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetUsers, "error while fetching users from db", err)
		return users, err
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetUsers, "records of user from db", users)
	return users, nil
}

func (store Postgress) GetUser(userID uuid.UUID) (model.User, error) {

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetUser, "fetching records of user from db", nil)
	var user model.User
	if err := store.DB.First(&user, "id = ?", userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			util.Log(model.LogLevelError, model.StorePackage, model.GetUser, "user record not found", err)
		} else {
			util.Log(model.LogLevelError, model.StorePackage, model.GetUser, "error while fetching user from db", err)
		}
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetUser, "records of user from db", user)
	return user, nil
}
