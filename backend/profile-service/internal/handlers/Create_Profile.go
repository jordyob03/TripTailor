package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"backend/profile-service/internal/db"

	"github.com/gin-gonic/gin"
)

// CreateProfile handles the creation (or updating) of a user profile
func CreateProfile(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var profileReq DBmodels.User

		// Bind the JSON body to the User struct
		if err := c.ShouldBindJSON(&profileReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		username := profileReq.Username
		_, err := DBmodels.GetUser(dbConn, username)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User %s not found", username)})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user data"})
			return
		}

		parsedTags := ParseTags(strings.Join(profileReq.Tags, ","))
		parsedLanguages := profileReq.Languages

		// Update the country if provided
		if profileReq.Country != "" {
			err = DBmodels.UpdateUserCountry(dbConn, username, profileReq.Country)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating country"})
				return
			}
		}

		// Update the languages if provided
		if len(parsedLanguages) > 0 {
			err = DBmodels.AddUserLanguage(dbConn, username, parsedLanguages)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating languages"})
				return
			}
		}

		// Update the tags if provided
		if len(parsedTags) > 0 {
			err = DBmodels.AddUserTag(dbConn, username, parsedTags)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating tags"})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
	}
}
