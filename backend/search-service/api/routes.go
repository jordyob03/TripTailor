package api

import (
	"backend/search-service/internal/handlers"
	"database/sql"
	"github.com/gorilla/mux"
	//"github.com/gin-gonic/gin"
)

// SetupRoutes initializes the routes for the service, passing the database connection
func SetupRoutes(router *mux.Router, db *sql.DB) {
	// Define a route for searching itineraries and pass the database connection
	router.HandleFunc("/search", handlers.SearchItineraries(db)).Methods("GET")
}

// func RegisterRoutes(r *gin.Engine, dbConn *sql.DB) {
// 	r.POST("/signup", handlers.SignUp(dbConn))
// 	r.POST("/signin", handlers.SignIn(dbConn))
// }
