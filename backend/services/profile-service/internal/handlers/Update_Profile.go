package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	db "github.com/jordyob03/TripTailor/backend/services/profile-service/internal/db"

	"github.com/gin-gonic/gin"
)

// UpdateProfile handles updating a user profile
func UpdateProfile(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var profileReq CreateProfileRequest

		// Bind the JSON body to the CreateProfileRequest struct
		if err := c.ShouldBindJSON(&profileReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		username := profileReq.Username
		if username == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
			return
		}

		// Fetch the user from the database
		user, err := db.GetUser(dbConn, username)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User %s not found", username)})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user"})
			return
		}

		// Update country if provided
		if profileReq.Country != "" {
			err = db.UpdateUserCountry(dbConn, username, profileReq.Country)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating country"})
				return
			}
		}

		// Update languages if provided
		if profileReq.Language != "" {
			user.Languages = strings.Split(profileReq.Language, ",")
			err = db.AddUserLanguage(dbConn, username, user.Languages)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating languages"})
				return
			}
		}

		// Update tags if provided
		if profileReq.Tags != "" {
			user.Tags = ParseTags(profileReq.Tags)
			err = db.AddUserTag(dbConn, username, user.Tags)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating tags"})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
	}
}
