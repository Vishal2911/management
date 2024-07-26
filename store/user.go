package store

import (
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
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
