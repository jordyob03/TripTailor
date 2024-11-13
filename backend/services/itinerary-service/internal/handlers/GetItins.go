package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	models "github.com/jordyob03/TripTailor/backend/services/itinerary-service/internal/models"

	"github.com/gin-gonic/gin"
)

type GetItinsRequest struct {
	Username string `json:"username"`
}

func GetItins(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetItinsRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := models.GetUser(dbConn, req.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
			return
		}

		var itineraries []models.Itinerary

		for _, post := range user.Posts {
			postID, err := strconv.Atoi(post)
			post, err := models.GetPost(dbConn, postID)
			if err != nil {
				fmt.Println("Failed to get post:", err)
				return
			}
			fmt.Printf("Post: %+v\n", post)

			itin, err := models.GetItinerary(dbConn, post.ItineraryId)
			if err != nil {
				fmt.Println("Failed to get itin:", err)
				return
			}
			itineraries = append(itineraries, itin)
		}

		fmt.Printf("Received Request: %+v\n", itineraries)

		// Respond to the client with the received data
		c.JSON(http.StatusOK, gin.H{
			"message":     "Itinerary received",
			"itineraries": itineraries,
		})

	}
}
