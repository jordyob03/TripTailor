package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jordyob03/TripTailor/backend/services/feed-service/api"
	_ "github.com/lib/pq"
)

func main() {
	// Set up database connection
	connStr := "postgres://postgres:password@db:5432/database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Set up router
	router := gin.Default()

	// Register the feed route
	api.SetupRoutes(router, db)

	// Start the server
	if err := router.Run(":8093"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
