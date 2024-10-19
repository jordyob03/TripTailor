package api

import (
	"backend/search-service/internal/handlers"
	"database/sql"
	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes the routes for the service, passing the database connection
func SetupRoutes(router *gin.Engine, db *sql.DB) {
	// Define the /search route and pass the database connection
	router.GET("/search", handlers.SearchItineraries(db))
}
