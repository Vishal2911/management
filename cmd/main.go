package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vishal2911/management/api"
	"github.com/vishal2911/management/controller"
)

// @title Managment
// @version 1.0
// @description API for managing School operations
// @host localhost:8000
// @BasePath /
// @schemes http https
func main() {
	api := api.APIRoutes{}
	controller := controller.Server{}
	routes := gin.Default()
	api.StartApp(routes, controller)

	routes.Run(":8000")
	// fmt.Printf("main server = %v\n", api)
}
