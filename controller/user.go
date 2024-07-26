package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
)

func (server Server) CreateUser(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.CreateUser, "creating new user", nil)
	user := model.User{}
	err := ctx.Bind(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateUser, "error while json binding", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	user.ID = uuid.New()
	user.CreatedAt = time.Now()

	err = server.PostgressDb.CreateUser(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateUser, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, user)

}

func (server Server) GetUser(ctx *gin.Context) {

	id := ctx.Param("id")

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetUser, "fetching user by id", map[string]interface{}{"id": id})

	uuid, err := uuid.Parse(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetUser, "invalid user ID format", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	user, err := server.PostgressDb.GetUser(uuid)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetUser, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, user)

}

func (server Server) GetUsers(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetUsers, "fetching all user", nil)

	users, err := server.PostgressDb.GetUsers()
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetUsers, "error while fetching users record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, users)

}
