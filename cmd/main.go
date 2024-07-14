package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vishal2911/management/api"
	"github.com/vishal2911/management/controller"
)

func main() {
	api := api.ApiRouts{}


	router := gin.Default()
	server := controller.Server{}
	api.StartApp(router , server)
	// fmt.Printf("main server = %v\n", api)
}
