package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jordyob03/TripTailor/backend/services/save-service/internal/db"
)

type GetPostRequest struct {
	boardId int `json:"boardId" form:"boardId"`
}

func GetPost(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetPostRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		Board, err := db.GetBoard(dbConn, req.boardId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Board"})
			return
		}

		IntPosts, err := db.StringsToInts(Board.Posts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert post IDs"})
			return
		}

		Posts := []db.Post{}

		for _, post := range IntPosts {
			post, err := db.GetPost(dbConn, post)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve post"})
				return
			}
			Posts = append(Posts, post)
		}

		c.JSON(http.StatusOK, gin.H{"Posts": Posts})
	}
}
