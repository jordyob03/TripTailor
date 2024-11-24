package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jordyob03/TripTailor/backend/services/search-service/internal/db"
)

type GetSearchRequest struct {
	SearchValue string  `json:"searchValue" form:"searchValue" binding:"required"`
	Price       float64 `json:"price" form:"price" binding:"required"`
}

func SearchItineraries(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetSearchRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			fmt.Println("Error parsing JSON: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		itineraries, err := db.GetScoredItineraries(dbConn, req.SearchValue, req.Price)
		if err != nil {
			fmt.Printf("Query error: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve itineraries", "details": err.Error()})
			return
		}

		c.JSON(http.StatusOK, itineraries)
	}
}
