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
		ProfileImage: db.WebImageToByte("https://media.licdn.com/dms/image/v2/D4D03AQHQVU25rQFoig/profile-displayphoto-shrink_200_200/profile-displayphoto-shrink_200_200/0/1682320154670?e=2147483647&v=beta&t=9QPPZBTsm9pJU7J2FP4aV1ZhJXHUXMFOO0JwAFxsDCU"),
		CoverImage:   db.WebImageToByte("https://media.licdn.com/dms/image/v2/D4E16AQH1arQjQSnB0g/profile-displaybackgroundimage-shrink_200_800/profile-displaybackgroundimage-shrink_200_800/0/1667348768354?e=2147483647&v=beta&t=ya0kKsPaqstxPqwGwe7P2YBmHVOsqy7rPujPS3Meje0"),
	}

	if userid, err := db.AddUser(DB, user); err != nil {
		log.Fatal("Error creating user:", err)
	} else {
		fmt.Printf("User %d created successfully!", userid)
	}

	if err := db.SaveProfileImage(DB, user.Username); err != nil {
		log.Fatal("Error saving user images:", err)
	}

	if err := db.SaveCoverImage(DB, user.Username); err != nil {
		log.Fatal("Error saving user images:", err)
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
		fmt.Printf("Post %d created successfully!", postid)
	}

	if err := db.AddPostImage(DB, post.PostId, "https://th.bing.com/th/id/OIP.EER8U-yHIt1XjHtBjVOYtAHaER?rs=1&pid=ImgDetMain"); err != nil {
		log.Fatal("Error adding post image:", err)
	}

	if err := db.AddPostImage(DB, post.PostId, "https://th.bing.com/th/id/R.07f2e310f8f5adc67f6d2789b953531e?rik=zL03KcOzeSjUZQ&pid=ImgRaw&r=0"); err != nil {
		log.Fatal("Error adding post image:", err)
	}

	post, err = db.GetAllPostImages(DB, post.PostId)
	if err != nil {
		log.Fatal("Error getting post images:", err)
	}
	for i, postImage := range db.post.PostImages {
		fmt.Printf("Post Image %d: %s\n", i, postImage)
		db.ByteToImage(postImage, fmt.Sprintf("post%dimage%d.jpg", post.PostId, i))
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/users", db.UserHandlerWrapper(DB))

	fmt.Println("Starting server on port 8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
