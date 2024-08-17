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

func (server Server) CreatePublisher(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.CreatePublisher, "creating new publisher", nil)
	publisher := model.Publisher{}
	err := ctx.ShouldBindJSON(&publisher)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreatePublisher, "error while json binding", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	publisher.ID = uuid.New()
	publisher.CreatedAt = time.Now()

	err = server.PostgressDb.CreatePublisher(&publisher)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreatePublisher, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, publisher)

}

func (server Server) GetPublisher(ctx *gin.Context) {

	id := ctx.Param("id")

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetPublisher, "fetching publisher by id", map[string]interface{}{"id": id})

	uuid, err := uuid.Parse(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetPublisher, "invalid publisher ID format", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid publisher ID format"})
		return
	}

	publisher, err := server.PostgressDb.GetPublisher(uuid)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetPublisher, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, publisher)

}

func (server Server) GetPublisherByFilter(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetPublisherByFilter, "fetching publisher by filter", nil)
	queryParams := ctx.Request.URL.Query()

	filter := util.ConvertQueryParams(queryParams)

	publisher, err := server.PostgressDb.GetPublisherByFilter(filter)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetPublisherByFilter, "error while fetching record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, publisher)

}

func (server Server) GetPublishers(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetPublishers, "fetching all publisher", nil)

	publishers, err := server.PostgressDb.GetPublishers()
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetPublishers, "error while fetching publishers record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, publishers)

}

func (server *Server) UpdatePublisher(c *gin.Context) error {

	var publisher model.Publisher
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.UpdatePublisher,
		"unmarshaling publisher data", nil)
	err := c.ShouldBindJSON(&publisher)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.UpdatePublisher,
			"error while unmarshaling payload", err)
		return fmt.Errorf("")
	}
	//validation is to be done here
	//DB call
	publisher.UpdatedAt = time.Now().UTC()
	err = server.PostgressDb.UpdatePublisher(&publisher)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.UpdatePublisher,
			"error while updating record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetPublishers,
		"successfully updated publisher record and setting response", publisher)
	c.JSON(http.StatusOK, publisher)
	return nil

}

func (server *Server) DeletePublisher(c *gin.Context) error {

	//validation is to be done here
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.DeletePublisher,
		"reading publisher id", nil)
	id := c.Param("id")
	if id == "" {
		util.Log(model.LogLevelError, model.ControllerPackage, model.DeletePublisher,
			"missing publisher id", nil)
		return fmt.Errorf("")
	}
	//DB call
	err := server.PostgressDb.DeletePublisher(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.DeletePublisher,
			"error while deleting publisher record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.DeletePublisher,
		"successfully deleted publisher record ", nil)
	return nil

}
