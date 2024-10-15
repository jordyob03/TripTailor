package models

type Itinerary struct {
	Description string
	Name        string
	Location    string
	Tags        []string
	Events      []Event
	Date        string
}

// Add methods for Itinerary here
