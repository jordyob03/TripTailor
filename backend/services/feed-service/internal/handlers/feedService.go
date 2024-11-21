package handlers

import (
	"database/sql"
	"fmt"
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

		fmt.Print("Got Tags", tagsParameters)

		user := c.Query("username")
		if user == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User is required"})
			return
		}

		fmt.Print("Got User")

		tags := strings.Split(tagsParameters, ",")
		if len(tags) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No tags specified"})
			return
		}

		fmt.Print("Got Tags 2")

		//Query the DB for itins with specified tags
		itineraries, err := db.QueryItinerariesByTags(database, tags, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		fmt.Print("Passed Quert")

		//Return the itineraries
		c.JSON(http.StatusOK, gin.H{
			"message":     "Itineraries retrieved successfully",
			"itineraries": itineraries,
		})
	}
}
