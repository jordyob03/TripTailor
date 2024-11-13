package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/jordyob03/TripTailor/backend/services/feed-service/internal/models"
)

type GetEventRequest struct {
	ItineraryId int `json:"itineraryId" form:"itineraryId" binding:"required"`
}

func GetEvent(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetEventRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		Itinerary, err := models.GetItinerary(dbConn, req.ItineraryId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve "})
			return
		}

		IntEvents, err := models.StringsToInts(Itinerary.Events)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert event IDs"})
			return
		}

		Events := []models.Event{}
		for _, event := range IntEvents {
			event, err := models.GetEvent(dbConn, event)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
				return
			}
			Events = append(Events, event)
		}

		c.JSON(http.StatusOK, gin.H{"Events": Events})
	}
}
