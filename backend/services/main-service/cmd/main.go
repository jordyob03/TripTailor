package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	db "github.com/jordyob03/TripTailor/backend/services/main-service/internal/db/models"
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

	user := db.User{
		Username:     "wmksherwani",
		Email:        "wmksherwani@mun.ca",
		Password:     "password",
		DateOfBirth:  time.Now(),
		Name:         "Waleed Sherwani",
		Country:      "Canada",
		Languages:    []string{"English", "Urdu"},
		Tags:         []string{"Travel", "Adventure"},
		Boards:       []string{},
		Posts:        []string{},
		ProfileImage: 0,
		CoverImage:   0,
	}

	if userid, err := db.AddUser(DB, user); err != nil {
		log.Fatal("Error creating user:", err)
	} else {
		fmt.Printf("User %d created successfully!\n", userid)
	}

	post := db.Post{
		PostId:       0,
		ItineraryId:  0,
		Title:        "My First Post",
		Description:  "This is my first post on TripTailor!",
		CreationDate: time.Now(),
		Username:     "wmksherwani",
		Tags:         []string{"Travel", "Adventure"},
		Boards:       []string{},
		PostImages:   []string{},
	}

	if postid, err := db.AddPost(DB, post); err != nil {
		log.Fatal("Error creating post:", err)
	} else {
		fmt.Printf("Post %d created successfully!\n", postid)
	}

	image1 := db.Image{
		ImageData: db.WebImageToByte("https://wallpapercave.com/wp/wp8484597.jpg"),
	}

	id, err := db.AddImage(DB, image1)
	if err != nil {
		log.Fatal("Error adding image:", err)
	} else {
		fmt.Printf("Image %d added successfully!\n", id)
	}

	image2 := db.Image{
		ImageData: db.WebImageToByte("https://wallpapercave.com/uwp/uwp4469044.png"),
	}

	id, err = db.AddImage(DB, image2)
	if err != nil {
		log.Fatal("Error adding image:", err)
	} else {
		fmt.Printf("Image %d added successfully!", id)
	}

	if err := db.AddPostImage(DB, post.PostId, 1); err != nil {
		log.Fatal("Error adding post image:", err)
	}

	if err := db.AddPostImage(DB, post.PostId, 1); err != nil {
		log.Fatal("Error adding post image:", err)
	}

	http.HandleFunc("/images/", db.ImageHandler(DB))
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/users", db.UserHandlerWrapper(DB))

	fmt.Println("Starting server on port 8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
