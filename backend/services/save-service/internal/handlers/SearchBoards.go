package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	db "github.com/jordyob03/TripTailor/backend/services/save-service/internal/models"
)

type SearchBoardRequest struct {
	Username   string `json:"username"`
	SearchTerm string `json:"searchTerm"`
}

func SearchBoards(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req SearchBoardRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		finalBoards, err := SearchBoardData(dbConn, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search boards"})
			return
		}

		c.JSON(http.StatusOK, finalBoards)
	}
}

func SearchBoardData(dbConn *sql.DB, req SearchBoardRequest) ([]db.Board, error) {
	var finalBoards []db.Board

	User, err := db.GetUser(dbConn, req.Username)
	if err != nil {
		return nil, err
	}

	intBoards, err := db.StringsToInts(User.Boards)
	if err != nil {
		return nil, err
	}

	for _, boardID := range intBoards {
		board, err := db.GetBoard(dbConn, boardID)
		if err != nil {
			return nil, err
		}

		for _, i := range []string{board.Name, board.Description} {
			if containsSubstring(i, req.SearchTerm) {
				finalBoards = append(finalBoards, board)
				break
			}
		}

		for _, tag := range board.Tags {
			if containsSubstring(tag, req.SearchTerm) {
				finalBoards = append(finalBoards, board)
				break
			}
		}

		intPosts, err := db.StringsToInts(board.Posts)
		if err != nil {
			return nil, err
		}

		for _, postID := range intPosts {
			post, err := db.GetPost(dbConn, postID)
			if err != nil {
				return nil, err
			}

			itinerary, err := db.GetItinerary(dbConn, post.ItineraryId)
			if err != nil {
				return nil, err
			}

			for _, i := range []string{itinerary.Name, itinerary.City, itinerary.Country} {
				if containsSubstring(i, req.SearchTerm) {
					finalBoards = append(finalBoards, board)
					break
				}
			}
		}

	}

	return finalBoards, nil

}

func containsSubstring(mainString, substring string) bool {
	escapedSubstring := regexp.QuoteMeta(substring)

	pattern := fmt.Sprintf(".*%s.*", escapedSubstring)
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}

	return re.MatchString(mainString)
}
