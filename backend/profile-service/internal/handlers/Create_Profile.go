package handlers

import (
	"backend/profile-service/internal/db"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// CreateProfile handles the creation (actually updating) of a user profile
func CreateProfile(DB *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var profileReq CreateProfileRequest

		// Bind the JSON body to the CreateProfileRequest struct
		if err := c.ShouldBindJSON(&profileReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Check if the user exists in the database
		username := c.Param("username")    // Pull username from URL param
		_, err := db.GetUser(DB, username) // No need to store existingUser if not used
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User %s not found", username)})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user data"})
			return
		}

		// Parse tags and languages from the profile request
		parsedTags := ParseTags(profileReq.Tags)
		parsedLanguages := strings.Split(profileReq.Language, ",")

		// Update only the fields provided by the request
		updates := make(map[string]interface{})
		if profileReq.Name != "" {
			updates["name"] = profileReq.Name
		}
		if profileReq.Country != "" {
			updates["country"] = profileReq.Country
		}
		if len(parsedLanguages) > 0 && parsedLanguages[0] != "" { // Ensure it's not an empty list
			updates["languages"] = pq.Array(parsedLanguages)
		}
		if len(parsedTags) > 0 && parsedTags[0] != "" {
			updates["tags"] = pq.Array(parsedTags)
		}

		if len(updates) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No fields to update"})
			return
		}

		// Update the user in the database
		err = db.UpdateRow(DB, "users", updates, "username = $1", username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user profile"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
	}
}
