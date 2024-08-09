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

func (server Server) CreateTeacher(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.CreateTeacher, "creating new teacher", nil)
	teacher := model.Teacher{}
	err := ctx.ShouldBindJSON(&teacher)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateTeacher, "error while json binding", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	teacher.ID = uuid.New()
	teacher.CreatedAt = time.Now()

	err = server.PostgressDb.CreateTeacher(&teacher)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateTeacher, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, teacher)

}

func (server Server) GetTeacher(ctx *gin.Context) {

	id := ctx.Param("id")

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetTeacher, "fetching teacher by id", map[string]interface{}{"id": id})

	uuid, err := uuid.Parse(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetTeacher, "invalid teacher ID format", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid teacher ID format"})
		return
	}

	teacher, err := server.PostgressDb.GetTeacher(uuid)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetTeacher, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, teacher)

}

func (server Server) GetTeacherByFilter(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetTeacherByFilter, "fetching teacher by filter", nil)
	queryParams := ctx.Request.URL.Query()

	filter := util.ConvertQueryParams(queryParams)

	teacher, err := server.PostgressDb.GetTeacherByFilter(filter)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetTeacherByFilter, "error while fetching record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, teacher)

}

func (server Server) GetTeachers(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetTeachers, "fetching all teacher", nil)

	teachers, err := server.PostgressDb.GetTeachers()
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetTeachers, "error while fetching teachers record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, teachers)

}

func (server *Server) UpdateTeacher(c *gin.Context) error {

	var teacher model.Teacher
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.UpdateTeacher,
		"unmarshaling teacher data", nil)
	err := c.ShouldBindJSON(&teacher)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.UpdateTeacher,
			"error while unmarshaling payload", err)
		return fmt.Errorf("")
	}
	//validation is to be done here
	//DB call
	teacher.UpdatedAt = time.Now().UTC()
	err = server.PostgressDb.UpdateTeacher(&teacher)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.UpdateTeacher,
			"error while updating record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetTeachers,
		"successfully updated teacher record and setting response", teacher)
	c.JSON(http.StatusOK, teacher)
	return nil

}

func (server *Server) DeleteTeacher(c *gin.Context) error {

	//validation is to be done here
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.DeleteTeacher,
		"reading teacher id", nil)
	id := c.Param("id")
	if id == "" {
		util.Log(model.LogLevelError, model.ControllerPackage, model.DeleteTeacher,
			"missing teacher id", nil)
		return fmt.Errorf("")
	}
	//DB call
	err := server.PostgressDb.DeleteTeacher(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.DeleteTeacher,
			"error while deleting teacher record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.DeleteTeacher,
		"successfully deleted teacher record ", nil)
	return nil

}
