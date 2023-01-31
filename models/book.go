package models

import (
	"github.com/meshachdamilare/go-docker/setup"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name     string `json:"name"`
	Author   string `json:"author"`
	Language string `json:"language"`
}

func init() {
	setup.Connection()
	db = setup.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int) (*Book, *gorm.DB) {
	var book Book
	db.Where("ID=?", Id).Find(&book)
	return &book, db
}

func DeleteBookById(Id int) Book {
	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return book
}
