package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishal2911/management/model"
)

func (s *Server) CreateSchool(ctx *gin.Context) {

	school := model.School{}

	if err := ctx.BindJSON(&school); err != nil {
		return
	}

	s.PostgressDb.CreateSchool(school)

	ctx.IndentedJSON(http.StatusCreated, school)

	
}
