package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/jordyob03/TripTailor/backend/services/itinerary-service/internal/models"
)

type CreateItinRequest struct {
	Name        string         `json:"name"`
	City        string         `json:"city"`
	Country     string         `json:"country"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	Languages   []string       `json:"languages"`
	Tags        []string       `json:"tags"`
	Events      []models.Event `json:"events"`
	PostId      int            `json:"postId"`
	Username    string         `json:"username"`
}

func CreateItin(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateItinRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("Received Itinerary: %+v\n", req)

		eventIdStrings := []string{}
		for _, event := range req.Events {
			eventIdStrings = append(eventIdStrings, fmt.Sprintf("%d", event.EventId)) // Convert EventId to string
		}

		itin := models.Itinerary{
			Name:        req.Name,
			City:        req.City,
			Country:     req.Country,
			Title:       req.Title,
			Description: req.Description,
			Price:       req.Price,
			Languages:   req.Languages,
			Tags:        req.Tags,
			Events:      eventIdStrings,
			PostId:      req.PostId,
			Username:    req.Username,
		}

		itinId, err := models.AddItinerary(dbConn, itin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create itinerary"})
			return
		}

		// Iterate over each event, set its ItineraryId, and add it to the database
		for _, event := range req.Events {
			// Set the ItineraryId for each event
			event.ItineraryId = itinId

			// Add the event to the database and retrieve the eventId
			eventId, err := models.AddEvent(dbConn, event)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to add event %s", event.Name)})
				return
			}

			// Link the event to the itinerary by adding the eventId to the itinerary's list of events
			err = models.AddItineraryEvent(dbConn, itinId, eventId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to link event %s to itinerary", event.Name)})
				return
			}
		}

		err = models.UpdateItineraryPrice(dbConn, itinId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update itinerary price"})
			return
		}

		// Respond to the client with the received data
		c.JSON(http.StatusOK, gin.H{
			"message": "Itinerary received",
			"itinId":  itinId,
		})

	}
}
