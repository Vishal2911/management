package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
)

func (api APIRoutes) SchoolRouts(routes *gin.Engine) {
	// Define routes
	group := routes.Group("school")
	{
		group.POST("/create", api.AuthMiddlewareComplete(), api.CreateSchool)
		group.GET("/:id", api.AuthMiddlewareComplete(), api.GetSchool)
		group.GET("/all", api.AuthMiddlewareComplete(), api.GetSchools)
		group.GET("/filter", api.GetSchoolsByFilter)
		group.PUT("/update/:id", api.AuthMiddlewareComplete(), api.UpdateSchool)
		group.DELETE("/delete/:id", api.AuthMiddlewareComplete(), api.DeleteSchool)
	}

}

// Handler to create a school
// @router /school/create [post]
// @summary Create a school
// @tags schools
// @accept json
// @produce json
// @param school body model.School true "School object"
// @success 201 {object} model.School
// @Security ApiKeyAuth
func (api APIRoutes) CreateSchool(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.CreateSchool, "creating new school", nil)
	api.Server.CreateSchool(ctx)
}

// Handler to get a school by ID
// @router /school/{id} [get]
// @summary Get a school by ID
// @tags schools
// @produce json
// @param id path string true "School ID"
// @success 200 {object} model.School
// @Security ApiKeyAuth
func (api APIRoutes) GetSchool(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetSchool, "fetching  school", nil)
	api.Server.GetSchool(ctx)
}

// Handler to get all schools
// @router /school/all [get]
// @summary Get all schools
// @tags schools
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.School
// @Security ApiKeyAuth
func (api APIRoutes) GetSchools(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetSchools, "fetching schools", nil)
	api.Server.GetSchools(ctx)
}

// Handler to get all schools based on filter
// @router /school/filter [get]
// @summary Get all schools based on given filters
// @tags schools
// @produce json
// @param domain query string false "email"
// @param id query string false "id"
// @param principle_id query string false "principle_id"
// @param active query bool false "active"
// @param created_by query string false "created_by"
// @param updated_by query string false "updated_by"
// @param admin_id query string false "admin_id"
// @param director_id query string false "director_id"
// @param hostel_id query string false "hostel_id"
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
// @success 200 {array} model.School
// @Security ApiKeyAuth
func (api APIRoutes) GetSchoolsByFilter(c *gin.Context) {
	api.Server.GetSchoolByFilter(c)
}

// Handler to update a school
// @router /school/update/{id} [put]
// @summary Update a school
// @tags schools
// @accept json
// @produce json
// @param id path string true "School ID"
// @param school body model.School true "School object"
// @success 200 {object} model.School
// @Security ApiKeyAuth
func (api APIRoutes) UpdateSchool(c *gin.Context) {
	api.Server.UpdateSchool(c)
}

// Handler to delete a school
// @router  /school/delete/{id} [delete]
// @summary Delete a school
// @tags schools
// @param id path string true "School ID"
// @Security ApiKeyAuth
func (api APIRoutes) DeleteSchool(c *gin.Context) {
	api.Server.DeleteSchool(c)
}

