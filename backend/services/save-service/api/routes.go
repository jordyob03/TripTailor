package api

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	handlers "github.com/jordyob03/TripTailor/backend/services/save-service/internal/handlers"
)

func test() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Log the received data
		fmt.Printf("Received data")

		// Send a simple response back
		c.JSON(200, gin.H{
			"message": "Data received successfully",
		})
	}
}

func RegisterRoutes(r *gin.Engine, dbConn *sql.DB) {
	r.POST("/itin-creation", handlers.CreateItin(dbConn))
}
