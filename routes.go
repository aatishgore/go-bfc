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

	// set comon header
	myRouter.Use(commonMiddleware)

	// set prefix of route
	subRouter := myRouter.PathPrefix("/api/").Subrouter()

	// routing list
	subRouter.HandleFunc("/add/incident", event.AddEvent).Methods("POST")
	subRouter.HandleFunc("/incident", event.GetEvents).Methods("GET")
	subRouter.HandleFunc("/incident/{id:[0-9]+}", event.GetDetail).Methods("GET")
	subRouter.HandleFunc("/notification", event.GetEventTypeCount).Methods("GET")

	// create server
	err := http.ListenAndServe(":8080", myRouter)

	// log a error
	log.Fatal(err)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
