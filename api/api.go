package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vishal2911/management/controller"
	_ "github.com/vishal2911/management/docs"
	"github.com/vishal2911/management/store"
)

type ApiRouts struct {
	Server controller.ServerOperations
}

func (api *ApiRouts) StartApp(router *gin.Engine, server controller.Server) {
	api.Server = &server
	api.Server.NewServer(store.Postgress{})

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// user routs
	api.UserRouts(router)

}
