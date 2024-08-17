package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
)

func (api APIRoutes) SubjectRouts(routes *gin.Engine) {
	// Define routes
	group := routes.Group("subject")
	{
		group.POST("/create", api.AuthMiddlewareComplete(), api.CreateSubject)
		group.GET("/:id", api.AuthMiddlewareComplete(), api.GetSubject)
		group.GET("/all", api.AuthMiddlewareComplete(), api.GetSubjects)
		group.GET("/filter", api.GetSubjectsByFilter)
		group.PUT("/update/:id", api.AuthMiddlewareComplete(), api.UpdateSubject)
		group.DELETE("/delete/:id", api.AuthMiddlewareComplete(), api.DeleteSubject)
	}

}

// Handler to create a subject
// @router /subject/create [post]
// @summary Create a subject
// @tags subjects
// @accept json
// @produce json
// @param subject body model.Subject true "Subject object"
// @success 201 {object} model.Subject
// @Security ApiKeyAuth
func (api APIRoutes) CreateSubject(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.CreateSubject, "creating new subject", nil)
	api.Server.CreateSubject(ctx)
}

// Handler to get a subject by ID
// @router /subject/{id} [get]
// @summary Get a subject by ID
// @tags subjects
// @produce json
// @param id path string true "Subject ID"
// @success 200 {object} model.Subject
// @Security ApiKeyAuth
func (api APIRoutes) GetSubject(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetSubject, "fetching  subject", nil)
	api.Server.GetSubject(ctx)
}

// Handler to get all subjects
// @router /subject/all [get]
// @summary Get all subjects
// @tags subjects
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Subject
// @Security ApiKeyAuth
func (api APIRoutes) GetSubjects(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetSubjects, "fetching subjects", nil)
	api.Server.GetSubjects(ctx)
}

// Handler to get all subjects based on filter
// @router /subject/filter [get]
// @summary Get all subjects based on given filters
// @tags subjects
// @produce json
// @param id query string false "id"
// @param school_id query string false "school_id"
// @param book_id query string false "book_id"
// @param active query bool false "active"
// @param created_by query string false "created_by"
// @param updated_by query string false "updated_by"
// @param name query string false "name"
// @param start_date query string false "start date"
// @param end_date query string false "end date"
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Subject
// @Security ApiKeyAuth
func (api APIRoutes) GetSubjectsByFilter(c *gin.Context) {
	api.Server.GetSubjectByFilter(c)
}

// Handler to update a subject
// @router /subject/update/{id} [put]
// @summary Update a subject
// @tags subjects
// @accept json
// @produce json
// @param id path string true "Subject ID"
// @param subject body model.Subject true "Subject object"
// @success 200 {object} model.Subject
// @Security ApiKeyAuth
func (api APIRoutes) UpdateSubject(c *gin.Context) {
	api.Server.UpdateSubject(c)
}

// Handler to delete a subject
// @router  /subject/delete/{id} [delete]
// @summary Delete a subject
// @tags subjects
// @param id path string true "Subject ID"
// @Security ApiKeyAuth
func (api APIRoutes) DeleteSubject(c *gin.Context) {
	api.Server.DeleteSubject(c)
}

