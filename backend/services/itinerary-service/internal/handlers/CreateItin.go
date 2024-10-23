package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jordyob03/TripTailor/backend/services/itinerary-service/internal/db"
	models "github.com/jordyob03/TripTailor/backend/services/itinerary-service/internal/models"
)

type Event struct {
	Name        string `json:"name"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Cost        string `json:"cost"`
}

type CreateItinRequest struct {
	Name        string   `json:"name" binding:"required"`
	User        string   `json:"username"`
	Location    string   `json:"location" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Cost        string   `json:"cost" binding:"required"`
	Tags        []string `json:"tags" binding:"required"`
	Events      []Event  `json:"events"`
}

func CreateItin(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateItinRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var itin models.Itinerary

		fmt.Printf("Received Itinerary: %+v\n", req)

		itin.Name = req.Name
		itin.City = req.Location
		itin.Country = req.Location
		itin.Tags = req.Tags
		itin.Username = req.User

		itinId, err := db.AddItinerary(dbConn, itin)
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
