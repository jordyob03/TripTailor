package handlers

import (
	"database/sql"
	"fmt"
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
			fmt.Println("Error: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve post"})
			return
		}

		fmt.Println("PostId: ", Post.PostId)

		// if Post.ItineraryId <= 0 {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid itineraryId 0"})
		// 	return
		// }

		if Post.ItineraryId <= 0 {
			Post.ItineraryId = 1
		}

		itinerary, err := db.GetItinerary(dbConn, Post.ItineraryId)
		if err != nil {
			fmt.Println("Error: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve itinerary"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Itinerary": itinerary})
	}
}
