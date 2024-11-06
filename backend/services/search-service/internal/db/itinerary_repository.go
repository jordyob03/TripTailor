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

	// Prepare SQL conditions based on provided params
	if v, ok := params["tags"].(string); ok && v != "" {
		tagArray = prepareTagArray(v)
		conditions = append(conditions, fmt.Sprintf("tags && %s", formatArrayForSQL(tagArray)))
	}
	if v, ok := params["languages"].(string); ok && v != "" {
		langArray = prepareTagArray(v)
		conditions = append(conditions, fmt.Sprintf("languages && %s", formatArrayForSQL(langArray)))
	}
	if v, ok := params["city"].(string); ok && v != "" {
		conditions = append(conditions, fmt.Sprintf("city = '%s'", v))
	}
	if v, ok := params["country"].(string); ok && v != "" {
		conditions = append(conditions, fmt.Sprintf("country = '%s'", v))
	}
	if v, ok := params["username"].(string); ok && v != "" {
		conditions = append(conditions, fmt.Sprintf("username = '%s'", v))
	}

	fmt.Printf("Conditions: %v\n", conditions)
	// Build the WHERE clause only if conditions exist
	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	// Construct the SQL query with scoring mechanism
	query := fmt.Sprintf(`
		SELECT itineraryid, name, city, country, languages, tags, events, postid, username, creationdate, lastupdate,
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

		if err := rows.Scan(
			&scored.Itinerary.ItineraryId, &scored.Itinerary.Name, &scored.Itinerary.City, &scored.Itinerary.Country,
			pq.Array(&scored.Itinerary.Languages), pq.Array(&scored.Itinerary.Tags), pq.Array(&scored.Itinerary.Events),
			&scored.Itinerary.PostId, &scored.Itinerary.Username, &scored.Itinerary.CreationDate, &scored.Itinerary.LastUpdate,
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
