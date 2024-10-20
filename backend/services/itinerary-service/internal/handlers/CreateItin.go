package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

		fmt.Printf("Received Itinerary: Name=%s, City=%s, Country=%s\n", req.Name, req.City, req.Country)

		// Respond to the client with the received data
		c.JSON(http.StatusOK, gin.H{
			"message": "Itinerary received",
			"data":    req,
		})

	}
}
