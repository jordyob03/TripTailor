package models

import (
	"time"
)

type Event struct {
	EventId      int       `json:"eventId"`
	Name         string    `json:"name"`
	Price        int       `json:"price"`
	Location     string    `json:"location"`
	Description  string    `json:"description"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	ItineraryIds []string  `json:"itineraryIds"`
	PhotoLinks   []string  `json:"photoLinks"`
}
