package main

import (
	event "BFC/modules/event/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/**
 * Init function initialize a all route and create http server on port 8080
 */

// Init is ...
func Init() {

	// create rounter object to handel a route
	myRouter := mux.NewRouter()

	// set prefix of route
	myRouter.PathPrefix("/api/")

	// routing list
	myRouter.HandleFunc("/add/event", event.AddEvent).Methods("POST")
	myRouter.HandleFunc("/list", event.GetEvents).Methods("GET")
	// myRouter.HandleFunc("/new-event-count", event.GetUnreadEventCount).Methods("GET")
	// myRouter.HandleFunc("/{id}/get", event.GetEvent).Methods("GET")

	// create server
	err := http.ListenAndServe(":8080", myRouter)

	// log a error
	log.Fatal(err)
}
