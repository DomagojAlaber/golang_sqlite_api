package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (event *Event) Save() {
	events = append(events, *event)
	// TODO: add to db
}

func GetAllEvents() ([]Event, error) {
	return events, nil
}
