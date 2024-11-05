package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jordyob03/TripTailor/backend/services/search-service/internal/models" // Import the models package
	"github.com/lib/pq"
)

// QueryItineraries builds a dynamic SQL query based on provided parameters
func QueryItineraries(db *sql.DB, params map[string]interface{}) ([]models.Itinerary, error) {
	baseQuery := `
		SELECT itineraryid, name, city, country, languages, tags, events, postid, username, creationdate, lastupdate
		FROM itineraries
		WHERE 1=1
	`

	var conditions []string
	var args []interface{}
	argCounter := 1

	// Build the query conditions dynamically
	for key, value := range params {
		switch key {
		case "username", "country", "city":
			// Properly format conditions without extra quotes
			conditions = append(conditions, fmt.Sprintf("%s = $%d", key, argCounter))
			args = append(args, value)
			argCounter++
		case "tags", "languages":
			val, ok := value.([]string)
			if ok && len(val) > 0 {
				conditions = append(conditions, fmt.Sprintf("%s && $%d", key, argCounter))
				args = append(args, pq.Array(val))
				argCounter++
			}
		}
	}

	// Safely join the conditions with proper spacing
	if len(conditions) > 0 {
		baseQuery += " AND " + strings.Join(conditions, " AND ")
	}

	// Print the query and parameters for debugging
	fmt.Println("Final Query:", baseQuery)
	fmt.Println("Query Args:", args)

	// Prepare the query to catch any syntax issues
	stmt, err := db.Prepare(baseQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	// Execute the query
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	// Process the query results
	var itineraries []models.Itinerary
	for rows.Next() {
		var itinerary models.Itinerary
		if err := rows.Scan(
			&itinerary.ItineraryId, &itinerary.Name, &itinerary.City, &itinerary.Country,
			pq.Array(&itinerary.Languages), pq.Array(&itinerary.Tags), pq.Array(&itinerary.Events),
			&itinerary.PostId, &itinerary.Username, &itinerary.CreationDate, &itinerary.LastUpdate,
		); err != nil {
			return nil, err
		}
		itineraries = append(itineraries, itinerary)
	}

	return itineraries, nil
}
