package api

import (
	"backend/profile-service/internal/handlers"
	"database/sql"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all routes for the profile service
func RegisterRoutes(router *gin.Engine, DB *sql.DB) {
	// Route to create a new profile
	router.POST("/profile/:username", handlers.CreateProfile(DB)) // Ensure username is part of the URL path

	// Route to update an existing profile
	router.PUT("/profile/:username", handlers.UpdateProfile(DB)) // Use the same username parameter for consistency
}
