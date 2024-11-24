package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jordyob03/TripTailor/backend/services/save-service/internal/db"
)

type AddBoardPostRequest struct {
	BoardId int `json:"boardId" form:"boardId"`
	PostId  int `json:"postId" form:"postId"`
}

func AddBoardPost(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AddBoardPostRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := db.AddBoardPost(dbConn, req.BoardId, req.PostId, true)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Board post added successfully"})

	}
}
