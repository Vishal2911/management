package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
)

func (api APIRoutes) LabRouts(routes *gin.Engine) {
	// Define routes
	group := routes.Group("lab")
	{
		group.POST("/create", api.AuthMiddlewareComplete(), api.CreateLab)
		group.GET("/:id", api.AuthMiddlewareComplete(), api.GetLab)
		group.GET("/all", api.AuthMiddlewareComplete(), api.GetLabs)
		group.GET("/filter", api.GetLabsByFilter)
		group.PUT("/update/:id", api.AuthMiddlewareComplete(), api.UpdateLab)
		group.DELETE("/delete/:id", api.AuthMiddlewareComplete(), api.DeleteLab)
	}

}

// Handler to create a lab
// @router /lab/create [post]
// @summary Create a lab
// @tags labs
// @accept json
// @produce json
// @param lab body model.Lab true "Lab object"
// @success 201 {object} model.Lab
// @Security ApiKeyAuth
func (api APIRoutes) CreateLab(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.CreateLab, "creating new lab", nil)
	api.Server.CreateLab(ctx)
}

// Handler to get a lab by ID
// @router /lab/{id} [get]
// @summary Get a lab by ID
// @tags labs
// @produce json
// @param id path string true "Lab ID"
// @success 200 {object} model.Lab
// @Security ApiKeyAuth
func (api APIRoutes) GetLab(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetLab, "fetching  lab", nil)
	api.Server.GetLab(ctx)
}

// Handler to get all labs
// @router /lab/all [get]
// @summary Get all labs
// @tags labs
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Lab
// @Security ApiKeyAuth
func (api APIRoutes) GetLabs(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetLabs, "fetching labs", nil)
	api.Server.GetLabs(ctx)
}

// Handler to get all labs based on filter
// @router /lab/filter [get]
// @summary Get all labs based on given filters
// @tags labs
// @produce json
// @param id query string false "id"
// @param school_id query string false "school_id"
// @param room_id query string false "room_id"
// @param number_of_equipments query int false "number_of_equipments"
// @param lab_assistant_id query string false "lab_assistant_id"
// @param active query bool false "active"
// @param created_by query string false "created_by"
// @param updated_by query string false "updated_by"
// @param lab_name query string false "lab_name"
// @param start_date query string false "start date"
// @param end_date query string false "end date"
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Lab
// @Security ApiKeyAuth
func (api APIRoutes) GetLabsByFilter(c *gin.Context) {
	api.Server.GetLabByFilter(c)
}

// Handler to update a lab
// @router /lab/update/{id} [put]
// @summary Update a lab
// @tags labs
// @accept json
// @produce json
// @param id path string true "Lab ID"
// @param lab body model.Lab true "Lab object"
// @success 200 {object} model.Lab
// @Security ApiKeyAuth
func (api APIRoutes) UpdateLab(c *gin.Context) {
	api.Server.UpdateLab(c)
}

// Handler to delete a lab
// @router  /lab/delete/{id} [delete]
// @summary Delete a lab
// @tags labs
// @param id path string true "Lab ID"
// @Security ApiKeyAuth
func (api APIRoutes) DeleteLab(c *gin.Context) {
	api.Server.DeleteLab(c)
}
