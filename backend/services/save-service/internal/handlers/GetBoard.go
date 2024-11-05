package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jordyob03/TripTailor/backend/services/save-service/internal/db"
)

type GetBoardRequest struct {
	BoardId int `json:"boardId"`
}

func GetBoard(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetBoardRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		board, err := db.GetBoard(dbConn, req.BoardId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve board"})
			return
		}

		var fullPosts []db.Post
		intPosts, err := db.StringsToInts(board.Posts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert post IDs"})
			return
		}

		for _, postID := range intPosts {
			post, err := db.GetPost(dbConn, postID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve post"})
				return
			}
			fullPosts = append(fullPosts, post)
		}

		c.JSON(http.StatusOK, gin.H{"posts": fullPosts})
	}
}
