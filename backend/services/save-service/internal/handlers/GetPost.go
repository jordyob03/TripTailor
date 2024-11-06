package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jordyob03/TripTailor/backend/services/save-service/internal/db"
)

type GetPostRequest struct {
	boardId int `form:"boardId"` // Use 'form' tag for query parameters
}

func GetPost(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetPostRequest

		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("Received boardId:", req.boardId)

		// Validate that boardId is greater than 0
		// if req.boardId <= 0 {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid boardId 0"})
		// 	return
		// }

		if req.boardId <= 0 {
			req.boardId = 1
		}

		// Fetch the board from the database
		Board, err := db.GetBoard(dbConn, req.boardId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Board"})
			return
		}

		// Convert the posts from string to int
		IntPosts, err := db.StringsToInts(Board.Posts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert post IDs"})
			return
		}

		// Fetch the posts for the board
		Posts := []db.Post{}
		for _, post := range IntPosts {
			post, err := db.GetPost(dbConn, post)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve post"})
				return
			}
			Posts = append(Posts, post)
		}

		// Return the posts to the client
		c.JSON(http.StatusOK, gin.H{"Posts": Posts})
	}
}
