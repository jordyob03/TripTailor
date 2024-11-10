package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	db "github.com/jordyob03/TripTailor/backend/services/main-service/internal/db/models"
	pack "github.com/jordyob03/TripTailor/backend/services/main-service/utils"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintln(w, "Hello World")
}

func main() {
	connStr := "postgres://postgres:password@db:5432/database?sslmode=disable"
	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer DB.Close()

	if err := db.InitDB(DB, connStr); err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.CloseDB(DB)

	if err := db.DeleteAllTables(DB); err != nil {
		log.Fatal("Error deleting tables:", err)
	}

	if err := db.CreateAllTables(DB); err != nil {
		log.Fatal("Error creating tables:", err)
	}

	if err := db.InitializeImageTable(DB); err != nil {
		log.Fatal("Error initializing image table:", err)
	}

	pack.PackImagesFromLocal("utils/packed_data/images", DB)
	pack.PackUsersFromJSON("utils/packed_data/users.json", DB)
	pack.PackEventFromJSON("utils/packed_data/events.json", DB)
	pack.PackItinsAndPostFromJSON("utils/packed_data/itineraries.json", DB)
	pack.PackBoardsFromJSON("utils/packed_data/boards.json", DB)

	http.HandleFunc("/images/", db.ImageHandler(DB))
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/users", db.UserHandlerWrapper(DB))

	fmt.Println("Starting server on port 8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
