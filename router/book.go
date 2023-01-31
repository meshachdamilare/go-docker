package router

import (
	"github.com/gorilla/mux"
	"github.com/meshachdamilare/go-docker/controller"
)

var Register = func(r *mux.Router) {
	r.HandleFunc("/createbook", controller.CreateBook).Methods("POST")
}
