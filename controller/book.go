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

func (server Server) CreateBook(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.CreateBook, "creating new book", nil)
	book := model.Book{}
	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateBook, "error while json binding", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	book.ID = uuid.New()
	book.CreatedAt = time.Now()

	err = server.PostgressDb.CreateBook(&book)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateBook, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, book)

}

func (server Server) GetBook(ctx *gin.Context) {

	id := ctx.Param("id")

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetBook, "fetching book by id", map[string]interface{}{"id": id})

	uuid, err := uuid.Parse(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetBook, "invalid book ID format", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID format"})
		return
	}

	book, err := server.PostgressDb.GetBook(uuid)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetBook, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, book)

}

func (server Server) GetBookByFilter(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetBookByFilter, "fetching book by filter", nil)
	queryParams := ctx.Request.URL.Query()

	filter := util.ConvertQueryParams(queryParams)

	book, err := server.PostgressDb.GetBookByFilter(filter)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetBookByFilter, "error while fetching record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, book)

}

func (server Server) GetBooks(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetBooks, "fetching all book", nil)

	books, err := server.PostgressDb.GetBooks()
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetBooks, "error while fetching books record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, books)

}

func (server *Server) UpdateBook(c *gin.Context) error {

	var book model.Book
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.UpdateBook,
		"unmarshaling book data", nil)
	err := c.ShouldBindJSON(&book)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.UpdateBook,
			"error while unmarshaling payload", err)
		return fmt.Errorf("")
	}
	//validation is to be done here
	//DB call
	book.UpdatedAt = time.Now().UTC()
	err = server.PostgressDb.UpdateBook(&book)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.UpdateBook,
			"error while updating record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetBooks,
		"successfully updated book record and setting response", book)
	c.JSON(http.StatusOK, book)
	return nil

}

func (server *Server) DeleteBook(c *gin.Context) error {

	//validation is to be done here
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.DeleteBook,
		"reading book id", nil)
	id := c.Param("id")
	if id == "" {
		util.Log(model.LogLevelError, model.ControllerPackage, model.DeleteBook,
			"missing book id", nil)
		return fmt.Errorf("")
	}
	//DB call
	err := server.PostgressDb.DeleteBook(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.DeleteBook,
			"error while deleting book record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.DeleteBook,
		"successfully deleted book record ", nil)
	return nil

}
