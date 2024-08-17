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

func (server Server) CreateAuthor(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.CreateAuthor, "creating new author", nil)
	author := model.Author{}
	err := ctx.ShouldBindJSON(&author)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateAuthor, "error while json binding", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	author.ID = uuid.New()
	author.CreatedAt = time.Now()

	err = server.PostgressDb.CreateAuthor(&author)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateAuthor, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, author)

}

func (server Server) GetAuthor(ctx *gin.Context) {

	id := ctx.Param("id")

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetAuthor, "fetching author by id", map[string]interface{}{"id": id})

	uuid, err := uuid.Parse(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetAuthor, "invalid author ID format", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID format"})
		return
	}

	author, err := server.PostgressDb.GetAuthor(uuid)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetAuthor, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, author)

}

func (server Server) GetAuthorByFilter(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetAuthorByFilter, "fetching author by filter", nil)
	queryParams := ctx.Request.URL.Query()

	filter := util.ConvertQueryParams(queryParams)

	author, err := server.PostgressDb.GetAuthorByFilter(filter)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetAuthorByFilter, "error while fetching record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, author)

}

func (server Server) GetAuthors(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetAuthors, "fetching all author", nil)

	authors, err := server.PostgressDb.GetAuthors()
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetAuthors, "error while fetching authors record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, authors)

}

func (server *Server) UpdateAuthor(c *gin.Context) error {

	var author model.Author
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.UpdateAuthor,
		"unmarshaling author data", nil)
	err := c.ShouldBindJSON(&author)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.UpdateAuthor,
			"error while unmarshaling payload", err)
		return fmt.Errorf("")
	}
	//validation is to be done here
	//DB call
	author.UpdatedAt = time.Now().UTC()
	err = server.PostgressDb.UpdateAuthor(&author)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.UpdateAuthor,
			"error while updating record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetAuthors,
		"successfully updated author record and setting response", author)
	c.JSON(http.StatusOK, author)
	return nil

}

func (server *Server) DeleteAuthor(c *gin.Context) error {

	//validation is to be done here
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.DeleteAuthor,
		"reading author id", nil)
	id := c.Param("id")
	if id == "" {
		util.Log(model.LogLevelError, model.ControllerPackage, model.DeleteAuthor,
			"missing author id", nil)
		return fmt.Errorf("")
	}
	//DB call
	err := server.PostgressDb.DeleteAuthor(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.DeleteAuthor,
			"error while deleting author record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.DeleteAuthor,
		"successfully deleted author record ", nil)
	return nil

}
