package store

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
	"gorm.io/gorm"
)

func (store Postgress) CreateTeacher(teacher *model.Teacher) error {

	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateTeacher, "creating new teacher", nil)
	response := store.DB.Create(teacher)
	if response.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, "error while creating new teacher", response.Error)
		return response.Error
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateTeacher, "Created new teacher", teacher)
	return nil
}

func (store Postgress) GetTeachers() ([]model.Teacher, error) {

	teachers := []model.Teacher{}
	util.Log(model.LogLevelInfo, model.StorePackage, model.GetTeachers, "fetching records of teacher from db", nil)
	if err := store.DB.Find(&teachers).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetTeachers, "error while fetching teachers from db", err)
		return teachers, err
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetTeachers, "records of teacher from db", teachers)
	return teachers, nil
}

func (store Postgress) GetTeacher(teacherID uuid.UUID) (model.Teacher, error) {

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetTeacher, "fetching records of teacher from db", nil)
	var teacher model.Teacher
	if err := store.DB.First(&teacher, "id = ?", teacherID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			util.Log(model.LogLevelError, model.StorePackage, model.GetTeacher, "teacher record not found", err)
		} else {
			util.Log(model.LogLevelError, model.StorePackage, model.GetTeacher, "error while fetching teacher from db", err)
		}
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetTeacher, "records of teacher from db", teacher)
	return teacher, nil
}

func (store Postgress) GetTeacherByFilter(filter map[string]interface{}) ([]model.Teacher, error) {

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetTeacherByFilter, "fetching records of teacher from db", nil)
	var teachers []model.Teacher
	query := store.DB

	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == model.StartDate || key == model.EndDate {
			continue
		}
		util.Log(model.LogLevelInfo, model.StorePackage, model.GetTeacherByFilter,
			"filters key", key+" value = "+fmt.Sprintf("%v", value))
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	setLimitAndPage(filter, query)
	setDateRangeFilter(filter, query)

	err := query.Find(&teachers).Error
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetTeacherByFilter,
			"error while reading teacher data", err)
		return nil, fmt.Errorf("error while fetching teacher records from DB, err = %v", err)
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetTeacherByFilter, "records of teachers from db", teachers)
	return teachers, nil
}

func (store Postgress) UpdateTeacher(teacher *model.Teacher) error {

	util.Log(model.LogLevelInfo, model.StorePackage, model.UpdateTeacher, "updating teacher data", *teacher)
	resp := store.DB.Save(teacher)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.UpdateTeacher,
			"error while updating teacher data", resp.Error)
		return fmt.Errorf("error while updating teacher record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.UpdateTeacher,
		"successfully updated teacher", nil)
	return nil
}

// DeleteTeacher is used to delete record by given teacherID
func (store Postgress) DeleteTeacher(teacherID string) error {

	var teacher model.Teacher
	util.Log(model.LogLevelInfo, model.StorePackage, model.DeleteTeacher, "deleting teacher data", teacherID)
	if err := store.DB.First(&teacher, `"id" = '`+teacherID+`'`).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.DeleteTeacher,
			"error while deleting teacher data", err)
		return fmt.Errorf("teacher not found for given id, ID = %v", teacherID)
	}
	resp := store.DB.Delete(teacher)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting teacher record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.DeleteTeacher,
		"successfully deleted teacher", nil)
	return nil
}

