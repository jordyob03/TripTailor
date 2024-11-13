package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jordyob03/TripTailor/backend/services/save-service/internal/db"
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

		Itinerary, err := db.GetItinerary(dbConn, req.ItineraryId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve "})
			return
		}

		IntEvents, err := db.StringsToInts(Itinerary.Events)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert event IDs"})
			return
		}

		Events := []db.Event{}
		for _, event := range IntEvents {
			event, err := db.GetEvent(dbConn, event)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
				return
			}
			Events = append(Events, event)
		}

		c.JSON(http.StatusOK, gin.H{"Events": Events})
	}
}
