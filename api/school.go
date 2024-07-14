package api

import "github.com/gin-gonic/gin"

func (api *ApiRouts) School() {


	Group :=api.router.Group("school")
    {
        Group.POST("/create", api.CreateSchool)
    }


}

func (api *ApiRouts) CreateSchool(ctx *gin.Context) {

	api.Server.CreateSchool(ctx)
}
