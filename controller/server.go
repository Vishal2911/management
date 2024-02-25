package controller

import (
	"fmt"

	"github.com/vishal2911/management/store"
)

type Server struct {
	PostgressDb store.SoteOperations
}

func (s *Server) NewServer(pgstore store.Postgress) {
	s.PostgressDb = &pgstore
	s.PostgressDb.NewStore()
	fmt.Printf("server = %v\n", s)
}

type ServerOperations interface{
	NewServer(pgstore store.Postgress)
}
