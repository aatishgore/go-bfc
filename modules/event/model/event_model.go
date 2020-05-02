package model

import (
	"BFC/utilities"

	"github.com/jinzhu/gorm"
)

// Event is ...
// type Event struct {
// 	ID          uint   `gorm:"primary_key;AUTO_INCREMENT;column:id"`
// 	EventName   string `gorm:"not null;column:event_name;type:varchar(50)"`
// 	EventSource string `gorm:"not null;column:event_source;type:varchar(100)"`
// 	EventDetail string `gorm:"not null;column:event_detail;type:varchar(255)"`
// 	EventType   string `gorm:"not null;column:event_type;type:enum('Error','Warning')"`
// 	isRead      int    `gorm:"not null;column:is_read;type:tinyint(1)"`
// CreatedAt   time.Time
// UpdatedAt   time.Time
// }

// Event is ...
type Event struct {
	ID          int    `json:"id"`
	EventName   string `json:"event_name"`
	EventSource string `json:"event_source"`
	EventDetail string `json:"event_detail"`
	EventType   string `json:"event_type"`
}

// Count is ...
type Count struct {
	EventType string `json:"event_type"`
	Count     int    `json:"count"`
}

// Events is ...
type Events []Event

var db *gorm.DB

// AddEvent is ...
// Add event is create a event in data
func AddEvent(event Event) Event {
	// connect to database
	db = utilities.DbConnect(db)

	//close database connection at end of database operation
	defer db.Close()

	// create event in db
	db.Create(&event)
	return event
}

// end : AddEvent

// GetEvents is ...
// GetEvents give you all events
func GetEvents() Events {
	// connect to database
	db = utilities.DbConnect(db)

	//close database connection at end of database operation
	defer db.Close()

	// create an empty object of Events
	events := Events{}

	// find all events
	db.Find(&events)

	// return an event list
	return events

}

// end : GetEvents

// GetEvent is ...
// give detail mach with id
func GetEvent(id string) Event {
	// connect to database
	db = utilities.DbConnect(db)

	//close database connection at end of database operation
	defer db.Close()

	// create an empty object of Events
	event := Event{}

	// find all events
	db.Find(&event, id)

	// return an event list
	return event

}

// end : GetEvents

// GetUnreadEventCount is ...
// give detail mach with id
func GetUnreadEventCount() uint64 {
	// connect to database
	db = utilities.DbConnect(db)

	//close database connection at end of database operation
	defer db.Close()

	// get count of unread event
	var count uint64
	db.Table("events").Where("is_read is null AND event_type IN (?, ?)", "error", "warning").Count(&count)

	// return an event list
	return count

}

// end : GetUnreadEventCount

// GetUnreadEventTypeCount is ...
// GetUnreadEventTypeCount give event type wise count
func GetUnreadEventTypeCount() []Count {
	// connect to database
	db = utilities.DbConnect(db)

	//close database connection at end of database operation
	defer db.Close()

	counts := []Count{}

	db.Table("events").Select("event_type,count(id) as count").Where("is_read is null AND event_type IN (?, ?)", "error", "warning").Group("event_type").Scan(&counts)

	return counts

}
