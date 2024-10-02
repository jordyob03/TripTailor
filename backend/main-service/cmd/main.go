// backend/main-service/cmd/main.go
package main

import (
	db "backend/db/models"
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintln(w, "Hello World")
}

func main() {
	connStr := "postgres://postgres:password@db:5432/database?sslmode=disable"

	if err := db.InitDB(connStr); err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.CloseDB()

	if err := db.CreateUserTable(); err != nil {
		log.Fatal("Error creating user table:", err)
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/users", db.UserHandler)

	fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
