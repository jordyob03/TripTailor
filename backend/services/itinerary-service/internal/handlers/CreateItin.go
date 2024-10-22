package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateItinRequest struct {
	Name        string   `json:"name" binding:"required"`
	Location    string   `json:"location" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Cost        string   `json:"cost" binding:"required"`
	Tags        []string `json:"tags" binding:"required"`
}

func CreateItin(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateItinRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//var itin models.Itinerary

		// test itin

		fmt.Printf("Received Itinerary: %+v\n", req)

		// itin.Name = req.Name
		// itin.City = req.City
		// itin.Country = req.City
		// itin.Tags = req.Tags
		// itin.Events = []string{}
		// itin.Username = "jordyob"

		//Needs to be added to db here

		// Respond to the client with the received data
		c.JSON(http.StatusOK, gin.H{
			"message": "Itinerary received",
		})

	}
}
