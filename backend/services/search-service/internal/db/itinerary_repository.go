package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jordyob03/TripTailor/backend/services/search-service/internal/models"
	"github.com/lib/pq" // PostgreSQL driver
)

func BuildQuery(searchString string, price float64) (string, []interface{}) {
	// Tokenize the search string into individual words
	tokens := strings.Fields(searchString)
	if len(tokens) == 0 && price <= 0 {
		return "", nil // No search terms or price provided
	}

	innerQuery := "SELECT *, (" // Start inner query with SELECT and begin total_score calculation
	scoringConditions := []string{}
	filterConditions := []string{} // Conditions for calculating total_score
	args := []interface{}{}        // Holds query parameter values
	argIndex := 1                  // Counter for argument placeholders

	// Add scoring for exact matches to the title
	scoringConditions = append(scoringConditions, fmt.Sprintf("(CASE WHEN title ILIKE $%d THEN 10 ELSE 0 END)", argIndex))
	args = append(args, "%"+searchString+"%")
	argIndex++

	// Add scoring for each token against multiple fields
	for _, token := range tokens {
		// Match token against the title
		scoringConditions = append(scoringConditions, fmt.Sprintf("(CASE WHEN title ILIKE '%%' || $%d || '%%' THEN 2 ELSE 0 END)", argIndex))
		args = append(args, token)
		argIndex++

		// Match token against city
		scoringConditions = append(scoringConditions, fmt.Sprintf("(CASE WHEN city ILIKE '%%' || $%d || '%%' THEN 5 ELSE 0 END)", argIndex))
		args = append(args, token)
		argIndex++

		// Match token against country
		scoringConditions = append(scoringConditions, fmt.Sprintf("(CASE WHEN country ILIKE '%%' || $%d || '%%' THEN 5 ELSE 0 END)", argIndex))
		args = append(args, token)
		argIndex++

		// Match token against username
		scoringConditions = append(scoringConditions, fmt.Sprintf("(CASE WHEN username ILIKE '%%' || $%d || '%%' THEN 1 ELSE 0 END)", argIndex))
		args = append(args, token)
		argIndex++

		// Match token against tags
		scoringConditions = append(scoringConditions, fmt.Sprintf("(SELECT COUNT(*) FROM unnest(tags) AS tag WHERE tag ILIKE '%%' || $%d || '%%')", argIndex))
		args = append(args, token)
		argIndex++

		// Match token against languages
		scoringConditions = append(scoringConditions, fmt.Sprintf("(SELECT COUNT(*) FROM unnest(languages) AS lang WHERE lang ILIKE '%%' || $%d || '%%')", argIndex))
		args = append(args, token)
		argIndex++
	}

	// Combine scoring conditions
	if len(scoringConditions) > 0 {
		innerQuery += strings.Join(scoringConditions, " + ")
	} else {
		innerQuery += "0"
	}
	innerQuery += ") AS total_score FROM itineraries"

	if price > 0 {
		filterConditions = append(filterConditions, fmt.Sprintf("price <= $%d", argIndex))
		args = append(args, price)
		argIndex++
	}

	// Add filter conditions to the WHERE clause
	if len(filterConditions) > 0 {
		innerQuery += " WHERE " + strings.Join(filterConditions, " AND ")
	}

	// Wrap the inner query in an outer query to apply the total_score filter
	outerQuery := `
        SELECT * 
        FROM (%s) AS scored_itineraries
        WHERE total_score >= 2
        ORDER BY total_score DESC`
	finalQuery := fmt.Sprintf(outerQuery, innerQuery)

	return finalQuery, args
}

// GetScoredItinerariesFromSearchString dynamically builds a query to filter and score itineraries from a single search string.
func GetScoredItineraries(db *sql.DB, searchString string, price float64) ([]models.ScoredItinerary, error) {
	fmt.Printf("Parsed search string: %s\n", searchString)
	query, args := BuildQuery(searchString, price)
	if query == "" {
		fmt.Println("No search terms provided.")
		return nil, nil
	}

	fmt.Printf("Executing query: %s with args: %+v\n", query, args)

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
