package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	handlers "github.com/jordyob03/TripTailor/backend/services/itinerary-service/internal/handlers"
)

// func test(c *gin.Context) {
// 	fmt.Println("Hello, a POST request was made to /itin-creation")

// 	// Respond to the client
// 	c.JSON(200, gin.H{
// 		"message": "Itinerary created",
// 	})
// }

func RegisterRoutes(r *gin.Engine, dbConn *sql.DB) {
	r.POST("/itin-creation", handlers.CreateItin(dbConn))
}
