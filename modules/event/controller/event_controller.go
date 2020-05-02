package controller

import (
	model "BFC/modules/event/model"
	service "BFC/modules/event/service"
	"strings"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// AddEvent action perfomed and wrire a response to destination
func AddEvent(w http.ResponseWriter, r *http.Request) {

	Event := model.Event{
		EventName:   r.PostFormValue("event_name"),
		EventSource: r.PostFormValue("event_source"),
		EventDetail: r.PostFormValue("event_source"),
		EventType:   strings.ToLower(r.PostFormValue("event_type")),
	}

	// call service to add an event and get response from service
	serviceResponse := service.AddEvent(Event)
	// writing a response
	json.NewEncoder(w).Encode(serviceResponse)
}

// GetEvents is featch a list of all events
func GetEvents(w http.ResponseWriter, r *http.Request) {
	// call service to get all events
	serviceResponse := service.GetAllEvent()

	// writing a response
	json.NewEncoder(w).Encode(serviceResponse)
}

// GetDetail is featch a list of all events
func GetDetail(w http.ResponseWriter, r *http.Request) {

	// get id from request
	vars := mux.Vars(r)

	// call service to get a details
	serviceResponse := service.GetEventDetail(vars["id"])

	// check data is present or not
	if serviceResponse.Data.ID == 0 {
		json.NewEncoder(w).Encode(service.NonEventData())
	} else {
		json.NewEncoder(w).Encode(serviceResponse)
	}

	// writing a response
}

// GetEventTypeCount is featch a count by event type which is unread
func GetEventTypeCount(w http.ResponseWriter, r *http.Request) {
	// call service to get a details
	serviceResponse := service.GetUnreadEventTypeCount()

	// writing a response
	json.NewEncoder(w).Encode(serviceResponse)
}

// GetTotalNotification is featch a count by event type which is unread
func GetTotalNotification(w http.ResponseWriter, r *http.Request) {
	notificationResponse := service.GetUnreadEventCount()
	json.NewEncoder(w).Encode(notificationResponse)

}
