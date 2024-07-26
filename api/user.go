package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
)

func (api ApiRouts) UserRouts(routes *gin.Engine) {
	// Define routes
	group := routes.Group("user")
	{
		group.POST("/create", api.CreateUser)
		group.GET("/:id", api.GetUser)
		group.GET("/all", api.GetUsers)
	}

}

// Handler to create a user
// @router /user/create [post]
// @summary Create a user
// @tags users
// @accept json
// @produce json
// @param user body model.User true "User object"
// @success 201 {object} model.User
func (api ApiRouts) CreateUser(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.CreateUser, "creating new user", nil)
	api.Server.CreateUser(ctx)
}

// Handler to get a user by ID
// @router /user/{id} [get]
// @summary Get a user by ID
// @tags users
// @produce json
// @param id path string true "User ID"
// @success 200 {object} model.User
func (api ApiRouts) GetUser(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetUser, "fetching  user", nil)
	api.Server.GetUser(ctx)
}

// Handler to get all users
// @router /user/all [get]
// @summary Get all users
// @tags users
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.User
func (api ApiRouts) GetUsers(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetUsers, "fetching users", nil)
	api.Server.GetUsers(ctx)
}