package models

import (
	"go-bookstore/pkg/config"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name" example:"A Tale of Two Cities" format:"string"`
	Author      string `json:"author" example:"Charles Dickens" format:"string"`
	Publication string `json:"publication" example:"Harper Collins" format:"string"`
}

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func (book *Book) CreateBook() *Book {
	db.Create(&book)
	return book
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return book
}
