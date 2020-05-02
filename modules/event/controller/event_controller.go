package controller

import (
	model "BFC/modules/event/model"
	service "BFC/modules/event/service"
	socket "BFC/utilities"
	"fmt"
	"io/ioutil"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// AddEvent action perfomed and wrire a response to destination
func AddEvent(w http.ResponseWriter, r *http.Request) {

	var Event model.Event

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error in reading body")
	}

	err = json.Unmarshal(body, &Event)

	// call service to add an event and get response from service
	serviceResponse := service.AddEvent(Event)

	notificationResponse := service.GetUnreadEventCount()
	serviceString, _ := json.Marshal(notificationResponse)
	socket.SendNotification(
		"success",
		string(serviceString),
	)
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
