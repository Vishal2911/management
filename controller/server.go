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

	//User controllers
	CreateUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	GetUserByFilter(ctx *gin.Context)
	UpdateUser(c *gin.Context) error
	DeleteUser(c *gin.Context) error
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)

	// School controllers
	CreateSchool(ctx *gin.Context)
	GetSchool(ctx *gin.Context)
	GetSchools(ctx *gin.Context)
	GetSchoolByFilter(ctx *gin.Context)
	UpdateSchool(c *gin.Context) error
	DeleteSchool(c *gin.Context) error

	// Class controllers
	CreateClass(ctx *gin.Context)
	GetClass(ctx *gin.Context)
	GetClasss(ctx *gin.Context)
	GetClassByFilter(ctx *gin.Context)
	UpdateClass(c *gin.Context) error
	DeleteClass(c *gin.Context) error

		// Teacher controllers
		CreateTeacher(ctx *gin.Context)
		GetTeacher(ctx *gin.Context)
		GetTeachers(ctx *gin.Context)
		GetTeacherByFilter(ctx *gin.Context)
		UpdateTeacher(c *gin.Context) error
		DeleteTeacher(c *gin.Context) error



			// Author controllers
			CreateAuthor(ctx *gin.Context)
			GetAuthor(ctx *gin.Context)
			GetAuthors(ctx *gin.Context)
			GetAuthorByFilter(ctx *gin.Context)
			UpdateAuthor(c *gin.Context) error
			DeleteAuthor(c *gin.Context) error


				// Book controllers
		CreateBook(ctx *gin.Context)
		GetBook(ctx *gin.Context)
		GetBooks(ctx *gin.Context)
		GetBookByFilter(ctx *gin.Context)
		UpdateBook(c *gin.Context) error
		DeleteBook(c *gin.Context) error

			// Lab controllers
			CreateLab(ctx *gin.Context)
			GetLab(ctx *gin.Context)
			GetLabs(ctx *gin.Context)
			GetLabByFilter(ctx *gin.Context)
			UpdateLab(c *gin.Context) error
			DeleteLab(c *gin.Context) error

				// Publisher controllers
		CreatePublisher(ctx *gin.Context)
		GetPublisher(ctx *gin.Context)
		GetPublishers(ctx *gin.Context)
		GetPublisherByFilter(ctx *gin.Context)
		UpdatePublisher(c *gin.Context) error
		DeletePublisher(c *gin.Context) error

			// Room controllers
			CreateRoom(ctx *gin.Context)
			GetRoom(ctx *gin.Context)
			GetRooms(ctx *gin.Context)
			GetRoomByFilter(ctx *gin.Context)
			UpdateRoom(c *gin.Context) error
			DeleteRoom(c *gin.Context) error


					// Subject controllers
					CreateSubject(ctx *gin.Context)
					GetSubject(ctx *gin.Context)
					GetSubjects(ctx *gin.Context)
					GetSubjectByFilter(ctx *gin.Context)
					UpdateSubject(c *gin.Context) error
					DeleteSubject(c *gin.Context) error
}
