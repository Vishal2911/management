package store

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
	"gorm.io/gorm"
)

func (store Postgress) CreateSubject(subject *model.Subject) error {

	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateSubject, "creating new subject", nil)
	response := store.DB.Create(subject)
	if response.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, "error while creating new subject", response.Error)
		return response.Error
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateSubject, "Created new subject", subject)
	return nil
}

func (store Postgress) GetSubjects() ([]model.Subject, error) {

	subjects := []model.Subject{}
	util.Log(model.LogLevelInfo, model.StorePackage, model.GetSubjects, "fetching records of subject from db", nil)
	if err := store.DB.Find(&subjects).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetSubjects, "error while fetching subjects from db", err)
		return subjects, err
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetSubjects, "records of subject from db", subjects)
	return subjects, nil
}

func (store Postgress) GetSubject(subjectID uuid.UUID) (model.Subject, error) {

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetSubject, "fetching records of subject from db", nil)
	var subject model.Subject
	if err := store.DB.First(&subject, "id = ?", subjectID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			util.Log(model.LogLevelError, model.StorePackage, model.GetSubject, "subject record not found", err)
		} else {
			util.Log(model.LogLevelError, model.StorePackage, model.GetSubject, "error while fetching subject from db", err)
		}
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetSubject, "records of subject from db", subject)
	return subject, nil
}

func (store Postgress) GetSubjectByFilter(filter map[string]interface{}) ([]model.Subject, error) {

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetSubjectByFilter, "fetching records of subject from db", nil)
	var subjects []model.Subject
	query := store.DB

	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == model.StartDate || key == model.EndDate {
			continue
		}
		util.Log(model.LogLevelInfo, model.StorePackage, model.GetSubjectByFilter,
			"filters key", key+" value = "+fmt.Sprintf("%v", value))
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	setLimitAndPage(filter, query)
	setDateRangeFilter(filter, query)

	err := query.Find(&subjects).Error
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetSubjectByFilter,
			"error while reading subject data", err)
		return nil, fmt.Errorf("error while fetching subject records from DB, err = %v", err)
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetSubjectByFilter, "records of subjects from db", subjects)
	return subjects, nil
}

func (store Postgress) UpdateSubject(subject *model.Subject) error {

	util.Log(model.LogLevelInfo, model.StorePackage, model.UpdateSubject, "updating subject data", *subject)
	resp := store.DB.Save(subject)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.UpdateSubject,
			"error while updating subject data", resp.Error)
		return fmt.Errorf("error while updating subject record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.UpdateSubject,
		"successfully updated subject", nil)
	return nil
}

// DeleteSubject is used to delete record by given subjectID
func (store Postgress) DeleteSubject(subjectID string) error {

	var subject model.Subject
	util.Log(model.LogLevelInfo, model.StorePackage, model.DeleteSubject, "deleting subject data", subjectID)
	if err := store.DB.First(&subject, `"id" = '`+subjectID+`'`).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.DeleteSubject,
			"error while deleting subject data", err)
		return fmt.Errorf("subject not found for given id, ID = %v", subjectID)
	}
	resp := store.DB.Delete(subject)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting subject record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.DeleteSubject,
		"successfully deleted subject", nil)
	return nil
}

