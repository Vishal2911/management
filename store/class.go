package store

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
	"gorm.io/gorm"
)

func (store Postgress) CreateClass(class *model.Class) error {

	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateClass, "creating new class", nil)
	response := store.DB.Create(class)
	if response.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, "error while creating new class", response.Error)
		return response.Error
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateClass, "Created new class", class)
	return nil
}

func (store Postgress) GetClasss() ([]model.Class, error) {

	classs := []model.Class{}
	util.Log(model.LogLevelInfo, model.StorePackage, model.GetClasss, "fetching records of class from db", nil)
	if err := store.DB.Find(&classs).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetClasss, "error while fetching classs from db", err)
		return classs, err
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetClasss, "records of class from db", classs)
	return classs, nil
}

func (store Postgress) GetClass(classID uuid.UUID) (model.Class, error) {

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetClass, "fetching records of class from db", nil)
	var class model.Class
	if err := store.DB.First(&class, "id = ?", classID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			util.Log(model.LogLevelError, model.StorePackage, model.GetClass, "class record not found", err)
		} else {
			util.Log(model.LogLevelError, model.StorePackage, model.GetClass, "error while fetching class from db", err)
		}
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetClass, "records of class from db", class)
	return class, nil
}

func (store Postgress) GetClassByFilter(filter map[string]interface{}) ([]model.Class, error) {

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetClassByFilter, "fetching records of class from db", nil)
	var classs []model.Class
	query := store.DB

	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == model.StartDate || key == model.EndDate {
			continue
		}
		util.Log(model.LogLevelInfo, model.StorePackage, model.GetClassByFilter,
			"filters key", key+" value = "+fmt.Sprintf("%v", value))
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	setLimitAndPage(filter, query)
	setDateRangeFilter(filter, query)

	err := query.Find(&classs).Error
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetClassByFilter,
			"error while reading class data", err)
		return nil, fmt.Errorf("error while fetching class records from DB, err = %v", err)
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetClassByFilter, "records of classs from db", classs)
	return classs, nil
}

func (store Postgress) UpdateClass(class *model.Class) error {

	util.Log(model.LogLevelInfo, model.StorePackage, model.UpdateClass, "updating class data", *class)
	resp := store.DB.Save(class)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.UpdateClass,
			"error while updating class data", resp.Error)
		return fmt.Errorf("error while updating class record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.UpdateClass,
		"successfully updated class", nil)
	return nil
}

// DeleteClass is used to delete record by given classID
func (store Postgress) DeleteClass(classID string) error {

	var class model.Class
	util.Log(model.LogLevelInfo, model.StorePackage, model.DeleteClass, "deleting class data", classID)
	if err := store.DB.First(&class, `"id" = '`+classID+`'`).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.DeleteClass,
			"error while deleting class data", err)
		return fmt.Errorf("class not found for given id, ID = %v", classID)
	}
	resp := store.DB.Delete(class)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting class record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.DeleteClass,
		"successfully deleted class", nil)
	return nil
}

