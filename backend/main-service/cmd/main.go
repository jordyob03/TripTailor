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

	user := db.User{
		Username:    "WyrdWyn4",
		Email:       "wmksherwani@mun.ca",
		Password:    "password",
		DateOfBirth: time.Date(2004, time.February, 4, 0, 0, 0, 0, time.UTC),
		Name:        "Waleed Mannan Khan Sherwani",
		Country:     "Pakistan",
		Languages:   []string{"English", "Urdu", "Punjabi"},
		Tags:        []string{},
		Boards:      []string{},
		Posts:       []string{},
	}

	if id, err := db.AddUser(user); err != nil {
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

	err = db.UpdateUserEmail("WyrdWyn4", "newemail@example.com")
	if err != nil {
		log.Println("Error updating email:", err)
	} else {
		fmt.Println("Email updated successfully!")
	}

	err = db.UpdateUserPassword("WyrdWyn4", "newpassword123")
	if err != nil {
		log.Println("Error updating password:", err)
	} else {
		fmt.Println("Password updated successfully!")
	}

	newDateOfBirth := time.Date(2000, 01, 01, 0, 0, 0, 0, time.UTC)
	err = db.UpdateUserDateOfBirth("WyrdWyn4", newDateOfBirth)
	if err != nil {
		log.Println("Error updating date of birth:", err)
	} else {
		fmt.Println("Date of birth updated successfully!")
	}

	err = db.UpdateUserCountry("WyrdWyn4", "Canada")
	if err != nil {
		log.Println("Error updating country:", err)
	} else {
		fmt.Println("Country updated successfully!")
	}

	err = db.AddUserBoard("WyrdWyn4", 1)
	if err != nil {
		log.Println("Error adding board:", err)
	} else {
		fmt.Println("Board added successfully!")
	}

	err = db.AddUserPost("WyrdWyn4", 2)
	if err != nil {
		log.Println("Error adding post:", err)
	} else {
		fmt.Println("Post added successfully!")
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/users", db.UserHandler)

	fmt.Println("Starting server on port 8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
