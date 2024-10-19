package main

import (
	"backend/search-service/api"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQL driver
	"log"
)

func main() {
	// Initialize the PostgreSQL connection
	connStr := "user=username dbname=itinerariesdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a new Gin router
	router := gin.Default()

	// Set up the /search route with the database connection
	api.SetupRoutes(router, db)

	// Start the HTTP server on port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
