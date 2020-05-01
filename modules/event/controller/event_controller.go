package controller

import (
	service "BFC/modules/event/service"
	socket "BFC/utilities"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// AddEvent is ...
// Add event action perfomed and wrire a response to destination
func AddEvent(w http.ResponseWriter, r *http.Request) {
	// call service to add an event and get response from service
	serviceResponse := service.AddEvent()
	// writing a response
	json.NewEncoder(w).Encode(serviceResponse)
}

// GetEvents is ...
// GetEvents is featch a list of all events
func GetEvents(w http.ResponseWriter, r *http.Request) {
	// call service to get all events
	serviceResponse := service.GetAllEvent()

	// writing a response
	json.NewEncoder(w).Encode(serviceResponse)
}

// GetDetail is ...
// GetEvents is featch a list of all events
func GetDetail(w http.ResponseWriter, r *http.Request) {

	// get id from request
	vars := mux.Vars(r)

	// call service to get a details
	serviceResponse := service.GetEventDetail(vars["id"])
	serviceString, _ := json.Marshal(serviceResponse)
	socket.SendNotification(
		"success",
		string(serviceString),
	)
	if serviceResponse.Data.ID == 0 {
		json.NewEncoder(w).Encode(service.NonEventData())
	} else {
		json.NewEncoder(w).Encode(serviceResponse)
	}

	// writing a response
}

// GetEventTypeCount is ...
// GetEventTypeCount is featch a count by event type which is unread
func GetEventTypeCount(w http.ResponseWriter, r *http.Request) {
	// call service to get a details
	serviceResponse := service.GetUnreadEventTypeCount()

	// writing a response
	json.NewEncoder(w).Encode(serviceResponse)
}
