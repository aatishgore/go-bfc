package service

import (
	"BFC/modules/event/model"
	"BFC/utilities"
	socket "BFC/utilities"
	"encoding/json"
	"log"
)

// EventResponse is set default event response
type EventResponse struct {
	Status bool        `json:"status"`
	Data   model.Event `json:"data"`
}

//EventsResponse is set response for event list
type EventsResponse struct {
	Status bool         `json:"status"`
	Data   model.Events `json:"data"`
}

// EventsTypeResponse is set response for event by type
type EventsTypeResponse struct {
	Status bool          `json:"status"`
	Data   []model.Count `json:"data"`
}

// UnreadMessageCount is giv count of unread message
type UnreadMessageCount struct {
	Status bool     `json:"status"`
	Data   []uint64 `json:"data"`
}

// response variable
var response string

// AddEvent create an event and sent response to service.
func AddEvent(Event model.Event) EventResponse {
	log.Println("Add event service initialize")
	RequiredEventType := []string{"warning", "error"}

	// call model to create a event
	eventData := model.AddEvent(Event)

	_, found := Find(RequiredEventType, Event.EventType)

	if found {
		// sending notification
		go SendNotification()

		// call instance api
		utilitieEventString, _ := json.Marshal(eventData)
		go utilities.CallInstanceAPI(utilitieEventString)

	}

	//create response to send a controller
	eventsResponse := EventResponse{true, eventData}

	return eventsResponse

}

// end : AddEvent

// GetAllEvent get all registered event
func GetAllEvent() EventsResponse {
	log.Println("Get all event service initialize")
	// call model to get a events
	eventList := model.GetEvents()

	// create event response struct
	eventsResponse := EventsResponse{true, eventList}

	return eventsResponse
}

// end : GetAllEvent

// GetEventDetail give event details
func GetEventDetail(id string) EventResponse {
	log.Println("Get event detail service initialize")
	// call model to get a events
	eventDetail := model.GetEvent(id)

	// create event response struct
	eventsResponse := EventResponse{true, eventDetail}

	// sending a response
	return eventsResponse
}

// end : GetEventDetail

// GetUnreadEventCount is ...
// get count of unread event
func GetUnreadEventCount() UnreadMessageCount {
	log.Println("Get unread event count service initialize")

	var count []uint64
	// call model to get a total event
	totalCount := model.GetUnreadEventCount()
	count = append(count, totalCount)
	response := UnreadMessageCount{true, count}
	return response
}

// end : GetUnreadEventCount

// GetUnreadEventTypeCount give count of unread event by event type
func GetUnreadEventTypeCount() EventsTypeResponse {
	log.Println("Get all event by count service initialize")
	// call model to get a total unread event
	countByEventType := model.GetUnreadEventTypeCount()

	// create event Count response struct
	eventsResponse := EventsTypeResponse{true, countByEventType}

	return eventsResponse

}

// end : GetUnreadEventTypeCount

// end : NonEventData

// Find a value in slice
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// end : Find

// SendNotification sending a notification count
func SendNotification() {
	notificationResponse := GetUnreadEventCount()
	serviceString, _ := json.Marshal(notificationResponse)
	socket.SendNotification(
		"success",
		string(serviceString),
	)

	log.Println("send notification on channel")

}

// end : SendNotification
