package models

import (
	"time"
)

type Itinerary struct {
	ItineraryId  int       `json:"itineraryId" db:"itineraryid"`
	Name         string    `json:"name"`
	City         string    `json:"city"`
	Country      string    `json:"country"`
	Languages    []string  `json:"languages"`
	Tags         []string  `json:"tags"`
	Events       []string  `json:"events"`
	PostId       int       `json:"postId"`
	Username     string    `json:"username"`
	CreationDate time.Time `json:"creationDate"`
	LastUpdate   time.Time `json:"lastUpdate"`
}

// Add methods for Itinerary here
