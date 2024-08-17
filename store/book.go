package store

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vishal2911/management/model"
	"github.com/vishal2911/management/util"
	"gorm.io/gorm"
)

func (store Postgress) CreateBook(book *model.Book) error {
	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateBook, "creating new book", nil)
	response := store.DB.Create(book)
	if response.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, "error while creating new book", response.Error)
		return response.Error
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateBook, "Created new book", book)
	return nil
}

func (store Postgress) GetBooks() ([]model.Book, error) {

	books := []model.Book{}
	util.Log(model.LogLevelInfo, model.StorePackage, model.GetBooks, "fetching records of book from db", nil)
	if err := store.DB.Find(&books).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetBooks, "error while fetching books from db", err)
		return books, err
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetBooks, "records of book from db", books)
	return books, nil
}

func (store Postgress) GetBook(bookID uuid.UUID) (model.Book, error) {

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetBook, "fetching records of book from db", nil)
	var book model.Book
	if err := store.DB.First(&book, "id = ?", bookID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			util.Log(model.LogLevelError, model.StorePackage, model.GetBook, "book record not found", err)
		} else {
			util.Log(model.LogLevelError, model.StorePackage, model.GetBook, "error while fetching book from db", err)
		}
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetBook, "records of book from db", book)
	return book, nil
}

func (store Postgress) GetBookByFilter(filter map[string]interface{}) ([]model.Book, error) {

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetBookByFilter, "fetching records of book from db", nil)
	var books []model.Book
	query := store.DB

	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == model.StartDate || key == model.EndDate {
			continue
		}
		util.Log(model.LogLevelInfo, model.StorePackage, model.GetBookByFilter,
			"filters key", key+" value = "+fmt.Sprintf("%v", value))
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	setLimitAndPage(filter, query)
	setDateRangeFilter(filter, query)

	err := query.Find(&books).Error
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetBookByFilter,
			"error while reading book data", err)
		return nil, fmt.Errorf("error while fetching book records from DB, err = %v", err)
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetBookByFilter, "records of books from db", books)
	return books, nil
}

func (store Postgress) UpdateBook(book *model.Book) error {

	util.Log(model.LogLevelInfo, model.StorePackage, model.UpdateBook, "updating book data", *book)
	resp := store.DB.Save(book)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.UpdateBook,
			"error while updating book data", resp.Error)
		return fmt.Errorf("error while updating book record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.UpdateBook,
		"successfully updated book", nil)
	return nil
}

// DeleteBook is used to delete record by given bookID
func (store Postgress) DeleteBook(bookID string) error {

	var book model.Book
	util.Log(model.LogLevelInfo, model.StorePackage, model.DeleteBook, "deleting book data", bookID)
	if err := store.DB.First(&book, `"id" = '`+bookID+`'`).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.DeleteBook,
			"error while deleting book data", err)
		return fmt.Errorf("book not found for given id, ID = %v", bookID)
	}
	resp := store.DB.Delete(book)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting book record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.DeleteBook,
		"successfully deleted book", nil)
	return nil
}

