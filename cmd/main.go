package main

import (
	"github.com/vishal2911/management/api"
	"github.com/vishal2911/management/controller"
)

func main() {
	api := api.ApiRouts{}
	api.StartApp(controller.Server{})
	// fmt.Printf("main server = %v\n", api)
}
