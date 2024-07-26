package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vishal2911/management/api"
	"github.com/vishal2911/management/controller"
)

func main() {
	api := api.ApiRouts{}
	controller := controller.Server{}
	routes:= gin.Default()
	api.StartApp(routes , controller)
	

	routes.Run(":8000")
	// fmt.Printf("main server = %v\n", api)
}
