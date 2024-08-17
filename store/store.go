package store

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgress struct {
	DB *gorm.DB
}

func (store *Postgress) NewStore() error {
	dsn := "host=localhost user=vishal password=password dbname=manage port=5432 sslmode=disable"

	util.Log(model.LogLevelInfo, model.StorePackage, model.NewStore, "creating new store", nil)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, "error while creating store", err)
		return err
	} else {
		store.DB = db
	}

	err = db.AutoMigrate(
		model.User{},
		model.School{},
		model.Class{},
		model.Teacher{},
		model.Book{},
		model.Room{},
		model.Lab{},
		model.Author{},
		model.Publisher{},
		model.Subject{},
	)
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, "error while running automigration", err)
		return err
	}

	fmt.Printf("db = %v\n", db)
	return nil
}

type SoteOperations interface {
	NewStore() error
	CreateUser(user *model.User) error
	GetUsers() ([]model.User, error)
	GetUser(uuid.UUID) (model.User, error)
	GetUserByFilter(filter map[string]interface{}) ([]model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(userID string) error
	SingIn(userSignIn model.UserSignIn) (*model.User, error)
	SignUp(user *model.User) error

	CreateSchool(school *model.School) error
	GetSchools() ([]model.School, error)
	GetSchool(uuid.UUID) (model.School, error)
	GetSchoolByFilter(filter map[string]interface{}) ([]model.School, error)
	UpdateSchool(school *model.School) error
	DeleteSchool(schoolID string) error


	CreateClass(class *model.Class) error
	GetClasss() ([]model.Class, error)
	GetClass(uuid.UUID) (model.Class, error)
	GetClassByFilter(filter map[string]interface{}) ([]model.Class, error)
	UpdateClass(class *model.Class) error
	DeleteClass(classID string) error


	CreateBook(book *model.Book) error
	GetBooks() ([]model.Book, error)
	GetBook(uuid.UUID) (model.Book, error)
	GetBookByFilter(filter map[string]interface{}) ([]model.Book, error)
	UpdateBook(book *model.Book) error
	DeleteBook(bookID string) error




	CreateLab(lab *model.Lab) error
	GetLabs() ([]model.Lab, error)
	GetLab(uuid.UUID) (model.Lab, error)
	GetLabByFilter(filter map[string]interface{}) ([]model.Lab, error)
	UpdateLab(lab *model.Lab) error
	DeleteLab(labID string) error	
	
	
	CreateAuthor(author *model.Author) error
	GetAuthors() ([]model.Author, error)
	GetAuthor(uuid.UUID) (model.Author, error)
	GetAuthorByFilter(filter map[string]interface{}) ([]model.Author, error)
	UpdateAuthor(author *model.Author) error
	DeleteAuthor(authorID string) error
	

	CreatePublisher(publisher *model.Publisher) error
	GetPublishers() ([]model.Publisher, error)
	GetPublisher(uuid.UUID) (model.Publisher, error)
	GetPublisherByFilter(filter map[string]interface{}) ([]model.Publisher, error)
	UpdatePublisher(publisher *model.Publisher) error
	DeletePublisher(publisherID string) error

	CreateRoom(room *model.Room) error
	GetRooms() ([]model.Room, error)
	GetRoom(uuid.UUID) (model.Room, error)
	GetRoomByFilter(filter map[string]interface{}) ([]model.Room, error)
	UpdateRoom(room *model.Room) error
	DeleteRoom(roomID string) error

	CreateTeacher(teacher *model.Teacher) error
	GetTeachers() ([]model.Teacher, error)
	GetTeacher(uuid.UUID) (model.Teacher, error)
	GetTeacherByFilter(filter map[string]interface{}) ([]model.Teacher, error)
	UpdateTeacher(teacher *model.Teacher) error
	DeleteTeacher(teacherID string) error

	CreateSubject(subject *model.Subject) error
	GetSubjects() ([]model.Subject, error)
	GetSubject(uuid.UUID) (model.Subject, error)
	GetSubjectByFilter(filter map[string]interface{}) ([]model.Subject, error)
	UpdateSubject(subject *model.Subject) error
	DeleteSubject(subjectID string) error
}
