package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/store"
	"github.com/vishal2911/management/util"
)

type Server struct {
	PostgressDb store.SoteOperations
}

func (s *Server) NewServer(pgstore store.Postgress) {
	util.SetLogger()
	util.Logger.Infof("Logger setup Done....\n")
	s.PostgressDb = &pgstore
	err := s.PostgressDb.NewStore()
	if err != nil {
		util.Logger.Errorf("error while creating store connection, err = %v\n", err)
		util.Log(model.LogLevelError, model.Controller, "NewServer", "error while creating store connection", err)
	} else {
		util.Logger.Infof("Connected with store\n")
		util.Log(model.LogLevelInfo, model.Controller, model.NewServer, "Connected with store", nil)
	}

	log.Printf("server = %v\n", s)
}

type ServerOperations interface {
	NewServer(pgstore store.Postgress)
	CreateUser(ctx *gin.Context)
}
