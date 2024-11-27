package api

import (
	"database/sql"

	"github.com/jordyob03/TripTailor/backend/services/feed-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes the routes for the service, passing the database connection
func SetupRoutes(router *gin.Engine, dbConn *sql.DB) {
	// Define the /search route and pass the database connection
	router.GET("/feed", handlers.FeedService(dbConn))
	router.GET("/itinerary", handlers.GetItinerary(dbConn))
	router.GET("/posts", handlers.GetPost(dbConn))
	router.GET("/events", handlers.GetEvent(dbConn))
}
