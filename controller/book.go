package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/meshachdamilare/go-docker/models"
	"github.com/meshachdamilare/go-docker/utils"
	"net/http"
	"strconv"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	Book := &models.Book{}
	utils.ParseBody(r, &Book)
	b := Book.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error converting the Id")
	}
	book, _ := models.GetBookById(int(Id))
	if book.ID == 0 {
		msg := fmt.Sprintf("Book Id: %d not found", Id)
		w.Write([]byte(msg))
		return
	}
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookUpdate := models.Book{}
	utils.ParseBody(r, &bookUpdate)
	fmt.Println(bookUpdate)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error converting the Id")
	}
	bookToUpdate, db := models.GetBookById(int(Id))
	if bookUpdate.Name != "" {
		bookToUpdate.Name = bookUpdate.Name
	}
	if bookUpdate.Author != "" {
		bookToUpdate.Author = bookUpdate.Author
	}
	if bookUpdate.Language != "" {
		bookToUpdate.Language = bookUpdate.Language
	}
	db.Save(&bookToUpdate)

	res, _ := json.Marshal(bookToUpdate)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error converting the Id")
	}
	book := models.DeleteBookById(int(Id))
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
