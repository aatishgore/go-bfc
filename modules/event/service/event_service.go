package service

import (
	"BFC/modules/event/model"
)

// EventResponse is ...
type EventResponse struct {
	Status bool        `json:"status"`
	Data   model.Event `json:"data"`
}

//EventNoDataResponse is set response for event no data found
type EventNoDataResponse struct {
	Status bool   `json:"status"`
	Data   string `json:"data"`
}

// EventsResponse is set response for event list
type EventsResponse struct {
	Status bool         `json:"status"`
	Data   model.Events `json:"data"`
}

// EventsTypeResponse is set response for event by type
type EventsTypeResponse struct {
	Status bool          `json:"status"`
	Data   []model.Count `json:"data"`
}

// EventsTypeResponseTt is ...
type UnreadMessageCount struct {
	Status bool     `json:"status"`
	Data   []uint64 `json:"data"`
}

// response variable
var response string

// AddEvent create an event and sent response to service.
func AddEvent(Event model.Event) EventResponse {

	// call model to create a event
	eventData := model.AddEvent(Event)

	//create response to send a controller
	eventsResponse := EventResponse{true, eventData}

	return eventsResponse

}

// end : AddEvent

// GetAllEvent get all registered event
func GetAllEvent() EventsResponse {

	// call model to get a events
	eventList := model.GetEvents()

	// create event response struct
	eventsResponse := EventsResponse{true, eventList}

	return eventsResponse
}

// end : GetAllEvent

// GetEventDetail give event details
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
func GetUnreadEventCount() UnreadMessageCount {
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

	// call model to get a total unread event
	countByEventType := model.GetUnreadEventTypeCount()

	// create event Count response struct
	eventsResponse := EventsTypeResponse{true, countByEventType}

	return eventsResponse

}

// end : GetUnreadEventTypeCount

// NonEventData set response if data was not set
func NonEventData() EventNoDataResponse {
	data := "No Event found"
	return EventNoDataResponse{true, data}
}

// end : NonEventData
