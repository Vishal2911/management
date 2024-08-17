package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
)

func (api APIRoutes) AuthorRouts(routes *gin.Engine) {
	// Define routes
	group := routes.Group("author")
	{
		group.POST("/create", api.AuthMiddlewareComplete(), api.CreateAuthor)
		group.GET("/:id", api.AuthMiddlewareComplete(), api.GetAuthor)
		group.GET("/all", api.AuthMiddlewareComplete(), api.GetAuthors)
		group.GET("/filter", api.GetAuthorsByFilter)
		group.PUT("/update/:id", api.AuthMiddlewareComplete(), api.UpdateAuthor)
		group.DELETE("/delete/:id", api.AuthMiddlewareComplete(), api.DeleteAuthor)
	}

}

// Handler to create a author
// @router /author/create [post]
// @summary Create a author
// @tags authors
// @accept json
// @produce json
// @param author body model.Author true "Author object"
// @success 201 {object} model.Author
// @Security ApiKeyAuth
func (api APIRoutes) CreateAuthor(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.CreateAuthor, "creating new author", nil)
	api.Server.CreateAuthor(ctx)
}

// Handler to get a author by ID
// @router /author/{id} [get]
// @summary Get a author by ID
// @tags authors
// @produce json
// @param id path string true "Author ID"
// @success 200 {object} model.Author
// @Security ApiKeyAuth
func (api APIRoutes) GetAuthor(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetAuthor, "fetching  author", nil)
	api.Server.GetAuthor(ctx)
}

// Handler to get all authors
// @router /author/all [get]
// @summary Get all authors
// @tags authors
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Author
// @Security ApiKeyAuth
func (api APIRoutes) GetAuthors(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetAuthors, "fetching authors", nil)
	api.Server.GetAuthors(ctx)
}

// Handler to get all authors based on filter
// @router /author/filter [get]
// @summary Get all authors based on given filters
// @tags authors
// @produce json
// @param id query string false "id"
// @param active query bool false "active"
// @param created_by query string false "created_by"
// @param updated_by query string false "updated_by"
// @param first_name query string false "first_name"
// @param middle_name query string false "middle_name"
// @param last_name query string false "last_name"
// @param lane query string false "lane"
// @param village query string false "village"
// @param city query string false "city"
// @param district query string false "district"
// @param pincode query int false "pincode"
// @param state query string false "state"
// @param start_date query string false "start date"
// @param end_date query string false "end date"
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Author
// @Security ApiKeyAuth
func (api APIRoutes) GetAuthorsByFilter(c *gin.Context) {
	api.Server.GetAuthorByFilter(c)
}

// Handler to update a author
// @router /author/update/{id} [put]
// @summary Update a author
// @tags authors
// @accept json
// @produce json
// @param id path string true "Author ID"
// @param author body model.Author true "Author object"
// @success 200 {object} model.Author
// @Security ApiKeyAuth
func (api APIRoutes) UpdateAuthor(c *gin.Context) {
	api.Server.UpdateAuthor(c)
}

// Handler to delete a author
// @router  /author/delete/{id} [delete]
// @summary Delete a author
// @tags authors
// @param id path string true "Author ID"
// @Security ApiKeyAuth
func (api APIRoutes) DeleteAuthor(c *gin.Context) {
	api.Server.DeleteAuthor(c)
}

