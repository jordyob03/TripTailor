package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/jordyob03/TripTailor/backend/services/save-service/internal/db"
)

func DeleteBoard(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		boardId, err := strconv.Atoi(c.Param("boardId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid boardId"})
			return
		}

		err = db.RemoveBoard(dbConn, boardId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
	}
}
