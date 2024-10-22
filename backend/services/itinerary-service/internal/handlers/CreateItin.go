package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	DBmodels "github.com/jordyob03/TripTailor/backend/services/itinerary-service/internal/db"
	models "github.com/jordyob03/TripTailor/backend/services/itinerary-service/internal/models"
)

type CreateItinRequest struct {
	Name    string `json:"Name" binding:"required"`
	City    string `json:"City" binding:"required"`
	Country string `json:"Country" binding:"required"`
}

func CreateItin(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateItinRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var itin models.Itinerary

		// test itin

		itin.Name = req.Name
		itin.City = req.City
		itin.Country = req.Country
		itin.Languages = []string{"English", "French"}
		itin.Tags = []string{"Family", "Vegetarian"}
		itin.Events = []string{}
		itin.Username = "jordyob"

		//Try adding Itin to the database where itinID
		//is the ID of the newly created itinerary
		itinID, err := DBmodels.AddItinerary(dbConn, itin)
		if err != nil {
			fmt.Printf("Error adding itinnerary to the database %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to create itinerary",
				"details": err.Error(),
			})
			return
		}

		//ENsure that the ItinID is returned from the database
		if itinID == 0 {
			fmt.Printf("Failed to retreve itinerary ID")
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Failed to create itinerary",
				"detail": "Invalid itinerary ID returned",
			})
			return
		}

		//Needs to be added to db here

		fmt.Printf("Received Itinerary: Name=%s, City=%s, Country=%s\n", req.Name, req.City, req.Country)

		// Respond to the client with the received data
		c.JSON(http.StatusOK, gin.H{
			"message":     "Itinerary received",
			"ItineraryID": itinID,
		})

	}
}
