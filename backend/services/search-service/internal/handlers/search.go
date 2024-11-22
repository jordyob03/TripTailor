package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jordyob03/TripTailor/backend/services/search-service/internal/db"
)

// SearchItineraries handles the search request from a single search bar
func SearchItineraries(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		searchString := c.Query("q")
		if searchString == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Search string cannot be empty"})
			return
		}

		priceParam := c.Query("price")
		var maxPrice float64
		if priceParam != "" {
			maxPrice, _ = strconv.ParseFloat(priceParam, 64)
		}

		itineraries, err := db.GetScoredItineraries(dbConn, searchString, maxPrice)
		if err != nil {
			fmt.Printf("Query error: %v\n", err) // Log the actual error
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve itineraries", "details": err.Error()})
			return
		}

		c.JSON(http.StatusOK, itineraries)
	}
}
