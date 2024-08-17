package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
)

func (api APIRoutes) PublisherRouts(routes *gin.Engine) {
	// Define routes
	group := routes.Group("publisher")
	{
		group.POST("/create", api.AuthMiddlewareComplete(), api.CreatePublisher)
		group.GET("/:id", api.AuthMiddlewareComplete(), api.GetPublisher)
		group.GET("/all", api.AuthMiddlewareComplete(), api.GetPublishers)
		group.GET("/filter", api.GetPublishersByFilter)
		group.PUT("/update/:id", api.AuthMiddlewareComplete(), api.UpdatePublisher)
		group.DELETE("/delete/:id", api.AuthMiddlewareComplete(), api.DeletePublisher)
	}

}

// Handler to create a publisher
// @router /publisher/create [post]
// @summary Create a publisher
// @tags publishers
// @accept json
// @produce json
// @param publisher body model.Publisher true "Publisher object"
// @success 201 {object} model.Publisher
// @Security ApiKeyAuth
func (api APIRoutes) CreatePublisher(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.CreatePublisher, "creating new publisher", nil)
	api.Server.CreatePublisher(ctx)
}

// Handler to get a publisher by ID
// @router /publisher/{id} [get]
// @summary Get a publisher by ID
// @tags publishers
// @produce json
// @param id path string true "Publisher ID"
// @success 200 {object} model.Publisher
// @Security ApiKeyAuth
func (api APIRoutes) GetPublisher(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetPublisher, "fetching  publisher", nil)
	api.Server.GetPublisher(ctx)
}

// Handler to get all publishers
// @router /publisher/all [get]
// @summary Get all publishers
// @tags publishers
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Publisher
// @Security ApiKeyAuth
func (api APIRoutes) GetPublishers(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetPublishers, "fetching publishers", nil)
	api.Server.GetPublishers(ctx)
}

// Handler to get all publishers based on filter
// @router /publisher/filter [get]
// @summary Get all publishers based on given filters
// @tags publishers
// @produce json
// @param id query string false "id"
// @param active query bool false "active"
// @param created_by query string false "created_by"
// @param updated_by query string false "updated_by"
// @param name query string false "name"
// @param start_date query string false "start date"
// @param end_date query string false "end date"
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Publisher
// @Security ApiKeyAuth
func (api APIRoutes) GetPublishersByFilter(c *gin.Context) {
	api.Server.GetPublisherByFilter(c)
}

// Handler to update a publisher
// @router /publisher/update/{id} [put]
// @summary Update a publisher
// @tags publishers
// @accept json
// @produce json
// @param id path string true "Publisher ID"
// @param publisher body model.Publisher true "Publisher object"
// @success 200 {object} model.Publisher
// @Security ApiKeyAuth
func (api APIRoutes) UpdatePublisher(c *gin.Context) {
	api.Server.UpdatePublisher(c)
}

// Handler to delete a publisher
// @router  /publisher/delete/{id} [delete]
// @summary Delete a publisher
// @tags publishers
// @param id path string true "Publisher ID"
// @Security ApiKeyAuth
func (api APIRoutes) DeletePublisher(c *gin.Context) {
	api.Server.DeletePublisher(c)
}

