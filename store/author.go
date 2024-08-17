package store

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
	"gorm.io/gorm"
)

func (store Postgress) CreateAuthor(teacher *model.Author) error {

	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateAuthor, "creating new teacher", nil)
	response := store.DB.Create(teacher)
	if response.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, "error while creating new teacher", response.Error)
		return response.Error
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateAuthor, "Created new teacher", teacher)
	return nil
}

func (store Postgress) GetAuthors() ([]model.Author, error) {

	teachers := []model.Author{}
	util.Log(model.LogLevelInfo, model.StorePackage, model.GetAuthors, "fetching records of teacher from db", nil)
	if err := store.DB.Find(&teachers).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetAuthors, "error while fetching teachers from db", err)
		return teachers, err
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetAuthors, "records of teacher from db", teachers)
	return teachers, nil
}

func (store Postgress) GetAuthor(teacherID uuid.UUID) (model.Author, error) {

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetAuthor, "fetching records of teacher from db", nil)
	var teacher model.Author
	if err := store.DB.First(&teacher, "id = ?", teacherID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			util.Log(model.LogLevelError, model.StorePackage, model.GetAuthor, "teacher record not found", err)
		} else {
			util.Log(model.LogLevelError, model.StorePackage, model.GetAuthor, "error while fetching teacher from db", err)
		}
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetAuthor, "records of teacher from db", teacher)
	return teacher, nil
}

func (store Postgress) GetAuthorByFilter(filter map[string]interface{}) ([]model.Author, error) {

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetAuthorByFilter, "fetching records of teacher from db", nil)
	var teachers []model.Author
	query := store.DB

	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == model.StartDate || key == model.EndDate {
			continue
		}
		util.Log(model.LogLevelInfo, model.StorePackage, model.GetAuthorByFilter,
			"filters key", key+" value = "+fmt.Sprintf("%v", value))
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	setLimitAndPage(filter, query)
	setDateRangeFilter(filter, query)

	err := query.Find(&teachers).Error
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetAuthorByFilter,
			"error while reading teacher data", err)
		return nil, fmt.Errorf("error while fetching teacher records from DB, err = %v", err)
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetAuthorByFilter, "records of teachers from db", teachers)
	return teachers, nil
}

func (store Postgress) UpdateAuthor(teacher *model.Author) error {

	util.Log(model.LogLevelInfo, model.StorePackage, model.UpdateAuthor, "updating teacher data", *teacher)
	resp := store.DB.Save(teacher)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.UpdateAuthor,
			"error while updating teacher data", resp.Error)
		return fmt.Errorf("error while updating teacher record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.UpdateAuthor,
		"successfully updated teacher", nil)
	return nil
}

// DeleteAuthor is used to delete record by given teacherID
func (store Postgress) DeleteAuthor(teacherID string) error {

	var teacher model.Author
	util.Log(model.LogLevelInfo, model.StorePackage, model.DeleteAuthor, "deleting teacher data", teacherID)
	if err := store.DB.First(&teacher, `"id" = '`+teacherID+`'`).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.DeleteAuthor,
			"error while deleting teacher data", err)
		return fmt.Errorf("teacher not found for given id, ID = %v", teacherID)
	}
	resp := store.DB.Delete(teacher)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting teacher record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.DeleteAuthor,
		"successfully deleted teacher", nil)
	return nil
}

