package main

import (
	api "backend/profile-service/api"
	"database/sql"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Database connection setup
	connStr := "postgres://postgres:password@localhost:5432/database?sslmode=disable"
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer dbConn.Close()

<<<<<<< HEAD
	// Ensure the connection is valid by pinging the database
	if err := dbConn.Ping(); err != nil {
		log.Fatal("Failed to ping the database:", err)
	}
=======
	api.SetupRoutes(r, connStr)
>>>>>>> origin/TT-24-njrc-Create-Profile

	// Pass the dbConn (the actual connection, not the connection string)
	api.SetupRoutes(r, dbConn)

	// Run the server
	r.Run(":8085")
}
