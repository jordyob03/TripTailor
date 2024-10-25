package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	db "github.com/jordyob03/TripTailor/backend/services/search-service/internal/db" // Import from internal/db

	"github.com/gin-gonic/gin"
)

// SearchItineraries handles search requests with optional filters
func SearchItineraries(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract query parameters
		params := map[string]interface{}{
			"country":   c.Query("country"),
			"city":      c.Query("city"),
			"tags":      c.QueryArray("tags"),
			"languages": c.QueryArray("languages"),
			"username":  c.Query("username"),
		}

		// Handle max cost (convert to integer if provided)
		if maxCost := c.Query("max_cost"); maxCost != "" {
			if cost, err := strconv.Atoi(maxCost); err == nil {
				params["max_cost"] = cost
			}
		}

		// Call the dynamic query function
		itineraries, err := db.QueryItineraries(dbConn, params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the itineraries as JSON
		c.JSON(http.StatusOK, itineraries)
	}
}
