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

func (api ApiRouts) CreateUser(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.CreateUser, "creating new user", nil)
	api.Server.CreateUser(ctx)
}

func (api ApiRouts) GetUser(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetUser, "fetching  user", nil)
	api.Server.GetUser(ctx)
}

func (api ApiRouts) GetUsers(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetUsers, "fetching users", nil)
	api.Server.GetUsers(ctx)
}
