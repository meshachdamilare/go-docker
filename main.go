package main

import (
	"github.com/gorilla/mux"
	"github.com/meshachdamilare/go-docker/router"
	"github.com/meshachdamilare/go-docker/setup"
	"log"
	"net/http"
)

func main() {
	setup.Connection()
	//db := setup.GetDB()
	r := mux.NewRouter()
	router.Register(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
