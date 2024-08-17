package store

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
	"gorm.io/gorm"
)

func (store Postgress) CreateRoom(room *model.Room) error {

	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateRoom, "creating new room", nil)
	response := store.DB.Create(room)
	if response.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, "error while creating new room", response.Error)
		return response.Error
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateRoom, "Created new room", room)
	return nil
}

func (store Postgress) GetRooms() ([]model.Room, error) {

	rooms := []model.Room{}
	util.Log(model.LogLevelInfo, model.StorePackage, model.GetRooms, "fetching records of room from db", nil)
	if err := store.DB.Find(&rooms).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetRooms, "error while fetching rooms from db", err)
		return rooms, err
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetRooms, "records of room from db", rooms)
	return rooms, nil
}

func (store Postgress) GetRoom(roomID uuid.UUID) (model.Room, error) {

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetRoom, "fetching records of room from db", nil)
	var room model.Room
	if err := store.DB.First(&room, "id = ?", roomID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			util.Log(model.LogLevelError, model.StorePackage, model.GetRoom, "room record not found", err)
		} else {
			util.Log(model.LogLevelError, model.StorePackage, model.GetRoom, "error while fetching room from db", err)
		}
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetRoom, "records of room from db", room)
	return room, nil
}

func (store Postgress) GetRoomByFilter(filter map[string]interface{}) ([]model.Room, error) {

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetRoomByFilter, "fetching records of room from db", nil)
	var rooms []model.Room
	query := store.DB

	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == model.StartDate || key == model.EndDate {
			continue
		}
		util.Log(model.LogLevelInfo, model.StorePackage, model.GetRoomByFilter,
			"filters key", key+" value = "+fmt.Sprintf("%v", value))
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	setLimitAndPage(filter, query)
	setDateRangeFilter(filter, query)

	err := query.Find(&rooms).Error
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetRoomByFilter,
			"error while reading room data", err)
		return nil, fmt.Errorf("error while fetching room records from DB, err = %v", err)
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetRoomByFilter, "records of rooms from db", rooms)
	return rooms, nil
}

func (store Postgress) UpdateRoom(room *model.Room) error {

	util.Log(model.LogLevelInfo, model.StorePackage, model.UpdateRoom, "updating room data", *room)
	resp := store.DB.Save(room)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.UpdateRoom,
			"error while updating room data", resp.Error)
		return fmt.Errorf("error while updating room record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.UpdateRoom,
		"successfully updated room", nil)
	return nil
}

// DeleteRoom is used to delete record by given roomID
func (store Postgress) DeleteRoom(roomID string) error {

	var room model.Room
	util.Log(model.LogLevelInfo, model.StorePackage, model.DeleteRoom, "deleting room data", roomID)
	if err := store.DB.First(&room, `"id" = '`+roomID+`'`).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.DeleteRoom,
			"error while deleting room data", err)
		return fmt.Errorf("room not found for given id, ID = %v", roomID)
	}
	resp := store.DB.Delete(room)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting room record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.DeleteRoom,
		"successfully deleted room", nil)
	return nil
}

