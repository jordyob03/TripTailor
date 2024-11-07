package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	handlers "github.com/jordyob03/TripTailor/backend/services/save-service/internal/handlers"
)

func RegisterRoutes(r *gin.Engine, dbConn *sql.DB) {
	r.GET("/boards", handlers.GetBoard(dbConn))
	r.GET("/posts", handlers.GetPost(dbConn))
	r.GET("/itineraries", handlers.GetItinerary(dbConn))
	r.GET("/events", handlers.GetEvent(dbConn))
	r.DELETE("/boards/:boardId/posts/:postId", handlers.DeleteBoardPost(dbConn))
	r.DELETE("/boards/:boardId", handlers.DeleteBoard(dbConn))
	r.POST("/boards", handlers.CreateBoard(dbConn))
}
