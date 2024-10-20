package api

import (
	"database/sql"
	"fmt"

	//handlers "github.com/jordyob03/TripTailor/backend/services/itinerary-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {

	fmt.Print("Hello")

	c.JSON(200, gin.H{
		"message": "Hello World",
	})

}

func RegisterRoutes(r *gin.Engine, dbConn *sql.DB) {
	r.POST("/itin-creation", test)
}
