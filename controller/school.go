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

func (server Server) CreateSchool(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.CreateSchool, "creating new school", nil)
	school := model.School{}
	err := ctx.ShouldBindJSON(&school)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateSchool, "error while json binding", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	school.ID = uuid.New()
	school.CreatedAt = time.Now()

	err = server.PostgressDb.CreateSchool(&school)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateSchool, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, school)

}

func (server Server) GetSchool(ctx *gin.Context) {

	id := ctx.Param("id")

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetSchool, "fetching school by id", map[string]interface{}{"id": id})

	uuid, err := uuid.Parse(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetSchool, "invalid school ID format", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid school ID format"})
		return
	}

	school, err := server.PostgressDb.GetSchool(uuid)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetSchool, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, school)

}

func (server Server) GetSchoolByFilter(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetSchoolByFilter, "fetching school by filter", nil)
	queryParams := ctx.Request.URL.Query()

	filter := util.ConvertQueryParams(queryParams)

	school, err := server.PostgressDb.GetSchoolByFilter(filter)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetSchoolByFilter, "error while fetching record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, school)

}

func (server Server) GetSchools(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetSchools, "fetching all school", nil)

	schools, err := server.PostgressDb.GetSchools()
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetSchools, "error while fetching schools record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, schools)

}

func (server *Server) UpdateSchool(c *gin.Context) error {

	var school model.School
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.UpdateSchool,
		"unmarshaling school data", nil)
	err := c.ShouldBindJSON(&school)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.UpdateSchool,
			"error while unmarshaling payload", err)
		return fmt.Errorf("")
	}
	//validation is to be done here
	//DB call
	school.UpdatedAt = time.Now().UTC()
	err = server.PostgressDb.UpdateSchool(&school)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.UpdateSchool,
			"error while updating record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetSchools,
		"successfully updated school record and setting response", school)
	c.JSON(http.StatusOK, school)
	return nil

}

func (server *Server) DeleteSchool(c *gin.Context) error {

	//validation is to be done here
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.DeleteSchool,
		"reading school id", nil)
	id := c.Param("id")
	if id == "" {
		util.Log(model.LogLevelError, model.ControllerPackage, model.DeleteSchool,
			"missing school id", nil)
		return fmt.Errorf("")
	}
	//DB call
	err := server.PostgressDb.DeleteSchool(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.DeleteSchool,
			"error while deleting school record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.DeleteSchool,
		"successfully deleted school record ", nil)
	return nil

}