package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/jordyob03/TripTailor/backend/services/itinerary-service/internal/models"
)

type CreateEventRequest struct {
	Name        string `json:"name"`
	Location    string `json:"location"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	Description string `json:"description"`
	Cost        string `json:"cost"`
}

type CreateItinRequest struct {
	Name        string               `json:"name"`
	City        string               `json:"city"`
	Country     string               `json:"country"`
	Description string               `json:"description"`
	Tags        []string             `json:"tags"`
	Events      []CreateEventRequest `json:"events"`
	Username    string               `json:"username"`
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
			Name:        req.Name,
			City:        req.City,
			Country:     req.Country,
			Description: req.Description,
			Tags:        req.Tags,
			Username:    req.Username,
		}

		// Add itin to db without events
		itinId, err := models.AddItinerary(dbConn, itin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create itinerary"})
			return
		}

		// Iterate over events, create event then add to event db to get event ID
		//eventIdStrings := []string{}
		for _, event := range req.Events {
			newEvent := models.Event{
				Name:    event.Name,
				Address: event.Location,

				// I commented these out because it won't work until these are converted to the right data type time and float I think
				//StartTime: event.StartTime,
				//EndTime: event.EndTime,
				Description: event.Description,
				//Cost: event.Cost,
				ItineraryId: itinId,
			}
			eventId, err := models.AddEvent(dbConn, newEvent)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to add event %s", event.Name)})
				return
			}
			err = models.AddItineraryEvent(dbConn, itinId, eventId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to link event %s to itinerary", event.Name)})
				return
			}
		}

		// Respond to the client with the received data
		c.JSON(http.StatusOK, gin.H{
			"message": "Itinerary received",
			"itinId":  itinId,
		})

	}
}
