package api

import (
	"database/sql"
	//"fmt"

	"github.com/gin-gonic/gin"
	handlers "github.com/jordyob03/TripTailor/backend/services/itinerary-service/internal/handlers"
)

// func test() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		// Log the received data
// 		fmt.Printf("Received data")

// 		// Send a simple response back
// 		c.JSON(200, gin.H{
// 			"message": "Data received successfully",
// 		})
// 	}
// }

func RegisterRoutes(r *gin.Engine, dbConn *sql.DB) {
	r.POST("/itin-creation", handlers.CreateItin(dbConn))
	r.POST("/get-user-itins", handlers.GetItins(dbConn))
}
