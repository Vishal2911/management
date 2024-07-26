package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
)

func (api ApiRouts) UserRouts(routes *gin.Engine) {
	group := routes.Group("user")
	{
		group.POST("/create", api.CreateUser)
	}

}

func (api ApiRouts) CreateUser(ctx *gin.Context) {

	util.Log(model.LogLevelInfo, model.ApiPackage, model.CreateUser, "creating new user", nil)
	api.Server.CreateUser(ctx)
}
