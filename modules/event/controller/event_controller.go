package controller

import (
	model "BFC/modules/event/model"
	service "BFC/modules/event/service"
	"io/ioutil"
	"log"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// AddEvent action perfomed and wrire a response to destination
func AddEvent(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling Add event api")
	var Event model.Event

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln("Unable to read a data")
	}

	err = json.Unmarshal(body, &Event)

	// call service to add an event and get response from service
	serviceResponse := service.AddEvent(Event)
	// writing a response
	log.Println("sending a response after add an event")
	json.NewEncoder(w).Encode(serviceResponse)
}

// GetEvents is featch a list of all events
func GetEvents(w http.ResponseWriter, r *http.Request) {
	log.Println("Call list of events")
	// call service to get all events
	serviceResponse := service.GetAllEvent()

	// writing a response
	log.Println("Send list of event")
	json.NewEncoder(w).Encode(serviceResponse)
}

// GetDetail is featch a list of all events
func GetDetail(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling get event detail api")
	// get id from request
	vars := mux.Vars(r)

	// call service to get a details
	serviceResponse := service.GetEventDetail(vars["id"])

	log.Println("Calling get event detail api")
	// writing a response
	json.NewEncoder(w).Encode(serviceResponse)
}

// GetEventTypeCount is featch a count by event type which is unread
func GetEventTypeCount(w http.ResponseWriter, r *http.Request) {
	log.Println("Call api to get event type count")
	// call service to get a details
	serviceResponse := service.GetUnreadEventTypeCount()

	// writing a response
	log.Println("Send count of event by type")
	json.NewEncoder(w).Encode(serviceResponse)
}

// GetTotalNotification is featch a count by event type which is unread
func GetTotalNotification(w http.ResponseWriter, r *http.Request) {
	notificationResponse := service.GetUnreadEventCount()
	json.NewEncoder(w).Encode(notificationResponse)

}
