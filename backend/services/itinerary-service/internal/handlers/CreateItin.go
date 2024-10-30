package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/jordyob03/TripTailor/backend/services/itinerary-service/internal/models"
)

type Event struct {
	Name        string  `json:"name"`
	StartTime   string  `json:"startTime"`
	EndTime     string  `json:"endTime"`
	Location    string  `json:"location"`
	Description string  `json:"description"`
	Cost        float64 `json:"cost"`
}

type CreateItinRequest struct {
	Name        string   `json:"name"`
	City        string   `json:"city"`
	Country     string   `json:"country"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Languages   []string `json:"languages"`
	Tags        []string `json:"tags"`
	Events      []Event  `json:"events"`
	PostId      int      `json:"postId"`
	Username    string   `json:"username"`
	ImageData   []byte   `json:"imageData"`
}

func CreateItin(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateItinRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("Received Itinerary: %+v\n", req)

		eventNames := make([]string, len(req.Events))
		for i, event := range req.Events {
			eventNames[i] = event.Name
		}

		var imageId int
		if len(req.ImageData) > 0 {
			image := models.Image{
				ImageData: req.ImageData,
				Metadata:  []string{"Itinerary image"},
			}

			var err error
			imageId, err = models.AddImage(dbConn, image)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store image"})
				return
			}
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
			Events:      eventNames,
			PostId:      req.PostId,
			Username:    req.Username,
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
			"imageId": imageId,
		})

	}
}
