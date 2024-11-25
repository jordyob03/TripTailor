package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jordyob03/TripTailor/backend/services/profile-service/internal/db"
)

// GetProfile handles retrieving a user profile based on the username
func GetProfile(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("username")
		fmt.Println("Received username:", username)

		// Check if the username is provided
		if username == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
			return
		}

		// Attempt to retrieve the user from the database
		user, err := db.GetUser(dbConn, username)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User %s not found", username)})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user data", "details": err.Error()})
			return
		}

		// Respond with the retrieved user profile data
		c.JSON(http.StatusOK, gin.H{
			"username":  user.Username,
			"name":      user.Name,
			"country":   user.Country,
			"tags":      user.Tags,
			"languages": user.Languages,
		})
	}
}
