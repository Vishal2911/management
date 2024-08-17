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

func (server Server) CreateSubject(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.CreateSubject, "creating new subject", nil)
	subject := model.Subject{}
	err := ctx.ShouldBindJSON(&subject)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateSubject, "error while json binding", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	subject.ID = uuid.New()
	subject.CreatedAt = time.Now()

	err = server.PostgressDb.CreateSubject(&subject)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateSubject, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, subject)

}

func (server Server) GetSubject(ctx *gin.Context) {

	id := ctx.Param("id")

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetSubject, "fetching subject by id", map[string]interface{}{"id": id})

	uuid, err := uuid.Parse(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetSubject, "invalid subject ID format", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subject ID format"})
		return
	}

	subject, err := server.PostgressDb.GetSubject(uuid)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetSubject, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, subject)

}

func (server Server) GetSubjectByFilter(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetSubjectByFilter, "fetching subject by filter", nil)
	queryParams := ctx.Request.URL.Query()

	filter := util.ConvertQueryParams(queryParams)

	subject, err := server.PostgressDb.GetSubjectByFilter(filter)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetSubjectByFilter, "error while fetching record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, subject)

}

func (server Server) GetSubjects(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetSubjects, "fetching all subject", nil)

	subjects, err := server.PostgressDb.GetSubjects()
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetSubjects, "error while fetching subjects record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, subjects)

}

func (server *Server) UpdateSubject(c *gin.Context) error {

	var subject model.Subject
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.UpdateSubject,
		"unmarshaling subject data", nil)
	err := c.ShouldBindJSON(&subject)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.UpdateSubject,
			"error while unmarshaling payload", err)
		return fmt.Errorf("")
	}
	//validation is to be done here
	//DB call
	subject.UpdatedAt = time.Now().UTC()
	err = server.PostgressDb.UpdateSubject(&subject)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.UpdateSubject,
			"error while updating record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetSubjects,
		"successfully updated subject record and setting response", subject)
	c.JSON(http.StatusOK, subject)
	return nil

}

func (server *Server) DeleteSubject(c *gin.Context) error {

	//validation is to be done here
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.DeleteSubject,
		"reading subject id", nil)
	id := c.Param("id")
	if id == "" {
		util.Log(model.LogLevelError, model.ControllerPackage, model.DeleteSubject,
			"missing subject id", nil)
		return fmt.Errorf("")
	}
	//DB call
	err := server.PostgressDb.DeleteSubject(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.DeleteSubject,
			"error while deleting subject record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.DeleteSubject,
		"successfully deleted subject record ", nil)
	return nil

}
