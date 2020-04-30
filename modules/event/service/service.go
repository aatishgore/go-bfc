package service

import (
	"fmt"
)

type Event struct {
	Id
	EventName
	EventSource
	EventEventDetail
	EventType
}

type Events struct {
	Event []Event `json:"event"`
}

func GetEventList() {
	fmt.Println("event service to get a list")
}
