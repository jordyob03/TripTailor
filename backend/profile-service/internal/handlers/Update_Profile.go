package handlers

import (
	"backend/profile-service/internal/db"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq" // Added pq package for array handling in PostgreSQL
)

// UpdateProfile handles updating a user profile
func UpdateProfile(DB *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username") // Pull username from URL param
		var profileReq CreateProfileRequest

		// Bind the JSON body to the CreateProfileRequest struct
		if err := c.ShouldBindJSON(&profileReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Fetch the user from the database
		user, err := db.GetUser(DB, username) // Changed from DBAuth.GetUser to db.GetUser
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User %s not found", username)})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user"})
			return
		}

		// Update the fields only if they are provided in the request
		if profileReq.Name != "" {
			user.Name = profileReq.Name
		}

		if profileReq.Country != "" {
			user.Country = profileReq.Country
		}

		if profileReq.Language != "" {
			user.Languages = strings.Split(profileReq.Language, ",")
		}

		if profileReq.Tags != "" {
			user.Tags = ParseTags(profileReq.Tags)
		}

		// Update user in the database
		err = db.UpdateRow(DB, "users", map[string]interface{}{
			"name":      user.Name,
			"country":   user.Country,
			"languages": pq.Array(user.Languages),
			"tags":      pq.Array(user.Tags),
		}, "username = $1", username)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user profile"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
	}
}
