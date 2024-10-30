package handlers

import (
	"database/sql"
	"net/http"
	//"strconv"

	db "github.com/jordyob03/TripTailor/backend/services/search-service/internal/db" // Import from internal/db

	"github.com/gin-gonic/gin"
)

// SearchItineraries handles search requests with optional filters
func SearchItineraries(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Collect query parameters dynamically
		params := map[string]interface{}{}

		// Add parameters only if they are present in the request
		if username := c.Query("username"); username != "" {
			params["username"] = username
		}
		if country := c.Query("country"); country != "" {
			params["country"] = country
		}
		if city := c.Query("city"); city != "" {
			params["city"] = city
		}
		if tags := c.QueryArray("tags"); len(tags) > 0 {
			params["tags"] = tags
		}
		if languages := c.QueryArray("languages"); len(languages) > 0 {
			params["languages"] = languages
		}

		// If no parameters are provided, return an error
		if len(params) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No query parameters provided"})
			return
		}

		// Query the database using the dynamic parameters
		itineraries, err := db.QueryItineraries(dbConn, params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the results as JSON
		c.JSON(http.StatusOK, itineraries)
	}
}
