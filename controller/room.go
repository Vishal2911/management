package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
)

func (server Server) CreateRoom(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.CreateRoom, "creating new room", nil)
	room := model.Room{}
	err := ctx.ShouldBindJSON(&room)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateRoom, "error while json binding", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	room.ID = uuid.New()
	room.CreatedAt = time.Now()

	err = server.PostgressDb.CreateRoom(&room)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateRoom, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, room)

}

func (server Server) GetRoom(ctx *gin.Context) {

	id := ctx.Param("id")

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetRoom, "fetching room by id", map[string]interface{}{"id": id})

	uuid, err := uuid.Parse(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetRoom, "invalid room ID format", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID format"})
		return
	}

	room, err := server.PostgressDb.GetRoom(uuid)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetRoom, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, room)

}

func (server Server) GetRoomByFilter(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetRoomByFilter, "fetching room by filter", nil)
	queryParams := ctx.Request.URL.Query()

	filter := util.ConvertQueryParams(queryParams)

	room, err := server.PostgressDb.GetRoomByFilter(filter)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetRoomByFilter, "error while fetching record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, room)

}

func (server Server) GetRooms(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetRooms, "fetching all room", nil)

	rooms, err := server.PostgressDb.GetRooms()
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetRooms, "error while fetching rooms record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, rooms)

}

func (server *Server) UpdateRoom(c *gin.Context) error {

	var room model.Room
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.UpdateRoom,
		"unmarshaling room data", nil)
	err := c.ShouldBindJSON(&room)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.UpdateRoom,
			"error while unmarshaling payload", err)
		return fmt.Errorf("")
	}
	//validation is to be done here
	//DB call
	room.UpdatedAt = time.Now().UTC()
	err = server.PostgressDb.UpdateRoom(&room)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.UpdateRoom,
			"error while updating record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetRooms,
		"successfully updated room record and setting response", room)
	c.JSON(http.StatusOK, room)
	return nil

}

func (server *Server) DeleteRoom(c *gin.Context) error {

	//validation is to be done here
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.DeleteRoom,
		"reading room id", nil)
	id := c.Param("id")
	if id == "" {
		util.Log(model.LogLevelError, model.ControllerPackage, model.DeleteRoom,
			"missing room id", nil)
		return fmt.Errorf("")
	}
	//DB call
	err := server.PostgressDb.DeleteRoom(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.DeleteRoom,
			"error while deleting room record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.DeleteRoom,
		"successfully deleted room record ", nil)
	return nil

}
