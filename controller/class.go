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

func (server Server) CreateClass(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.CreateClass, "creating new class", nil)
	class := model.Class{}
	err := ctx.ShouldBindJSON(&class)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateClass, "error while json binding", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	class.ID = uuid.New()
	class.CreatedAt = time.Now()

	err = server.PostgressDb.CreateClass(&class)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateClass, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, class)

}

func (server Server) GetClass(ctx *gin.Context) {

	id := ctx.Param("id")

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetClass, "fetching class by id", map[string]interface{}{"id": id})

	uuid, err := uuid.Parse(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetClass, "invalid class ID format", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID format"})
		return
	}

	class, err := server.PostgressDb.GetClass(uuid)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetClass, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, class)

}

func (server Server) GetClassByFilter(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetClassByFilter, "fetching class by filter", nil)
	queryParams := ctx.Request.URL.Query()

	filter := util.ConvertQueryParams(queryParams)

	class, err := server.PostgressDb.GetClassByFilter(filter)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetClassByFilter, "error while fetching record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, class)

}

func (server Server) GetClasss(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetClasss, "fetching all class", nil)

	classs, err := server.PostgressDb.GetClasss()
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetClasss, "error while fetching classs record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, classs)

}

func (server *Server) UpdateClass(c *gin.Context) error {

	var class model.Class
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.UpdateClass,
		"unmarshaling class data", nil)
	err := c.ShouldBindJSON(&class)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.UpdateClass,
			"error while unmarshaling payload", err)
		return fmt.Errorf("")
	}
	//validation is to be done here
	//DB call
	class.UpdatedAt = time.Now().UTC()
	err = server.PostgressDb.UpdateClass(&class)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.UpdateClass,
			"error while updating record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetClasss,
		"successfully updated class record and setting response", class)
	c.JSON(http.StatusOK, class)
	return nil

}

func (server *Server) DeleteClass(c *gin.Context) error {

	//validation is to be done here
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.DeleteClass,
		"reading class id", nil)
	id := c.Param("id")
	if id == "" {
		util.Log(model.LogLevelError, model.ControllerPackage, model.DeleteClass,
			"missing class id", nil)
		return fmt.Errorf("")
	}
	//DB call
	err := server.PostgressDb.DeleteClass(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.DeleteClass,
			"error while deleting class record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.DeleteClass,
		"successfully deleted class record ", nil)
	return nil

}
