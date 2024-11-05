package api

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	handlers "github.com/jordyob03/TripTailor/backend/services/save-service/internal/handlers"
)

func test() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log the received data
		fmt.Printf("Received data")

		// Send a simple response back
		c.JSON(200, gin.H{
			"message": "Data received successfully",
		})
	}
}

func RegisterRoutes(r *gin.Engine, dbConn *sql.DB) {
	r.POST("/save-service-test", test())
	r.POST("/boards/add", handlers.AddBoard(dbConn))
	r.DELETE("/boards/delete", handlers.DeleteBoard(dbConn))
	r.POST("/boards/:boardId/posts/add", handlers.AddBoardPost(dbConn))
	r.DELETE("/boards/:boardId/posts/:postId", handlers.DeleteBoardPost(dbConn))
	r.POST("/boards/search", handlers.SearchBoards(dbConn))
	r.GET("/boards/:boardId/posts", handlers.GetBoard(dbConn))
}
