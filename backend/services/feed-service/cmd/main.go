package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jordyob03/TripTailor/backend/services/feed-service/api"
	"github.com/jordyob03/TripTailor/backend/services/feed-service/handlers"
	_ "github.com/lib/pq"
)

func main() {
	// Set up database connection
	connStr := "user=username dbname=trip_tailor sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Set up router
	router := gin.Default()

	// Register the feed route
	api.SetupRoutes("/feed", handlers.FeedService(db))

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
