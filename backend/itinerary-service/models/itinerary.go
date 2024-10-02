package models

// Import the models package from event-service
import "backend/event-service/models"

// Itinerary struct definition
type Itinerary struct {
	ItineraryDescription string
	ItineraryName        string
	ItineraryLocation    string
	ItineraryTags        []string
	ItineraryEvents      []models.Event // Use Event struct from imported package
	ItineraryDate        string
}

// Add methods for Itinerary here
