package controller

import (
	service "BFC/modules/event/service"
	"fmt"
	"net/http"
)

// AddEvent is ...
// Add event action perfomed and wrire a response to destination
func AddEvent(w http.ResponseWriter, r *http.Request) {
	// call service to add an event and get response from service
	serviceResponse := service.AddEvent()

	// writing a response
	fmt.Fprintln(w, serviceResponse)
}

// GetEvents is ...
// GetEvents is featch a list of all events
func GetEvents(w http.ResponseWriter, r *http.Request) {
	serviceResponse := service.GetAllEvent()

	// writing a response
	fmt.Fprintln(w, serviceResponse)
}
