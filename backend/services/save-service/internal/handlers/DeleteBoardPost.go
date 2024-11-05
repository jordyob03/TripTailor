package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jordyob03/TripTailor/backend/services/save-service/internal/db"
)

type DeleteBoardPostRequest struct {
	BoardId int `json:"boardId"`
	PostId  int `json:"postId"`
}

func DeleteBoardPost(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req DeleteBoardPostRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if err := db.RemoveBoardPost(dbConn, req.BoardId, req.PostId); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
	}
}
