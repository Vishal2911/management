package api

import (
	"github.com/vishal2911/management/controller"
	"github.com/vishal2911/management/store"
)

type ApiRouts struct {
	Server controller.ServerOperations
}

func (api *ApiRouts) StartApp(server controller.Server) {
	api.Server = &server
	api.Server.NewServer(store.Postgress{})

}
