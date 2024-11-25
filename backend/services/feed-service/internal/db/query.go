package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/jordyob03/TripTailor/backend/services/feed-service/internal/models"
	"github.com/lib/pq"
)

// QueryItinerariesByTags fetches itineraries that contain all of the inputted tags
func QueryItinerariesByTags(database *sql.DB, tags []string) ([]models.Itinerary, error) {
	// Build the condition for the query
	tagsConditions := []string{}
	args := []interface{}{}
	for i, tag := range tags {
		tagsConditions = append(tagsConditions, fmt.Sprintf("tags @> $%d", i+1))
		args = append(args, pq.Array([]string{tag}))
	}

	// Join all conditions with AND
	tagsCondition := strings.Join(tagsConditions, " AND ")

	// Construct the query
	query := fmt.Sprintf(`
		SELECT itineraryId, city, country, title, description, price, languages, tags, events, postId, username
		FROM itineraries
		WHERE %s;
	`, tagsCondition)

	// Execute the query
	var itineraries []models.Itinerary
	rows, err := database.Query(query, args...)
	if err != nil {
		log.Printf("Error querying itineraries by tags: %v\n", err)
		return nil, fmt.Errorf("failed to query itineraries by tags: %w", err)
	}
	defer rows.Close()

	// Scan the results
	for rows.Next() {
		var itinerary models.Itinerary
		if err := rows.Scan(
			&itinerary.ItineraryId,
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
			log.Printf("Error scanning itinerary: %v\n", err)
			return nil, err
		}
		itineraries = append(itineraries, itinerary)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over itineraries: %v\n", err)
		return nil, err
	}

	fmt.Println(itineraries)

	return itineraries, nil
}
