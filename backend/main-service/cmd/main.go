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

	if err := db.DeleteAllTables(); err != nil {
		log.Fatal("Error deleting tables:", err)
	}

	if err := db.CreateAllTables(); err != nil {
		log.Fatal("Error creating tables:", err)
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

	post := db.Post{
		ItineraryId:  1,
		Title:        "My First Post",
		ImageLink:    "http://example.com/image.jpg",
		Description:  "This is a description of my first post.",
		CreationDate: time.Now(),
		Username:     "WyrdWyn4",
		Tags:         []string{"travel", "adventure"},
	}

	fmt.Println("Testing AddPost...")
	postId, err := db.AddPost(post)
	if err != nil {
		log.Fatalf("AddPost failed: %v", err)
	}
	fmt.Printf("Inserted post with ID: %d\n", postId)

	fmt.Println("Testing GetPost...")
	retrievedPost, err := db.GetPost(postId)
	if err != nil {
		log.Fatalf("GetPost failed: %v", err)
	}
	fmt.Printf("Retrieved Post: %+v\n", retrievedPost)

	newDescription := "Updated description for my first post."
	fmt.Println("Testing UpdatePostDescription...")
	if err := db.UpdatePostDescription(postId, newDescription); err != nil {
		log.Fatalf("UpdatePostDescription failed: %v", err)
	}

	retrievedPost, err = db.GetPost(postId)
	if err != nil {
		log.Fatalf("GetPost failed: %v", err)
	}
	fmt.Printf("Updated Post: %+v\n", retrievedPost)

	fmt.Println("Testing AddPostTag...")
	if err := db.AddPostTag(postId, "newTag"); err != nil {
		log.Fatalf("AddPostTag failed: %v", err)
	}

	retrievedPost, err = db.GetPost(postId)
	if err != nil {
		log.Fatalf("GetPost failed: %v", err)
	}
	fmt.Printf("Post after adding tag: %+v\n", retrievedPost)

	fmt.Println("Testing RemovePostTag...")
	if err := db.RemovePostTag(postId, "newTag"); err != nil {
		log.Fatalf("RemovePostTag failed: %v", err)
	}

	retrievedPost, err = db.GetPost(postId)
	if err != nil {
		log.Fatalf("GetPost failed: %v", err)
	}
	fmt.Printf("Post after removing tag: %+v\n", retrievedPost)

	fmt.Println("Testing RemovePost...")
	if err := db.RemovePost(postId); err != nil {
		log.Fatalf("RemovePost failed: %v", err)
	} else {
		fmt.Printf("Post with ID %d successfully removed.\n", postId)
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/users", db.UserHandler)

	fmt.Println("Starting server on port 8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
