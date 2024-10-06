package models

// Import the models package from event-service
import "backend/event-service/models"

type Itinerary struct {
	Description string
	Name        string
	Location    string
	Tags        []string
	Events      []models.Event
	Date        string
}

// Add methods for Itinerary here
