package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/jordyob03/TripTailor/backend/services/feed-service/internal/models"
)

type GetPostRequest struct {
	UserId string `json:"UserId" form:"UserId" binding:"required"`
}

func GetPost(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetPostRequest

		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		User, err := models.GetUser(dbConn, req.UserId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve User"})
			return
		}

		IntPosts, err := models.StringsToInts(User.Posts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert post IDs"})
			return
		}

		Posts := []models.Post{}
		for _, post := range IntPosts {
			post, err := models.GetPost(dbConn, post)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve post"})
				return
			}
			Posts = append(Posts, post)
		}

		c.JSON(http.StatusOK, gin.H{"Posts": Posts})
	}
}
