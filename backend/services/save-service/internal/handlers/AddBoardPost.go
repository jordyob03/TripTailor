package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jordyob03/TripTailor/backend/services/save-service/internal/db"
)

type AddBoardPostRequest struct {
	BoardId int `json:"boardId"`
	PostId  int `json:"postId"`
}

func AddBoardPost(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AddBoardPostRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		err := db.AddBoardPost(dbConn, req.BoardId, req.PostId, true)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add post to board"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Post added successfully", "boardId": req.BoardId, "postId": req.PostId})
	}
}
