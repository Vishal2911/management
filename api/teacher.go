package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
)

func (api APIRoutes) TeacherRouts(routes *gin.Engine) {
	// Define routes
	group := routes.Group("teacher")
	{
		group.POST("/create", api.AuthMiddlewareComplete(), api.CreateTeacher)
		group.GET("/:id", api.AuthMiddlewareComplete(), api.GetTeacher)
		group.GET("/all", api.AuthMiddlewareComplete(), api.GetTeachers)
		group.GET("/filter", api.GetTeachersByFilter)
		group.PUT("/update/:id", api.AuthMiddlewareComplete(), api.UpdateTeacher)
		group.DELETE("/delete/:id", api.AuthMiddlewareComplete(), api.DeleteTeacher)
	}

}

// Handler to create a teacher
// @router /teacher/create [post]
// @summary Create a teacher
// @tags teachers
// @accept json
// @produce json
// @param teacher body model.Teacher true "Teacher object"
// @success 201 {object} model.Teacher
// @Security ApiKeyAuth
func (api APIRoutes) CreateTeacher(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.CreateTeacher, "creating new teacher", nil)
	api.Server.CreateTeacher(ctx)
}

// Handler to get a teacher by ID
// @router /teacher/{id} [get]
// @summary Get a teacher by ID
// @tags teachers
// @produce json
// @param id path string true "Teacher ID"
// @success 200 {object} model.Teacher
// @Security ApiKeyAuth
func (api APIRoutes) GetTeacher(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetTeacher, "fetching  teacher", nil)
	api.Server.GetTeacher(ctx)
}

// Handler to get all teachers
// @router /teacher/all [get]
// @summary Get all teachers
// @tags teachers
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Teacher
// @Security ApiKeyAuth
func (api APIRoutes) GetTeachers(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetTeachers, "fetching teachers", nil)
	api.Server.GetTeachers(ctx)
}

// Handler to get all teachers based on filter
// @router /teacher/filter [get]
// @summary Get all teachers based on given filters
// @tags teachers
// @produce json
// @param email query string false "email"
// @param id query string false "id"
// @param school_id query string false "school_id"
// @param password query string false "password"
// @param active query bool false "active"
// @param created_by query string false "created_by"
// @param updated_by query string false "updated_by"
// @param first_name query string false "first_name"
// @param middle_name query string false "middle_name"
// @param last_name query string false "last_name"
// @param salary query int false "salary"
// @param lane query string false "lane"
// @param village query string false "village"
// @param city query string false "city"
// @param district query string false "district"
// @param pincode query int false "pincode"
// @param state query string false "state"
// @param subjects_id query string false "subjects id"
// @param joining_date query string false "joining_date"
// @param skills query string false "skills"
// @param classes_id query string false "classes_id"
// @param teacher_name query string false "teacher_name"
// @param start_date query string false "start date"
// @param end_date query string false "end date"
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Teacher
// @Security ApiKeyAuth
func (api APIRoutes) GetTeachersByFilter(c *gin.Context) {
	api.Server.GetTeacherByFilter(c)
}

// Handler to update a teacher
// @router /teacher/update/{id} [put]
// @summary Update a teacher
// @tags teachers
// @accept json
// @produce json
// @param id path string true "Teacher ID"
// @param teacher body model.Teacher true "Teacher object"
// @success 200 {object} model.Teacher
// @Security ApiKeyAuth
func (api APIRoutes) UpdateTeacher(c *gin.Context) {
	api.Server.UpdateTeacher(c)
}

// Handler to delete a teacher
// @router  /teacher/delete/{id} [delete]
// @summary Delete a teacher
// @tags teachers
// @param id path string true "Teacher ID"
// @Security ApiKeyAuth
func (api APIRoutes) DeleteTeacher(c *gin.Context) {
	api.Server.DeleteTeacher(c)
}

