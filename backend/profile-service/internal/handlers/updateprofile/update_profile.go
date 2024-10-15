package updateprofile

import (
	"backend/profile-service/internal/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler for updating a profile.
func UpdateProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var updatedProfile handlers.CreateProfileRequest

		// Bind and validate the request body.
		if err := c.ShouldBindJSON(&updatedProfile); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the profile exists.
		existingProfile, exists := handlers.Profiles[id]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
			return
		}

		// Ensure at least 3 tags are provided.
		tags := handlers.ParseTags(updatedProfile.Tags)
		if len(tags) < 3 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "At least 3 tags must be selected"})
			return
		}

		// Update the profile fields.
		existingProfile.Language = updatedProfile.Language
		existingProfile.Country = updatedProfile.Country
		existingProfile.Tags = updatedProfile.Tags
		existingProfile.Name = updatedProfile.Name
		existingProfile.Username = updatedProfile.Username

		// Store the updated profile.
		handlers.Profiles[id] = existingProfile

		// Respond with the updated profile.
		c.JSON(http.StatusOK, gin.H{
			"message": "Profile updated successfully",
			"profile": existingProfile,
		})
	}
}
