package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jordyob03/TripTailor/backend/services/feed-service/internal/models"
	"github.com/lib/pq"
)

// QueryItinerariesByTags fetches itinerareis that contain any of the inputted tags
func QueryItinerariesByTags(database *sql.DB, tags []string) ([]models.Itinerary, error) {
	query := `
	SELECT itineraryId, name, city, country, title, description, price, languages, tags, events, postId, username
	FROM itineraries
	WHERE tags && $1;`

	var itineraries []models.Itinerary
	rows, err := database.Query(query, pq.Array(tags))
	if err != nil {
		log.Printf("Error querying itineraries by tags: %v\n", err)
		return nil, fmt.Errorf("failed to query itineraries by tags: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var itinerary models.Itinerary
		if err := rows.Scan(
			&itinerary.ItineraryId,
			&itinerary.Name,
			&itinerary.City,
			&itinerary.Country,
			&itinerary.Title,
			&itinerary.Description,
			&itinerary.Price,
			pq.Array(&itinerary.Languages),
			pq.Array(&itinerary.Tags),
			pq.Array(&itinerary.Events),
			&itinerary.PostId,
			&itinerary.Username,
		); err != nil {
			log.Printf("Error scanning itinerar: %v\n", err)
			return nil, err
		}
		itineraries = append(itineraries, itinerary)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over itineraries: %v\n", err)
		return nil, err
	}

	return itineraries, nil
}
