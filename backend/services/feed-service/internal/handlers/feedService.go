package handlers

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	db "github.com/jordyob03/TripTailor/backend/services/feed-service/internal/db"
)

// FeedService handles feed requests based on tags
func FeedService(database *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		//Extract and parse tags from query parameters
		tagsParameters := c.Query("tags")
		if tagsParameters == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Tags parameter is required"})
			return
		}

		tags := strings.Split(tagsParameters, ",")
		if len(tags) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No tags specified"})
			return
		}

		//Query the DB for itins with specified tags
		itineraries, err := db.QueryItinerariesByTags(database, tags)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		//Return the itineraries
		c.JSON(http.StatusOK, gin.H{
			"message":     "Itineraries retrieved successfully",
			"itineraries": itineraries,
		})
	}
}
