package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/jordyob03/TripTailor/backend/services/save-service/internal/db"
)

type AddBoardRequest struct {
	Username  string `json:"username" form:"username"`
	BoardName string `json:"boardname" form:"boardname"`
}

func AddBoard(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AddBoardRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		NewBoard := db.Board{
			Name:         req.BoardName,
			CreationDate: time.Now(),
			Description:  "",
			Username:     req.Username,
			Posts:        []string{},
			Tags:         []string{},
		}

		boardId, err := db.AddBoard(dbConn, NewBoard)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"boardId": boardId})
	}
}
