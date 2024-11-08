package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jordyob03/TripTailor/backend/services/search-service/internal/models"
	"github.com/lib/pq" // PostgreSQL driver
)

// formatArrayForSQL formats a slice of strings for SQL array operations
func formatArrayForSQL(arr []string) string {
	quoted := make([]string, len(arr))
	for i, v := range arr {
		quoted[i] = fmt.Sprintf("'%s'", v)
	}
	return fmt.Sprintf("ARRAY[%s]::text[]", strings.Join(quoted, ","))
}

// prepareTagArray splits a comma-separated string into a slice
func prepareTagArray(tagString string) []string {
	return strings.Split(tagString, ",")
}

func QueryItineraries(db *sql.DB, params map[string]interface{}) ([]models.Itinerary, error) {
	var conditions []string
	var tagArray, langArray []string

	for key, value := range params {
		switch key {
		case "tags":
			if tagArray, ok := value.([]string); ok && len(tagArray) > 0 {
				tagArray := prepareTagArray(tagArray[0])
				conditions = append(conditions, fmt.Sprintf("tags && %s", formatArrayForSQL(tagArray)))
			}
		case "languages":
			if langArray, ok := value.([]string); ok && len(langArray) > 0 {
				langArray := prepareTagArray(langArray[0])
				conditions = append(conditions, fmt.Sprintf("languages && %s", formatArrayForSQL(langArray)))
			}
		case "city", "country", "username", "name", "title":
			if v, ok := value.(string); ok && v != "" {
				conditions = append(conditions, fmt.Sprintf("%s ILIKE '%%%s%%'", key, v)) // ILIKE for case-insensitive partial matches
			}
		case "price":
			if price, ok := value.(float64); ok && price > 0 {
				conditions = append(conditions, fmt.Sprintf("price <= %f", price)) // Example: search for itineraries <= the given price
			}
		}
	}

	fmt.Printf("Conditions: %v\n", conditions)
	// Build the WHERE clause only if conditions exist
	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	// Construct the SQL query with scoring mechanism
	query := fmt.Sprintf(`
		SELECT itineraryid, name, city, country, title, description, price, languages, tags, events, postid, username,
			array_length(array(select unnest(tags) intersect select unnest(%s)), 1) AS tag_match_count,
			array_length(array(select unnest(languages) intersect select unnest(%s)), 1) AS lang_match_count,
			array_length(array(select unnest(tags) intersect select unnest(%s)), 1) +
			array_length(array(select unnest(languages) intersect select unnest(%s)), 1) AS total_match_count
		FROM itineraries
		%s
		ORDER BY total_match_count DESC,
		lang_match_count DESC, 
        tag_match_count DESC;
	`, formatArrayForSQL(tagArray), formatArrayForSQL(langArray), formatArrayForSQL(tagArray), formatArrayForSQL(langArray), whereClause)

	// Print the final query and arguments for debugging
	fmt.Printf("Final Query: %s\n", query)
	//fmt.Printf("Query Args: %v\n", args)

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	// Process the rows into ScoredItinerary objects
	var scoredItineraries []models.ScoredItinerary
	for rows.Next() {
		var scored models.ScoredItinerary
		var tagMatchCount, languageMatchCount, totalMatchCount sql.NullInt64

		// Update the scanning process to account for all fields in Itinerary and the match counts
		if err := rows.Scan(
			&scored.Itinerary.ItineraryId, &scored.Itinerary.Name, &scored.Itinerary.City, &scored.Itinerary.Country,
			&scored.Itinerary.Title, &scored.Itinerary.Description, &scored.Itinerary.Price,
			pq.Array(&scored.Itinerary.Languages), pq.Array(&scored.Itinerary.Tags), pq.Array(&scored.Itinerary.Events),
			&scored.Itinerary.PostId, &scored.Itinerary.Username,
			&tagMatchCount, &languageMatchCount, &totalMatchCount,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		// Convert sql.NullInt64 to regular int if it's not null
		if tagMatchCount.Valid {
			scored.TagMatchCount = int(tagMatchCount.Int64)
		}
		if languageMatchCount.Valid {
			scored.LanguageMatchCount = int(languageMatchCount.Int64)
		}
		if totalMatchCount.Valid {
			scored.TotalMatchCount = int(totalMatchCount.Int64)
		}

		scoredItineraries = append(scoredItineraries, scored)
	}

	// Extract only the Itinerary part to return
	itineraries := make([]models.Itinerary, len(scoredItineraries))
	for i, scored := range scoredItineraries {
		itineraries[i] = scored.Itinerary
	}

	return itineraries, nil
}
