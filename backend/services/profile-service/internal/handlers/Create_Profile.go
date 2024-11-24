package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	db "github.com/jordyob03/TripTailor/backend/services/profile-service/internal/db"
)

// CreateProfile handles the creation (or updating) of a user profile
func CreateProfile(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var profileReq db.User

		// Bind the JSON body to the User struct
		if err := c.ShouldBindJSON(&profileReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
			return
		}

		// Check if the username is provided
		if profileReq.Username == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
			return
		}

		// Attempt to retrieve the user from the database
		user, err := db.GetUser(dbConn, profileReq.Username)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User %s not found", profileReq.Username)})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user data", "details": err.Error()})
			return
		}

		// Parse and validate the provided tags and languages
		parsedLanguages := ParseLang(strings.Join(profileReq.Languages, ","))

		// Update the country if provided
		if profileReq.Country != "" {
			err = db.UpdateUserCountry(dbConn, profileReq.Username, profileReq.Country)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating country", "details": err.Error()})
				return
			}
		}

		// Update the languages if provided
		if len(parsedLanguages) > 0 {
			err = db.AddUserLanguage(dbConn, profileReq.Username, parsedLanguages)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating languages", "details": err.Error()})
				return
			}
		}

		if profileReq.Tags != nil {
			for _, rtag := range user.Tags {
				err = db.RemoveUserTag(dbConn, profileReq.Username, []string{rtag})
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Removing updated tags"})
					return
				}
				println("Removed tag: ", rtag)
			}

			for _, atag := range profileReq.Tags {
				err = db.AddUserTag(dbConn, profileReq.Username, []string{atag})
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Adding updated tags"})
					return
				}
				println("Added tag: ", atag)
			}
		}

		if profileReq.Name != "" {
			err = db.UpdateName(dbConn, profileReq.Username, profileReq.Name)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating name", "details": err.Error()})
				return
			}
		}

		// Respond with success if no errors occurred
		c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully", "user": user})
	}
}
