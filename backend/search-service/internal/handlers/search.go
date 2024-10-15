package handlers

import (
	"backend/search-service/internal/db"
	"database/sql"
	"encoding/json"
	"net/http"
)

// SearchItineraries returns an HTTP handler function with a database connection
func SearchItineraries(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		country := r.URL.Query().Get("country")
		city := r.URL.Query().Get("city")

		if country == "" || city == "" {
			http.Error(w, "Missing query parameters", http.StatusBadRequest)
			return
		}

		// Use the database connection passed from the main function
		itineraries, err := db.QueryItinerariesByLocation(database, country, city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(itineraries)
	}
}
