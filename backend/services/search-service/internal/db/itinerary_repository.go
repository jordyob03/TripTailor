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

func BuildQuery(params map[string]interface{}) (string, []interface{}) {
	innerQuery := "SELECT *, ("     // Start inner query with SELECT and begin total_score calculation
	scoringConditions := []string{} // Conditions for calculating total_score
	args := []interface{}{}         // Holds query parameter values
	argIndex := 1                   // Counter for argument placeholders

	// Add scoring condition for name
	if name, ok := params["name"].(string); ok && name != "" {
		scoringConditions = append(scoringConditions, fmt.Sprintf("(CASE WHEN $%d::text IS NOT NULL AND name ILIKE '%%' || $%d || '%%' THEN 1 ELSE 0 END)", argIndex, argIndex))
		args = append(args, name)
		argIndex++
	}

	// Add scoring condition for city
	if city, ok := params["city"].(string); ok && city != "" {
		scoringConditions = append(scoringConditions, fmt.Sprintf("(CASE WHEN $%d::text IS NOT NULL AND city ILIKE '%%' || $%d || '%%' THEN 1 ELSE 0 END)", argIndex, argIndex))
		args = append(args, city)
		argIndex++
	}

	// Add scoring condition for country
	if country, ok := params["country"].(string); ok && country != "" {
		scoringConditions = append(scoringConditions, fmt.Sprintf("(CASE WHEN $%d::text IS NOT NULL AND country ILIKE '%%' || $%d || '%%' THEN 1 ELSE 0 END)", argIndex, argIndex))
		args = append(args, country)
		argIndex++
	}

	// Add scoring for languages
	if languages, ok := params["languages"].([]string); ok && len(languages) > 0 {
		scoringConditions = append(scoringConditions, fmt.Sprintf("(SELECT COUNT(*) FROM unnest($%d::text[]) AS lang WHERE lang = ANY(languages))", argIndex))
		args = append(args, pq.Array(languages))
		argIndex++
	}

	// Add scoring for tags
	if tags, ok := params["tags"].([]string); ok && len(tags) > 0 {
		scoringConditions = append(scoringConditions, fmt.Sprintf("(SELECT COUNT(*) FROM unnest($%d::text[]) AS tag WHERE tag = ANY(tags))", argIndex))
		args = append(args, pq.Array(tags))
		argIndex++
	}

	// Add scoring for price
	if price, ok := params["price"].(float64); ok && price > 0 {
		scoringConditions = append(scoringConditions, fmt.Sprintf("(CASE WHEN price <= $%d THEN 1 ELSE 0 END)", argIndex))
		args = append(args, price)
		argIndex++
	}

	// Add scoring for title
	if title, ok := params["title"].(string); ok && title != "" {
		scoringConditions = append(scoringConditions, fmt.Sprintf("(CASE WHEN $%d::text IS NOT NULL AND title ILIKE '%%' || $%d || '%%' THEN 1 ELSE 0 END)", argIndex, argIndex))
		args = append(args, title)
		argIndex++
	}

	// Add scoring for username
	if username, ok := params["username"].(string); ok && username != "" {
		scoringConditions = append(scoringConditions, fmt.Sprintf("(CASE WHEN $%d::text IS NOT NULL AND username ILIKE '%%' || $%d || '%%' THEN 1 ELSE 0 END)", argIndex, argIndex))
		args = append(args, username)
		argIndex++
	}

	// Default score to 0 if no scoring conditions exist
	if len(scoringConditions) > 0 {
		innerQuery += strings.Join(scoringConditions, " + ")
	} else {
		innerQuery += "0"
	}
	innerQuery += ") AS total_score FROM itineraries"

	// Wrap the inner query in an outer query to apply the total_score filter
	outerQuery := fmt.Sprintf(`
        SELECT * 
        FROM (%s) AS scored_itineraries
        WHERE total_score > 0
        ORDER BY total_score DESC`, innerQuery)

	return outerQuery, args
}

// GetScoredItineraries dynamically builds a query to filter and score itineraries
func GetScoredItineraries(db *sql.DB, params map[string]interface{}) ([]models.ScoredItinerary, error) {
	fmt.Printf("Parsed parameters: %+v\n", params)
	query, args := BuildQuery(params)
	fmt.Printf("Executing query: %s with args: %+v\n", query, args)

	fmt.Printf("Executing query: %s with params: %+v\n", query, params)

	// Execute the query with parameters
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Printf("Query execution error: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	// Parse results
	itineraries := []models.ScoredItinerary{}
	for rows.Next() {
		var scored models.ScoredItinerary
		err := rows.Scan(
			&scored.Itinerary.ItineraryId,
			&scored.Itinerary.Name,
			&scored.Itinerary.City,
			&scored.Itinerary.Country,
			&scored.Itinerary.Title,
			&scored.Itinerary.Description,
			&scored.Itinerary.Price,
			pq.Array(&scored.Itinerary.Languages),
			pq.Array(&scored.Itinerary.Tags),
			pq.Array(&scored.Itinerary.Events),
			&scored.Itinerary.PostId,
			&scored.Itinerary.Username,
			&scored.TotalMatchCount, // Map the calculated score to TotalMatchCount
		)
		if err != nil {
			fmt.Printf("Row scan error: %v\n", err)
			return nil, err
		}
		itineraries = append(itineraries, scored)
	}

	return itineraries, nil
}
