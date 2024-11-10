package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

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

		for _, event := range req.Events {

			fmt.Print(event.StartTime)

			//If these work as I think they will, both start and end time should be converted into time objects now, also layout has to
			//be "3:04 PM" as this is the reference time for the package. shoulkd handle tiem strings in hh:mm AM/PM format
			layout := "3:04 PM"
			startTime, err := time.Parse(layout, event.StartTime)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid start time format for event %s", event.Name)})
				return
			}

			endTime, err := time.Parse(layout, event.EndTime)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid end time format for event %s", event.Name)})
				return
			}

			eventCost, err := strconv.ParseFloat(event.Cost, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid cost format for event %s", event.Name)})
				return
			}

			newEvent := models.Event{
				Name:    event.Name,
				Address: event.Location,
				// I commented these out because it won't work until these are converted to the right data type time and float I think
				StartTime:   startTime,
				EndTime:     endTime,
				Description: event.Description,
				Cost:        eventCost,
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

		// Update price (this won't work until the data types are fixed)

		// err = models.UpdateItineraryPrice(dbConn, itinId)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed update price"})
		// 	return
		// }

		post := models.Post{
			ItineraryId:  itinId,
			CreationDate: time.Now(),
			Username:     req.Username,
		}

		// Create post
		postId, err := models.AddPost(dbConn, post)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
			return
		}

		models.AddUserPost(dbConn, req.Username, postId)

		// Respond to the client with the received data
		c.JSON(http.StatusOK, gin.H{
			"message": "Itinerary received",
			"itinId":  itinId,
			"postId":  postId,
		})

	}
}
