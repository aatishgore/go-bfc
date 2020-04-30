package main

import (
	event "BFC/modules/event/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Init() {

	// create rounter object to handel a route
	myRouter := mux.NewRouter().StrictSlash(true)

	// set prefix of route
	myRouter.PathPrefix("/api/")

	// routing list
	myRouter.HandleFunc("/list", event.GetList).Methods("GET")

	// create server
	err := http.ListenAndServe(":8080", myRouter)

	log.Fatal(err)

}
