package main

import (
	"backend/search-service/api" // Import the new routes package
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	// Initialize the PostgreSQL connection
	connStr := "user=username dbname=itinerariesdb sslmode=disable"
	database, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	// Initialize the router
	router := mux.NewRouter()

	// Set up the routes and pass the database connection
	api.SetupRoutes(router, database)

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", router))
}
