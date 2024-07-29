package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vishal2911/management/controller"
	"github.com/vishal2911/management/store"
)

type ApiRouts struct {
	Server controller.ServerOperations
}

func (api *ApiRouts) StartApp(router *gin.Engine, server controller.Server) {
	api.Server = &server
	api.Server.NewServer(store.Postgress{})
	
	// user routs
	api.UserRouts(router)

}
