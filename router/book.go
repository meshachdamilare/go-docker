package router

import (
	"github.com/gorilla/mux"
	"github.com/meshachdamilare/go-docker/controller"
)

var Register = func(r *mux.Router) {
	r.HandleFunc("/books", controller.CreateBook).Methods("POST")
	r.HandleFunc("/books", controller.GetAllBooks).Methods("GET")
	r.HandleFunc("/books/{bookId}", controller.GetBook).Methods("GET")
	r.HandleFunc("/books/{bookId}", controller.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{bookId}", controller.DeleteBook).Methods("DELETE")
}
