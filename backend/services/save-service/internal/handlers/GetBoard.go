package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jordyob03/TripTailor/backend/services/save-service/internal/db"
)

type GetBoardRequest struct {
	Username string `json:"username" form:"username"`
}

func GetBoard(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetBoardRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := db.GetUser(dbConn, req.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
			return
		}

		Intboards, err := db.StringsToInts(user.Boards)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert board IDs"})
			return
		}

		boards := []db.Board{}

		for _, board := range Intboards {
			board, err := db.GetBoard(dbConn, board)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve board"})
				return
			}
			boards = append(boards, board)
		}

		c.JSON(http.StatusOK, gin.H{"boards": boards})
	}
}
