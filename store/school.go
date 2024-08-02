package store

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
	"gorm.io/gorm"
)

func (store Postgress) CreateSchool(school *model.School) error {

	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateSchool, "creating new school", nil)
	response := store.DB.Create(school)
	if response.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, "error while creating new school", response.Error)
		return response.Error
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateSchool, "Created new school", school)
	return nil
}

func (store Postgress) GetSchools() ([]model.School, error) {

	schools := []model.School{}
	util.Log(model.LogLevelInfo, model.StorePackage, model.GetSchools, "fetching records of school from db", nil)
	if err := store.DB.Find(&schools).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetSchools, "error while fetching schools from db", err)
		return schools, err
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetSchools, "records of school from db", schools)
	return schools, nil
}

func (store Postgress) GetSchool(schoolID uuid.UUID) (model.School, error) {

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetSchool, "fetching records of school from db", nil)
	var school model.School
	if err := store.DB.First(&school, "id = ?", schoolID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			util.Log(model.LogLevelError, model.StorePackage, model.GetSchool, "school record not found", err)
		} else {
			util.Log(model.LogLevelError, model.StorePackage, model.GetSchool, "error while fetching school from db", err)
		}
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetSchool, "records of school from db", school)
	return school, nil
}

func (store Postgress) GetSchoolByFilter(filter map[string]interface{}) ([]model.School, error) {

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetSchoolByFilter, "fetching records of school from db", nil)
	var schools []model.School
	query := store.DB

	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == model.StartDate || key == model.EndDate {
			continue
		}
		util.Log(model.LogLevelInfo, model.StorePackage, model.GetSchoolByFilter,
			"filters key", key+" value = "+fmt.Sprintf("%v", value))
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	err := query.Find(&schools).Error
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetSchoolByFilter,
			"error while reading school data", err)
		return nil, fmt.Errorf("error while fetching school records from DB, err = %v", err)
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetSchoolByFilter, "records of schools from db", schools)
	return schools, nil
}


func (store Postgress) UpdateSchool(school *model.School) error {

	util.Log(model.LogLevelInfo, model.StorePackage, model.UpdateSchool, "updating school data", *school)
	resp := store.DB.Save(school)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.UpdateSchool,
			"error while updating school data", resp.Error)
		return fmt.Errorf("error while updating school record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.UpdateSchool,
		"successfully updated school", nil)
	return nil
}

// DeleteSchool is used to delete record by given schoolID
func (store Postgress) DeleteSchool(schoolID string) error {

	var school model.School
	util.Log(model.LogLevelInfo, model.StorePackage, model.DeleteSchool, "deleting school data", schoolID)
	if err := store.DB.First(&school, `"id" = '`+schoolID+`'`).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.DeleteSchool,
			"error while deleting school data", err)
		return fmt.Errorf("school not found for given id, ID = %v", schoolID)
	}
	resp := store.DB.Delete(school)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting school record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.DeleteSchool,
		"successfully deleted school", nil)
	return nil
}
