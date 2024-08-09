package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
)

func (api APIRoutes) ClassRouts(routes *gin.Engine) {
	// Define routes
	group := routes.Group("class")
	{
		group.POST("/create", api.AuthMiddlewareComplete(), api.CreateClass)
		group.GET("/:id", api.AuthMiddlewareComplete(), api.GetClass)
		group.GET("/all", api.AuthMiddlewareComplete(), api.GetClasss)
		group.GET("/filter", api.GetClasssByFilter)
		group.PUT("/update/:id", api.AuthMiddlewareComplete(), api.UpdateClass)
		group.DELETE("/delete/:id", api.AuthMiddlewareComplete(), api.DeleteClass)
	}

}

// Handler to create a class
// @router /class/create [post]
// @summary Create a class
// @tags classs
// @accept json
// @produce json
// @param class body model.Class true "Class object"
// @success 201 {object} model.Class
// @Security ApiKeyAuth
func (api APIRoutes) CreateClass(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.CreateClass, "creating new class", nil)
	api.Server.CreateClass(ctx)
}

// Handler to get a class by ID
// @router /class/{id} [get]
// @summary Get a class by ID
// @tags classs
// @produce json
// @param id path string true "Class ID"
// @success 200 {object} model.Class
// @Security ApiKeyAuth
func (api APIRoutes) GetClass(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetClass, "fetching  class", nil)
	api.Server.GetClass(ctx)
}

// Handler to get all classs
// @router /class/all [get]
// @summary Get all classs
// @tags classs
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Class
// @Security ApiKeyAuth
func (api APIRoutes) GetClasss(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetClasss, "fetching classs", nil)
	api.Server.GetClasss(ctx)
}

// Handler to get all classs based on filter
// @router /class/filter [get]
// @summary Get all classs based on given filters
// @tags classs
// @produce json
// @param id query string false "id"
// @param school_id query string false "school_id"
// @param active query bool false "active"
// @param created_by query string false "created_by"
// @param updated_by query string false "updated_by"
// @param class_tracher_id query string false "class_tracher_id"
// @param trachers_id query string false "trachers_id"
// @param subjects_id query string false "subjects_id"
// @param floor_number query string false "floor_number"
// @param class_name query string false "class_name"
// @param start_date query string false "start date"
// @param end_date query string false "end date"
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Class
// @Security ApiKeyAuth
func (api APIRoutes) GetClasssByFilter(c *gin.Context) {
	api.Server.GetClassByFilter(c)
}

// Handler to update a class
// @router /class/update/{id} [put]
// @summary Update a class
// @tags classs
// @accept json
// @produce json
// @param id path string true "Class ID"
// @param class body model.Class true "Class object"
// @success 200 {object} model.Class
// @Security ApiKeyAuth
func (api APIRoutes) UpdateClass(c *gin.Context) {
	api.Server.UpdateClass(c)
}

// Handler to delete a class
// @router  /class/delete/{id} [delete]
// @summary Delete a class
// @tags classs
// @param id path string true "Class ID"
// @Security ApiKeyAuth
func (api APIRoutes) DeleteClass(c *gin.Context) {
	api.Server.DeleteClass(c)
}

