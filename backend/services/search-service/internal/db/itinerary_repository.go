package db

import (
	"database/sql"
	"strings"

	"github.com/jordyob03/TripTailor/backend/services/search-service/internal/models" // Import the models package
)

// QueryItinerariesByLocation queries the database for itineraries based on country and city
func QueryItinerariesByLocation(db *sql.DB, country, city string) ([]models.Itinerary, error) {
	rows, err := db.Query(`
		SELECT itinerary_id, name, city, country, languages, tags, events, post_id, username, creation_date, last_update
		FROM itineraries
		WHERE country = $1 AND city = $2`, country, city)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	itineraries := []models.Itinerary{}
	for rows.Next() {
		var itinerary models.Itinerary
		var languages, tags, events string

		// Scan each row into the itinerary struct
		if err := rows.Scan(
			&itinerary.ItineraryId, &itinerary.Name, &itinerary.City, &itinerary.Country,
			&languages, &tags, &events, &itinerary.PostId, &itinerary.Username,
			&itinerary.CreationDate, &itinerary.LastUpdate,
		); err != nil {
			return nil, err
		}

		// Convert comma-separated strings to slices
		itinerary.Languages = splitTags(languages)
		itinerary.Tags = splitTags(tags)
		itinerary.Events = splitTags(events)

		itineraries = append(itineraries, itinerary)
	}
	return itineraries, nil
}

// Helper function to split comma-separated values into a slice of strings
func splitTags(input string) []string {
	if input == "" {
		return []string{}
	}
	return strings.Split(input, ",")
}
