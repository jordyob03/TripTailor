package main

import (
	db "backend/db/models"
	"fmt"
	"log"
	"net/http"
	"time"
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

	if err := db.DeleteTable("users"); err != nil {
		log.Fatal("Error deleting user table:", err)
	}

	if err := db.CreateUserTable(); err != nil {
		log.Fatal("Error creating user table:", err)
	}

	// Test user
	user := db.User{
		Username:    "WyrdWyn4",
		Email:       "wmksherwani@mun.ca",
		Password:    "password",
		DateOfBirth: time.Date(2004, time.February, 4, 0, 0, 0, 0, time.UTC),
		Name:        "Waleed Mannan Khan Sherwani",
		Country:     "Pakistan",
		Languages:   []string{"English", "Urdu", "Punjabi"},
		Tags:        []string{},
	}

	if id, err := db.AddUser(user.Username, user.Email, user.Password, user.DateOfBirth, user.Name, user.Country, user.Languages, user.Tags); err != nil {
		log.Fatal("Error inserting user:", err)
	} else {
		fmt.Println("Inserted user with id:", id)
	}

	if user, err := db.GetUser("WyrdWyn4"); err != nil {
		log.Fatal("Error getting user:", err)
	} else {
		fmt.Println("User:")
		fmt.Println(user)
	}

	err := db.AddUserLanguage("WyrdWyn4", []string{"English", "Arabic", "Pashto"})
	if err != nil {
		fmt.Println("Error adding languages:", err)
	}

	err = db.RemoveUserLanguage("WyrdWyn4", []string{"Pashto", "Spanish"})
	if err != nil {
		fmt.Println("Error removing language:", err)
	}

	err = db.AddUserTag("WyrdWyn4", []string{"Family", "Student"})
	if err != nil {
		fmt.Println("Error adding tags:", err)
	}

	err = db.RemoveUserTag("WyrdWyn4", []string{"Student", "Friend"})
	if err != nil {
		fmt.Println("Error removing tag:", err)
	}

	// if err := db.DeleteUser("WyrdWyn4"); err != nil {
	// 	log.Fatal("Error deleting user:", err)
	// }

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/users", db.UserHandler)

	fmt.Println("Starting server on port 8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
