package handlers

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/jordyob03/TripTailor/backend/services/search-service/internal/db"

	"github.com/gin-gonic/gin"
)

// SearchItineraries handles search requests with optional filters for tags and languages
func SearchItineraries(database *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse query parameters
		country := c.Query("country")
		city := c.Query("city")
		username := c.Query("username")

		// Parse comma-separated tags and languages into slices
		tags := parseCommaSeparatedParam(c.Query("tags"))
		languages := parseCommaSeparatedParam(c.Query("languages"))

		// Build params map with parsed values to pass to QueryItineraries
		params := map[string]interface{}{
			"country":   country,
			"city":      city,
			"username":  username,
			"tags":      tags,
			"languages": languages,
		}

		// Call the database function
		itineraries, err := db.QueryItineraries(database, params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the search results as JSON
		c.JSON(http.StatusOK, itineraries)
	}
}

// Helper function to parse a comma-separated string into a slice of strings
func parseCommaSeparatedParam(param string) []string {
	if param == "" {
		return nil
	}
	return strings.Split(param, ",")
}
