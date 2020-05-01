package service

import (
	"BFC/modules/event/model"
	"BFC/utilities"
	"encoding/json"

	"fmt"
)

// AddEvent is ...
// AddEvent create an event and sent response to service.
func AddEvent() string {
	// define response variable
	var response string

	event := &model.Event{
		EventName:   "this is a event Name",
		EventSource: "this is a demo event source",
		EventDetail: "this is a event detail",
		EventType:   "Warning",
	}

	// call model to create a event
	eventResponse := model.AddEvent(event)

	// check event is created or not
	if eventResponse != 0 {
		fmt.Println("Record has Created")
		// convert response into json string and encodeing to response fro controller
		response = utilities.ResponseEncode(false, "Record was created")
	} else {
		fmt.Println("Record has not created")
		// encoding to response for controller
		response = utilities.ResponseEncode(false, "")
	}

	return response

}

// end : AddEvent

// GetAllEvent is ...
// this service all event
func GetAllEvent() string {
	// call model to get a events
	eventList := model.GetEvents()

	// decode into json to send a response
	eventJSONEncoded, _ := json.Marshal(eventList)

	// encode in json and return a response
	response := utilities.ResponseEncode(false, string(eventJSONEncoded))
	return response
}

// end : GetAllEvent
