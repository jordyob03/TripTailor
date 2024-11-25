package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jordyob03/TripTailor/backend/services/feed-service/internal/db"
)

// FeedService handles feed requests based on tags
func FeedService(database *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Extract the "tags" query parameter
		tagsParam := c.DefaultQuery("tags", "")
		if tagsParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Tags parameter is required"})
			return
		}

		// Parse the tags as a JSON array
		var tags []string
		if err := json.Unmarshal([]byte(tagsParam), &tags); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format for tags"})
			return
		}

		// Log the tags to verify
		fmt.Println("Got Tags:", tags)

		// Query the DB for itineraries with the specified tags
		itineraries, err := db.QueryItinerariesByTags(database, tags)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the itineraries
		c.JSON(http.StatusOK, gin.H{
			"message":     "Itineraries retrieved successfully",
			"itineraries": itineraries,
		})
	}
}
