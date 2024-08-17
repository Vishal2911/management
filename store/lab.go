package store

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
	"gorm.io/gorm"
)

func (store Postgress) CreateLab(lab *model.Lab) error {

	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateLab, "creating new lab", nil)
	response := store.DB.Create(lab)
	if response.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, "error while creating new lab", response.Error)
		return response.Error
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateLab, "Created new lab", lab)
	return nil
}

func (store Postgress) GetLabs() ([]model.Lab, error) {

	labs := []model.Lab{}
	util.Log(model.LogLevelInfo, model.StorePackage, model.GetLabs, "fetching records of lab from db", nil)
	if err := store.DB.Find(&labs).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetLabs, "error while fetching labs from db", err)
		return labs, err
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetLabs, "records of lab from db", labs)
	return labs, nil
}

func (store Postgress) GetLab(labID uuid.UUID) (model.Lab, error) {

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetLab, "fetching records of lab from db", nil)
	var lab model.Lab
	if err := store.DB.First(&lab, "id = ?", labID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			util.Log(model.LogLevelError, model.StorePackage, model.GetLab, "lab record not found", err)
		} else {
			util.Log(model.LogLevelError, model.StorePackage, model.GetLab, "error while fetching lab from db", err)
		}
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetLab, "records of lab from db", lab)
	return lab, nil
}

func (store Postgress) GetLabByFilter(filter map[string]interface{}) ([]model.Lab, error) {

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetLabByFilter, "fetching records of lab from db", nil)
	var labs []model.Lab
	query := store.DB

	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == model.StartDate || key == model.EndDate {
			continue
		}
		util.Log(model.LogLevelInfo, model.StorePackage, model.GetLabByFilter,
			"filters key", key+" value = "+fmt.Sprintf("%v", value))
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	setLimitAndPage(filter, query)
	setDateRangeFilter(filter, query)

	err := query.Find(&labs).Error
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetLabByFilter,
			"error while reading lab data", err)
		return nil, fmt.Errorf("error while fetching lab records from DB, err = %v", err)
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetLabByFilter, "records of labs from db", labs)
	return labs, nil
}

func (store Postgress) UpdateLab(lab *model.Lab) error {

	util.Log(model.LogLevelInfo, model.StorePackage, model.UpdateLab, "updating lab data", *lab)
	resp := store.DB.Save(lab)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.UpdateLab,
			"error while updating lab data", resp.Error)
		return fmt.Errorf("error while updating lab record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.UpdateLab,
		"successfully updated lab", nil)
	return nil
}

// DeleteLab is used to delete record by given labID
func (store Postgress) DeleteLab(labID string) error {

	var lab model.Lab
	util.Log(model.LogLevelInfo, model.StorePackage, model.DeleteLab, "deleting lab data", labID)
	if err := store.DB.First(&lab, `"id" = '`+labID+`'`).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.DeleteLab,
			"error while deleting lab data", err)
		return fmt.Errorf("lab not found for given id, ID = %v", labID)
	}
	resp := store.DB.Delete(lab)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting lab record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.DeleteLab,
		"successfully deleted lab", nil)
	return nil
}

