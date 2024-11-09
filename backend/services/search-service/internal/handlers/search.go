package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/jordyob03/TripTailor/backend/services/search-service/internal/db"

	"github.com/gin-gonic/gin"
)

// SearchItineraries handles the search request
func SearchItineraries(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := map[string]interface{}{
			"name":      c.Query("name"),
			"city":      c.Query("city"),
			"country":   c.Query("country"),
			"title":     c.Query("title"),
			"price":     parseFloatParam(c.Query("price")),
			"languages": parseArrayParam(c.Query("languages")),
			"tags":      parseArrayParam(c.Query("tags")),
			"username":  c.Query("username"),
		}

		itineraries, err := db.GetScoredItineraries(dbConn, params)
		if err != nil {
			fmt.Printf("Query error: %v\n", err) // Log the actual error
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve itineraries", "details": err.Error()})
			return
		}

		c.JSON(http.StatusOK, itineraries)
	}
}

func parseArrayParam(param string) []string {
	if param == "" {
		return nil // Pass nil for NULL parameters
	}
	return strings.Split(param, ",")
}

func parseFloatParam(param string) float64 {
	if param == "" {
		return 0
	}
	value, _ := strconv.ParseFloat(param, 64)
	return value
}
