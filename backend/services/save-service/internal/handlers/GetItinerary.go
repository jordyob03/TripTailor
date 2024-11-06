package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jordyob03/TripTailor/backend/services/save-service/internal/db"
)

type GetItineraryRequest struct {
	PostId int `json:"postId" form:"postId" binding:"required"`
}

func GetItinerary(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetItineraryRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		Post, err := db.GetPost(dbConn, req.PostId)
		if err != nil {
			fmt.Println("Error: ", err, "Post ID: ", req.PostId)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve post"})
			return
		}

		fmt.Println("Post: ", Post)
		fmt.Println("Itinerary ID: ", Post.ItineraryId)

		itinerary, err := db.GetItinerary(dbConn, Post.ItineraryId)
		if err != nil {
			fmt.Println("Error: ", err, "Itinerary ID: ", Post.ItineraryId)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve itinerary"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Itinerary": itinerary})
	}
}
