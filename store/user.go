package store

import (
	"fmt"

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

func (store Postgress) GetUserByFilter(filter map[string]interface{}) ([]model.User, error) {

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetUserByFilter, "fetching records of user from db", nil)
	var users []model.User
	query := store.DB

	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == model.StartDate || key == model.EndDate {
			continue
		}
		util.Log(model.LogLevelInfo, model.StorePackage, model.GetUserByFilter,
			"filters key", key+" value = "+fmt.Sprintf("%v", value))
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	err := query.Find(&users).Error
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetUserByFilter,
			"error while reading user data", err)
		return nil, fmt.Errorf("error while fetching user records from DB, err = %v", err)
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetUserByFilter, "records of users from db", users)
	return users, nil
}


func (store Postgress) UpdateUser(user *model.User) error {

	util.Log(model.LogLevelInfo, model.StorePackage, model.UpdateUser, "updating user data", *user)
	resp := store.DB.Save(user)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.UpdateUser,
			"error while updating user data", resp.Error)
		return fmt.Errorf("error while updating user record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.UpdateUser,
		"successfully updated user", nil)
	return nil
}

// DeleteUser is used to delete record by given userID
func (store Postgress) DeleteUser(userID string) error {

	var user model.User
	util.Log(model.LogLevelInfo, model.StorePackage, model.DeleteUser, "deleting user data", userID)
	if err := store.DB.First(&user, `"id" = '`+userID+`'`).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.DeleteUser,
			"error while deleting user data", err)
		return fmt.Errorf("user not found for given id, ID = %v", userID)
	}
	resp := store.DB.Delete(user)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting user record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.DeleteUser,
		"successfully deleted user", nil)
	return nil
}