package db

import (
	"backend/search-service/internal/models"
	"database/sql"
	"github.com/lib/pq"
	"strings"
)

// QueryItinerariesByLocation queries the database for itineraries by country and city,
// including their associated events.
func QueryItinerariesByLocation(db *sql.DB, country, city string) ([]models.Itinerary, error) {
	// Query the itineraries based on location
	itineraryRows, err := db.Query(`
		SELECT id, description, name, location, date, tags 
		FROM itineraries
		WHERE country=$1 AND city=$2`, country, city)
	if err != nil {
		return nil, err
	}
	defer itineraryRows.Close()

	// Create a map to store itineraries by ID so we can associate events later
	itinerariesMap := make(map[int]*models.Itinerary)

	for itineraryRows.Next() {
		var itinerary models.Itinerary
		var itineraryID int
		var tags string

		// Scan itinerary fields
		err := itineraryRows.Scan(&itineraryID, &itinerary.Description, &itinerary.Name, &itinerary.Location, &itinerary.Date, &tags)
		if err != nil {
			return nil, err
		}

		// Convert tags (comma-separated) into a slice
		itinerary.Tags = splitTags(tags)

		// Store the itinerary in the map by ID
		itinerariesMap[itineraryID] = &itinerary
	}

	// If there are no itineraries, return an empty slice
	if len(itinerariesMap) == 0 {
		return []models.Itinerary{}, nil
	}

	// Query the events associated with the itineraries
	itineraryIDs := make([]int, 0, len(itinerariesMap))
	for id := range itinerariesMap {
		itineraryIDs = append(itineraryIDs, id)
	}

	// Query for events associated with the itineraries
	eventRows, err := db.Query(`
		SELECT id, itinerary_id, name, date, time, location, price, description, photos
		FROM events
		WHERE itinerary_id = ANY($1)`, pq.Array(itineraryIDs))
	if err != nil {
		return nil, err
	}
	defer eventRows.Close()

	// Process events and associate them with the correct itineraries
	for eventRows.Next() {
		var event models.Event
		var itineraryID int
		var photos string

		// Scan event fields
		err := eventRows.Scan(&event.Id, &itineraryID, &event.Name, &event.Date, &event.Time, &event.Location, &event.Price, &event.Description, &photos)
		if err != nil {
			return nil, err
		}

		// Convert photos (comma-separated) into a slice
		event.Photos = splitPhotos(photos)

		// Add the event to the correct itinerary
		itinerary := itinerariesMap[itineraryID]
		itinerary.Events = append(itinerary.Events, event)
	}

	// Collect the itineraries from the map into a slice
	itineraries := make([]models.Itinerary, 0, len(itinerariesMap))
	for _, itinerary := range itinerariesMap {
		itineraries = append(itineraries, *itinerary)
	}

	return itineraries, nil
}

// Helper function to split comma-separated tags into a slice of strings
func splitTags(tags string) []string {
	if tags == "" {
		return []string{}
	}
	return strings.Split(tags, ",")
}

// Helper function to split comma-separated photos into a slice of strings
func splitPhotos(photos string) []string {
	if photos == "" {
		return []string{}
	}
	return strings.Split(photos, ",")
}
