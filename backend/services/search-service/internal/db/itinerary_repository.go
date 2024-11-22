package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jordyob03/TripTailor/backend/services/search-service/internal/models"
	"github.com/lib/pq" // PostgreSQL driver
)

func BuildQuery(params map[string]interface{}) (string, []interface{}) {
	innerQuery := "SELECT *, ("     // Start inner query with SELECT and begin total_score calculation
	scoringConditions := []string{} // Conditions for calculating total_score
	filterConditions := []string{}  // Conditions for WHERE clause (filters)
	args := []interface{}{}         // Holds query parameter values
	argIndex := 1                   // Counter for argument placeholders
	paramCount := 0                 // Count of specified parameters

	// City as a filter
	if city, ok := params["city"].(string); ok && city != "" {
		filterConditions = append(filterConditions, fmt.Sprintf("city ILIKE '%%' || $%d || '%%'", argIndex))
		scoringConditions = append(scoringConditions, fmt.Sprintf("(CASE WHEN city ILIKE '%%' || $%d || '%%' THEN 1 ELSE 0 END)", argIndex))
		args = append(args, city)
		argIndex++
	}

	if country, ok := params["country"].(string); ok && country != "" {
		filterConditions = append(filterConditions, fmt.Sprintf("country ILIKE '%%' || $%d || '%%'", argIndex))
		scoringConditions = append(scoringConditions, fmt.Sprintf("(CASE WHEN country ILIKE '%%' || $%d || '%%' THEN 1 ELSE 0 END)", argIndex))
		args = append(args, country)
		argIndex++
	}

	// Add scoring for languages
	if languages, ok := params["languages"].([]string); ok && len(languages) > 0 {
		scoringConditions = append(scoringConditions, fmt.Sprintf("(SELECT COUNT(*) FROM unnest($%d::text[]) AS lang WHERE lang = ANY(languages))", argIndex))
		args = append(args, pq.Array(languages))
		paramCount++
		argIndex++
	}

	// Add scoring for tags
	if tags, ok := params["tags"].([]string); ok && len(tags) > 0 {
		scoringConditions = append(scoringConditions, fmt.Sprintf("(SELECT COUNT(*) FROM unnest($%d::text[]) AS tag WHERE tag = ANY(tags))", argIndex))
		args = append(args, pq.Array(tags))
		paramCount++
		argIndex++
	}

	// Add scoring for price
	if price, ok := params["price"].(float64); ok && price > 0 {
		scoringConditions = append(scoringConditions, fmt.Sprintf("(CASE WHEN price <= $%d THEN 1 ELSE 0 END)", argIndex))
		args = append(args, price)
		paramCount++
		argIndex++
	}

	// Add scoring for title
	if title, ok := params["title"].(string); ok && title != "" {
		scoringConditions = append(scoringConditions, fmt.Sprintf("(CASE WHEN $%d::text IS NOT NULL AND title ILIKE '%%' || $%d || '%%' THEN 1 ELSE 0 END)", argIndex, argIndex))
		args = append(args, title)
		paramCount++
		argIndex++
	}

	// Add scoring for username
	if username, ok := params["username"].(string); ok && username != "" {
		scoringConditions = append(scoringConditions, fmt.Sprintf("(CASE WHEN $%d::text IS NOT NULL AND username ILIKE '%%' || $%d || '%%' THEN 1 ELSE 0 END)", argIndex, argIndex))
		args = append(args, username)
		paramCount++
		argIndex++
	}

	// Default score to 0 if no scoring conditions exist
	if len(scoringConditions) > 0 {
		innerQuery += strings.Join(scoringConditions, " + ")
	} else {
		innerQuery += "0"
	}
	innerQuery += ") AS total_score FROM itineraries"

	// Add filter conditions to the WHERE clause
	if len(filterConditions) > 0 {
		innerQuery += " WHERE " + strings.Join(filterConditions, " AND ")
	}

	// Determine the total_score threshold based on the number of parameters
	threshold := 1
	if paramCount >= 3 {
		threshold = 2
	}

	// Wrap the inner query in an outer query to apply the total_score filter
	outerQuery := fmt.Sprintf(`
        SELECT * 
        FROM (%s) AS scored_itineraries
        WHERE total_score >= %d
        ORDER BY total_score DESC`, innerQuery, threshold)

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
