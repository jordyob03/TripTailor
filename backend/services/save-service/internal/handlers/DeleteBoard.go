package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jordyob03/TripTailor/backend/services/save-service/internal/db"
)

type DeleteBoardRequest struct {
	BoardId int `json:"boardId"`
}

func DeleteBoard(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req DeleteBoardRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.RemoveBoard(dbConn, req.BoardId); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete board"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Board deleted successfully"})
	}
}
