package store

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
	"gorm.io/gorm"
)

func (store Postgress) CreatePublisher(publisher *model.Publisher) error {

	util.Log(model.LogLevelInfo, model.StorePackage, model.CreatePublisher, "creating new publisher", nil)
	response := store.DB.Create(publisher)
	if response.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, "error while creating new publisher", response.Error)
		return response.Error
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.CreatePublisher, "Created new publisher", publisher)
	return nil
}

func (store Postgress) GetPublishers() ([]model.Publisher, error) {

	publishers := []model.Publisher{}
	util.Log(model.LogLevelInfo, model.StorePackage, model.GetPublishers, "fetching records of publisher from db", nil)
	if err := store.DB.Find(&publishers).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetPublishers, "error while fetching publishers from db", err)
		return publishers, err
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetPublishers, "records of publisher from db", publishers)
	return publishers, nil
}

func (store Postgress) GetPublisher(publisherID uuid.UUID) (model.Publisher, error) {

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetPublisher, "fetching records of publisher from db", nil)
	var publisher model.Publisher
	if err := store.DB.First(&publisher, "id = ?", publisherID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			util.Log(model.LogLevelError, model.StorePackage, model.GetPublisher, "publisher record not found", err)
		} else {
			util.Log(model.LogLevelError, model.StorePackage, model.GetPublisher, "error while fetching publisher from db", err)
		}
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetPublisher, "records of publisher from db", publisher)
	return publisher, nil
}

func (store Postgress) GetPublisherByFilter(filter map[string]interface{}) ([]model.Publisher, error) {

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetPublisherByFilter, "fetching records of publisher from db", nil)
	var publishers []model.Publisher
	query := store.DB

	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == model.StartDate || key == model.EndDate {
			continue
		}
		util.Log(model.LogLevelInfo, model.StorePackage, model.GetPublisherByFilter,
			"filters key", key+" value = "+fmt.Sprintf("%v", value))
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	setLimitAndPage(filter, query)
	setDateRangeFilter(filter, query)

	err := query.Find(&publishers).Error
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetPublisherByFilter,
			"error while reading publisher data", err)
		return nil, fmt.Errorf("error while fetching publisher records from DB, err = %v", err)
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetPublisherByFilter, "records of publishers from db", publishers)
	return publishers, nil
}

func (store Postgress) UpdatePublisher(publisher *model.Publisher) error {

	util.Log(model.LogLevelInfo, model.StorePackage, model.UpdatePublisher, "updating publisher data", *publisher)
	resp := store.DB.Save(publisher)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.UpdatePublisher,
			"error while updating publisher data", resp.Error)
		return fmt.Errorf("error while updating publisher record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.UpdatePublisher,
		"successfully updated publisher", nil)
	return nil
}

// DeletePublisher is used to delete record by given publisherID
func (store Postgress) DeletePublisher(publisherID string) error {

	var publisher model.Publisher
	util.Log(model.LogLevelInfo, model.StorePackage, model.DeletePublisher, "deleting publisher data", publisherID)
	if err := store.DB.First(&publisher, `"id" = '`+publisherID+`'`).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.DeletePublisher,
			"error while deleting publisher data", err)
		return fmt.Errorf("publisher not found for given id, ID = %v", publisherID)
	}
	resp := store.DB.Delete(publisher)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting publisher record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.DeletePublisher,
		"successfully deleted publisher", nil)
	return nil
}

