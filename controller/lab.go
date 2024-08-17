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

func (server Server) CreateLab(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.CreateLab, "creating new lab", nil)
	lab := model.Lab{}
	err := ctx.ShouldBindJSON(&lab)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateLab, "error while json binding", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	lab.ID = uuid.New()
	lab.CreatedAt = time.Now()

	err = server.PostgressDb.CreateLab(&lab)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateLab, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, lab)

}

func (server Server) GetLab(ctx *gin.Context) {

	id := ctx.Param("id")

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetLab, "fetching lab by id", map[string]interface{}{"id": id})

	uuid, err := uuid.Parse(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetLab, "invalid lab ID format", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lab ID format"})
		return
	}

	lab, err := server.PostgressDb.GetLab(uuid)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetLab, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, lab)

}

func (server Server) GetLabByFilter(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetLabByFilter, "fetching lab by filter", nil)
	queryParams := ctx.Request.URL.Query()

	filter := util.ConvertQueryParams(queryParams)

	lab, err := server.PostgressDb.GetLabByFilter(filter)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetLabByFilter, "error while fetching record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, lab)

}

func (server Server) GetLabs(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetLabs, "fetching all lab", nil)

	labs, err := server.PostgressDb.GetLabs()
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetLabs, "error while fetching labs record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, labs)

}

func (server *Server) UpdateLab(c *gin.Context) error {

	var lab model.Lab
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.UpdateLab,
		"unmarshaling lab data", nil)
	err := c.ShouldBindJSON(&lab)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.UpdateLab,
			"error while unmarshaling payload", err)
		return fmt.Errorf("")
	}
	//validation is to be done here
	//DB call
	lab.UpdatedAt = time.Now().UTC()
	err = server.PostgressDb.UpdateLab(&lab)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.UpdateLab,
			"error while updating record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetLabs,
		"successfully updated lab record and setting response", lab)
	c.JSON(http.StatusOK, lab)
	return nil

}

func (server *Server) DeleteLab(c *gin.Context) error {

	//validation is to be done here
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.DeleteLab,
		"reading lab id", nil)
	id := c.Param("id")
	if id == "" {
		util.Log(model.LogLevelError, model.ControllerPackage, model.DeleteLab,
			"missing lab id", nil)
		return fmt.Errorf("")
	}
	//DB call
	err := server.PostgressDb.DeleteLab(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.DeleteLab,
			"error while deleting lab record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.DeleteLab,
		"successfully deleted lab record ", nil)
	return nil

}
