package store

import (
	"github.com/vishal2911/management/model"
	"gorm.io/gorm"
)

func (store *Postgress) CreateSchool(school model.School) *gorm.DB {

	return store.DB.Create(&school)

}
