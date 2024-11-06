package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/jordyob03/TripTailor/backend/services/search-service/internal/db"

	"github.com/gin-gonic/gin"
)

// SearchItineraries handles the search request
func SearchItineraries(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse query parameters into a map
		params := map[string]interface{}{
			"tags":      c.Query("tags"),
			"languages": c.Query("languages"),
			"country":   c.Query("country"),
			"city":      c.Query("city"),
			"username":  c.Query("username"),
		}

		// Debugging: Print parsed params
		fmt.Printf("Parsed Params: %v\n", params)

		// Query the database
		itineraries, err := db.QueryItineraries(dbConn, params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the itineraries as JSON
		c.JSON(http.StatusOK, itineraries)
	}
}
