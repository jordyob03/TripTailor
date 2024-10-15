package api

import (
	CreateProfilehandlers "backend/profile-service/internal/handlers/createprofile"
	UpdateProfilehandlers "backend/profile-service/internal/handlers/updateprofile"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all routes for the profile service
func RegisterRoutes(router *gin.Engine) {
	router.POST("/profile", CreateProfilehandlers.CreateProfile())    // Ensure the function returns gin.HandlerFunc
	router.PUT("/profile/:id", UpdateProfilehandlers.UpdateProfile()) // Use consistent route parameter
}
