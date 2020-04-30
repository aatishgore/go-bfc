package model

import "fmt"

type EventModel struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	EventName   string `gorm:"not null;column:event_name;size:50"`
	EventSource string `gorm:"not null;column:event_source;size:100"`
	EventDetail string `gorm:"not null;column:event_detail;size:255"`
	EventType   string `gorm:"not null;column:event_type;size:50"`
}

func GetListEvent(limit, offset int) {
	fmt.Println("inside a model list")
}
