package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
)

func (api APIRoutes) RoomRouts(routes *gin.Engine) {
	// Define routes
	group := routes.Group("room")
	{
		group.POST("/create", api.AuthMiddlewareComplete(), api.CreateRoom)
		group.GET("/:id", api.AuthMiddlewareComplete(), api.GetRoom)
		group.GET("/all", api.AuthMiddlewareComplete(), api.GetRooms)
		group.GET("/filter", api.GetRoomsByFilter)
		group.PUT("/update/:id", api.AuthMiddlewareComplete(), api.UpdateRoom)
		group.DELETE("/delete/:id", api.AuthMiddlewareComplete(), api.DeleteRoom)
	}

}

// Handler to create a room
// @router /room/create [post]
// @summary Create a room
// @tags rooms
// @accept json
// @produce json
// @param room body model.Room true "Room object"
// @success 201 {object} model.Room
// @Security ApiKeyAuth
func (api APIRoutes) CreateRoom(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.CreateRoom, "creating new room", nil)
	api.Server.CreateRoom(ctx)
}

// Handler to get a room by ID
// @router /room/{id} [get]
// @summary Get a room by ID
// @tags rooms
// @produce json
// @param id path string true "Room ID"
// @success 200 {object} model.Room
// @Security ApiKeyAuth
func (api APIRoutes) GetRoom(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetRoom, "fetching  room", nil)
	api.Server.GetRoom(ctx)
}

// Handler to get all rooms
// @router /room/all [get]
// @summary Get all rooms
// @tags rooms
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Room
// @Security ApiKeyAuth
func (api APIRoutes) GetRooms(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetRooms, "fetching rooms", nil)
	api.Server.GetRooms(ctx)
}

// Handler to get all rooms based on filter
// @router /room/filter [get]
// @summary Get all rooms based on given filters
// @tags rooms
// @produce json
// @param id query string false "id"
// @param school_id query string false "school_id"
// @param width query int false "width"

// @param hight query int false "hight"
// @param length query int false "length"
// @param room_number query int false "room_number"
// @param floor_number query int false "floor_number"
// @param room_type query string false "room_type"
// @param room_name query string false "room_name"
// @param ac query int bool "ac"
// @param stared query boolean false "stared"
// @param active query bool false "active"
// @param created_by query string false "created_by"
// @param updated_by query string false "updated_by"
// @param name query string false "name"
// @param start_date query string false "start date"
// @param end_date query string false "end date"
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Room
// @Security ApiKeyAuth
func (api APIRoutes) GetRoomsByFilter(c *gin.Context) {
	api.Server.GetRoomByFilter(c)
}

// Handler to update a room
// @router /room/update/{id} [put]
// @summary Update a room
// @tags rooms
// @accept json
// @produce json
// @param id path string true "Room ID"
// @param room body model.Room true "Room object"
// @success 200 {object} model.Room
// @Security ApiKeyAuth
func (api APIRoutes) UpdateRoom(c *gin.Context) {
	api.Server.UpdateRoom(c)
}

// Handler to delete a room
// @router  /room/delete/{id} [delete]
// @summary Delete a room
// @tags rooms
// @param id path string true "Room ID"
// @Security ApiKeyAuth
func (api APIRoutes) DeleteRoom(c *gin.Context) {
	api.Server.DeleteRoom(c)
}

