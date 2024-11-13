package main

import (
	"database/sql"

	"log"

	api "github.com/jordyob03/TripTailor/backend/services/save-service/api"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Save"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	connStr := "postgres://postgres:password@db:5432/database?sslmode=disable"
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer dbConn.Close()

	api.RegisterRoutes(r, dbConn)

	r.Run(":8086")

}
