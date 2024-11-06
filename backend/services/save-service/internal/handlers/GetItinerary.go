package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jordyob03/TripTailor/backend/services/save-service/internal/db"
)

type GetItineraryRequest struct {
	postId int `json:"postId" form:"postId"`
}

func GetItinerary(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetItineraryRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// if req.postId <= 0 {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid postId 0"})
		// 	return
		// }

		if req.postId <= 0 {
			req.postId = 1
		}

		Post, err := db.GetPost(dbConn, req.postId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve post"})
			return
		}

		itinerary, err := db.GetItinerary(dbConn, Post.ItineraryId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve itinerary"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Itinerary": itinerary})
	}
}
