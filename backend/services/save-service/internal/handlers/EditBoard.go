package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jordyob03/TripTailor/backend/services/save-service/internal/db"
)

type EditBoardRequest struct {
	BoardId     int    `json:"boardId" form:"boardId"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}

func EditBoard(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req EditBoardRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.BoardId == 0 || req.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "boardId and name are required fields"})
			return
		}

		err := db.UpdateBoardName(dbConn, req.BoardId, req.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = db.UpdateBoardDescription(dbConn, req.BoardId, req.Description)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"boardId": req.BoardId})
	}
}
