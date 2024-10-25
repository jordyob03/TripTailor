package db

import (
	"database/sql"
	"fmt"
	"github.com/jordyob03/TripTailor/backend/services/search-service/internal/models" // Import the models package
	"github.com/lib/pq"
	"strings"
)

// QueryItineraries builds a dynamic query based on the provided parameters
func QueryItineraries(db *sql.DB, params map[string]interface{}) ([]models.Itinerary, error) {
	baseQuery := `
		SELECT itineraryid, name, city, country, languages, tags, events, postid, username, creationdate, lastupdate, cost
		FROM itineraries
		WHERE 1=1
	`
	// Query parts to build dynamically
	var conditions []string
	var args []interface{}
	argCounter := 1

	// Dynamically build conditions based on provided parameters
	for key, value := range params {
		switch key {
		case "country", "city", "username":
			if value != "" {
				conditions = append(conditions, fmt.Sprintf("%s = $%d", key, argCounter))
				args = append(args, value)
				argCounter++
			}
		case "max_cost":
			conditions = append(conditions, fmt.Sprintf("cost <= $%d", argCounter))
			args = append(args, value)
			argCounter++
		case "tags", "languages":
			if len(value.([]string)) > 0 {
				conditions = append(conditions, fmt.Sprintf("%s && $%d", key, argCounter)) // PostgreSQL array intersection
				args = append(args, pq.Array(value))
				argCounter++
			}
		}
	}

	// Join conditions to build the final query
	query := baseQuery + strings.Join(conditions, " AND ")

	// Execute the query
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Process the results
	itineraries := []models.Itinerary{}
	for rows.Next() {
		var itinerary models.Itinerary
		if err := rows.Scan(
			&itinerary.ItineraryId, &itinerary.Name, &itinerary.City, &itinerary.Country,
			pq.Array(&itinerary.Languages), pq.Array(&itinerary.Tags), pq.Array(&itinerary.Events),
			&itinerary.PostId, &itinerary.Username, &itinerary.CreationDate, &itinerary.LastUpdate,
			&itinerary.Cost,
		); err != nil {
			return nil, err
		}
		itineraries = append(itineraries, itinerary)
	}
	return itineraries, nil
}
