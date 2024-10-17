package routes

import (
	"backend/profile-service/internal/handlers"
	"database/sql"

	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up all the routes for the application.
func SetupRoutes(router *gin.Engine, dbConn *sql.DB) {
	// Group for user-related routes
	userGroup := router.Group("/user")
	{
		// Create or update a user profile
		userGroup.POST("/create", handlers.CreateProfile(dbConn))

		// Update an existing user profile
		userGroup.PUT("/update", handlers.UpdateProfile(dbConn))
	}
}
