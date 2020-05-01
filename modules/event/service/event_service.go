package service

import (
	"BFC/modules/event/model"
	"fmt"
)

// EventResponse is ...
type EventResponse struct {
	Status bool        `json:"status"`
	Data   model.Event `json:"data"`
}

// EventsResponse is ...
type EventsResponse struct {
	Status bool         `json:"status"`
	Data   model.Events `json:"data"`
}

// EventsTypeResponse is ...
type EventsTypeResponse struct {
	Status bool          `json:"status"`
	Data   []model.Count `json:"data"`
}

// define response variable
var response string

// AddEvent is ...
// AddEvent create an event and sent response to service.
func AddEvent() EventResponse {
	event := model.Event{
		EventName:   "this is a event Name",
		EventSource: "this is a demo event source",
		EventDetail: "this is a event detail",
		EventType:   "Warning",
	}

	// call model to create a event
	eventData := model.AddEvent(event)

	//create response to send a controller
	eventsResponse := EventResponse{true, eventData}

	return eventsResponse

}

// end : AddEvent

// GetAllEvent is ...
// this service all event
func GetAllEvent() EventsResponse {

	// call model to get a events
	eventList := model.GetEvents()

	// create event response struct
	eventsResponse := EventsResponse{true, eventList}

	return eventsResponse
}

// end : GetAllEvent

// GetEventDetail is ...
// this service five event detail
func GetEventDetail(id string) EventResponse {

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
func GetUnreadEventCount() {

	// call model to get a total event
	totalCount := model.GetUnreadEventCount()

	fmt.Println("total Count", totalCount)
}

// end : GetUnreadEventCount

// GetUnreadEventTypeCount is ...
// get count of unread event
func GetUnreadEventTypeCount() EventsTypeResponse {

	// call model to get a total unread event
	countByEventType := model.GetUnreadEventTypeCount()

	// create event Count response struct
	eventsResponse := EventsTypeResponse{true, countByEventType}

	return eventsResponse

}

// end : GetUnreadEventTypeCount
