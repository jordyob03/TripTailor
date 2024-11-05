package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jordyob03/TripTailor/backend/services/search-service/internal/models"
	"github.com/lib/pq"
)

// Helper function to parse comma-separated parameters into a slice
func parseCommaSeparatedParam(param string) []string {
	if param == "" {
		return []string{}
	}
	return strings.Split(param, ",")
}

// Helper function to convert a slice of strings to []interface{} for pq.Array
func toInterfaceSlice(slice []string) []interface{} {
	interfaceSlice := make([]interface{}, len(slice))
	for i, v := range slice {
		interfaceSlice[i] = v
	}
	return interfaceSlice
}

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

	// Dynamically build query conditions
	for key, value := range params {
		switch key {
		case "username", "country", "city":
			if value != "" {
				conditions = append(conditions, fmt.Sprintf("%s = $%d", key, argCounter))
				args = append(args, value)
				argCounter++
			}
		case "tags":
			// Use the overlap operator (&&) for tags
			if tagArray, ok := value.([]string); ok && len(tagArray) > 0 {
				fmt.Printf("Parsed Tags Array: %v\n", tagArray) // Debug output
				conditions = append(conditions, fmt.Sprintf("tags && $%d", argCounter))
				args = append(args, pq.Array(toInterfaceSlice(tagArray)))
				argCounter++
			}
		case "languages":
			// Use the containment operator (@>) for languages
			if langArray, ok := value.([]string); ok && len(langArray) > 0 {
				fmt.Printf("Parsed Languages Array: %v\n", langArray) // Debug output
				conditions = append(conditions, fmt.Sprintf("languages @> $%d", argCounter))
				args = append(args, pq.Array(toInterfaceSlice(langArray)))
				argCounter++
			}
		}
	}

	// Combine base query with dynamically built conditions
	query := baseQuery + strings.Join(conditions, " AND ")

	// Print the final query and arguments for debugging
	fmt.Printf("Final Query: %s\n", query)
	fmt.Printf("Query Args: %+v\n", args)

	// Prepare and execute the query
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	// Process results
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
