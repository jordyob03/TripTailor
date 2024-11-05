package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/jordyob03/TripTailor/backend/services/save-service/internal/db"
)

type AddBoardRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Username    string `json:"username"`
}

func AddBoard(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AddBoardRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		board := db.Board{
			Name:         req.Name,
			Description:  req.Description,
			Username:     req.Username,
			CreationDate: time.Now(),
		}

		boardId, err := db.AddBoard(dbConn, board)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add board"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Board added successfully", "boardId": boardId})
	}
}
