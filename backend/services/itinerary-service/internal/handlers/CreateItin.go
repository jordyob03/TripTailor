package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	models "github.com/jordyob03/TripTailor/backend/services/itinerary-service/internal/models"
)

type CreateItinRequest struct {
	Name      string   `json:"name"`
	City      string   `json:"city"`
	Country   string   `json:"country"`
	Languages []string `json:"languages"`
	Tags      []string `json:"tags"`
	Events    []string `json:"events"`
	PostId    int      `json:"postId"`
	Username  string   `json:"username"`
}

func CreateItin(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateItinRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("Received Itinerary: %+v\n", req)

		itin := models.Itinerary{
			Name:         req.Name,
			City:         req.City,
			Country:      req.Country,
			Languages:    req.Languages,
			Tags:         req.Tags,
			Events:       req.Events,
			PostId:       req.PostId,
			Username:     req.Username,
			CreationDate: time.Now(),
			LastUpdate:   time.Now(),
		}

		itinId, err := models.AddItinerary(dbConn, itin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create itinerary"})
			return
		}

		// Respond to the client with the received data
		c.JSON(http.StatusOK, gin.H{
			"message": "Itinerary received",
			"itinId":  itinId,
		})

	}
}
