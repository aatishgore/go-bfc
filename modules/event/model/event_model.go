package model

import (
	"BFC/utilities"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// Event is ...
type Event struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	EventName   string `gorm:"not null;column:event_name;type:varchar(50)"`
	EventSource string `gorm:"not null;column:event_source;type:varchar(100)"`
	EventDetail string `gorm:"not null;column:event_detail;type:varchar(255)"`
	EventType   string `gorm:"not null;column:event_type;type:enum('Error','Warning')"`
	isRead      int    `gorm:"not null;column:event_type;type:tinyint(1)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Events is ...
type Events struct {
	Event []Event `json:"event"`
}

var db *gorm.DB

// AddEvent is ...
// Add event is create a event in data
func AddEvent(event *Event) uint {
	// connect to database
	db = utilities.DbConnect(db)

	//close database connection at end of database operation
	defer db.Close()

	// create event in db
	db.Create(&event)
	return 1
}

// end : AddEvent

// GetEvents is ...
func GetEvents() Events {
	// connect to database
	db = utilities.DbConnect(db)

	//close database connection at end of database operation
	defer db.Close()

	events := Events{}

	db.Find(&events)

	fmt.Println(events)

	return events

}

// end : GetEvents

// func createEventTable(db *gorm.DB) {
// 	db.AutoMigrate(&Event{})
// }
