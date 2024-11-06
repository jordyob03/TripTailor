package models

import (
	"time"
)

type ScoredItinerary struct {
	Itinerary
	TagMatchCount      int `json:"tagMatchCount"`
	LanguageMatchCount int `json:"languageMatchCount"`
	TotalMatchCount    int `json:"totalMatchCount"`
}

type Itinerary struct {
	ItineraryId  int       `json:"itineraryId"`
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
	//Cost         float32   `json:"Cost"`
}

// Add methods for Itinerary here
