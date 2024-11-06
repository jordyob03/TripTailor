package db

import (
	"database/sql"
	"strings"

	"github.com/jordyob03/TripTailor/backend/services/search-service/internal/models" // Import the models package
	"github.com/lib/pq"
)

type Itinerary struct {
	ItineraryId int      `json:"itineraryId"`
	Name        string   `json:"name"`
	City        string   `json:"city"`
	Country     string   `json:"country"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Languages   []string `json:"languages"`
	Tags        []string `json:"tags"`
	Events      []string `json:"events"`
	PostId      int      `json:"postId"`
	Username    string   `json:"username"`
}

// QueryItinerariesByLocation queries the database for itineraries based on country and city
func QueryItinerariesByLocation(db *sql.DB, country, city string) ([]models.Itinerary, error) {
	rows, err := db.Query(`
		SELECT itineraryid, name, city, country, title, description, price, languages, tags, events, postid, username
		FROM itineraries
		WHERE country = $1 AND city = $2`, country, city)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	itineraries := []models.Itinerary{}
	for rows.Next() {
		var itinerary models.Itinerary
		var languages, tags, events []string

		// Scan each row into the itinerary struct
		if err := rows.Scan(
			&itinerary.ItineraryId, &itinerary.Name, &itinerary.City, &itinerary.Country,
			&itinerary.Title, &itinerary.Description, &itinerary.Price,
			pq.Array(&languages), pq.Array(&tags), pq.Array(&events), &itinerary.PostId, &itinerary.Username,
		); err != nil {
			return nil, err
		}

		// Convert comma-separated strings to slices
		itinerary.Languages = splitTags(strings.Join(languages, ",")) // Convert []string to string
		itinerary.Tags = splitTags(strings.Join(tags, ","))           // Convert []string to string
		itinerary.Events = splitTags(strings.Join(events, ","))       // Convert []string to string

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
