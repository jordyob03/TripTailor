package handlers

import (
	"database/sql"
	"net/http"

	db "github.com/jordyob03/TripTailor/backend/services/search-service/internal/db" // Import from internal/db

	"github.com/gin-gonic/gin"
)

// SearchItineraries handles the search request
func SearchItineraries(database *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract query parameters
		country := c.Query("country")
		city := c.Query("city")

		// Validate query parameters
		if country == "" || city == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing query parameters"})
			return
		}

		// Query the database for itineraries using the db package
		itineraries, err := db.QueryItinerariesByLocation(database, country, city)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the itineraries as JSON
		c.JSON(http.StatusOK, itineraries)
	}
}
