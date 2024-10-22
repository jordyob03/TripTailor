package main

import (
	"database/sql"
	"log"

	"github.com/jordyob03/TripTailor/backend/services/search-service/api"

	//"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Connect to the database
	db, err := sql.Open("postgres", "postgres://postgres:password@db:5432/database?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a new Gin router
	router := gin.Default()

	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Set up routes
	api.SetupRoutes(router, db)

	// Start the server on port 8084
	if err := router.Run(":8084"); err != nil {
		log.Fatal(err)
	}
}
