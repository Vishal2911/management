package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
)

func (api APIRoutes) BookRouts(routes *gin.Engine) {
	// Define routes
	group := routes.Group("book")
	{
		group.POST("/create", api.AuthMiddlewareComplete(), api.CreateBook)
		group.GET("/:id", api.AuthMiddlewareComplete(), api.GetBook)
		group.GET("/all", api.AuthMiddlewareComplete(), api.GetBooks)
		group.GET("/filter", api.GetBooksByFilter)
		group.PUT("/update/:id", api.AuthMiddlewareComplete(), api.UpdateBook)
		group.DELETE("/delete/:id", api.AuthMiddlewareComplete(), api.DeleteBook)
	}

}

// Handler to create a book
// @router /book/create [post]
// @summary Create a book
// @tags books
// @accept json
// @produce json
// @param book body model.Book true "Book object"
// @success 201 {object} model.Book
// @Security ApiKeyAuth
func (api APIRoutes) CreateBook(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.CreateBook, "creating new book", nil)
	api.Server.CreateBook(ctx)
}

// Handler to get a book by ID
// @router /book/{id} [get]
// @summary Get a book by ID
// @tags books
// @produce json
// @param id path string true "Book ID"
// @success 200 {object} model.Book
// @Security ApiKeyAuth
func (api APIRoutes) GetBook(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetBook, "fetching  book", nil)
	api.Server.GetBook(ctx)
}

// Handler to get all books
// @router /book/all [get]
// @summary Get all books
// @tags books
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Book
// @Security ApiKeyAuth
func (api APIRoutes) GetBooks(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetBooks, "fetching books", nil)
	api.Server.GetBooks(ctx)
}

// Handler to get all books based on filter
// @router /book/filter [get]
// @summary Get all books based on given filters
// @tags books
// @produce json
// @param id query string false "id"
// @param school_id query string false "school_id"
// @param edition_number query string false "edition_number"
// @param active query bool false "active"
// @param created_by query string false "created_by"
// @param updated_by query string false "updated_by"
// @param book_name query string false "book_name"
// @param start_date query string false "start date"
// @param end_date query string false "end date"
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Book
// @Security ApiKeyAuth
func (api APIRoutes) GetBooksByFilter(c *gin.Context) {
	api.Server.GetBookByFilter(c)
}

// Handler to update a book
// @router /book/update/{id} [put]
// @summary Update a book
// @tags books
// @accept json
// @produce json
// @param id path string true "Book ID"
// @param book body model.Book true "Book object"
// @success 200 {object} model.Book
// @Security ApiKeyAuth
func (api APIRoutes) UpdateBook(c *gin.Context) {
	api.Server.UpdateBook(c)
}

// Handler to delete a book
// @router  /book/delete/{id} [delete]
// @summary Delete a book
// @tags books
// @param id path string true "Book ID"
// @Security ApiKeyAuth
func (api APIRoutes) DeleteBook(c *gin.Context) {
	api.Server.DeleteBook(c)
}

