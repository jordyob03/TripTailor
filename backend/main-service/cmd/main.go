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

	board := db.Board{
		Name:        "Sample Board",
		Description: "This is a test board",
		Username:    "WyrdWyn4",
		Posts:       []string{},
		Tags:        []string{"travel", "fun"},
	}

	fmt.Println("Testing AddBoard...")
	err := db.AddBoard(board)
	if err != nil {
		log.Fatalf("AddBoard failed: %v", err)
	}

	boardID := 1

	fmt.Println("Testing GetBoard...")
	retrievedBoard, err := db.GetBoard(boardID)
	if err != nil {
		log.Fatalf("GetBoard failed: %v", err)
	}
	fmt.Printf("Retrieved Board: %+v\n", retrievedBoard)

	fmt.Println("Testing UpdateBoardName...")
	err = db.UpdateBoardName(boardID, "New Board Name")
	if err != nil {
		log.Fatalf("UpdateBoardName failed: %v", err)
	}

	fmt.Println("Testing UpdateBoardDescription...")
	err = db.UpdateBoardDescription(boardID, "New Board Description")
	if err != nil {
		log.Fatalf("UpdateBoardDescription failed: %v", err)
	}

	fmt.Println("Testing UpdateBoardCreationDate...")
	err = db.UpdateBoardCreationDate(boardID, time.Now())
	if err != nil {
		log.Fatalf("UpdateBoardCreationDate failed: %v", err)
	}

	fmt.Println("Testing UpdateBoardDescription...")
	err = db.UpdateBoardDescription(boardID, "New Board Description")
	if err != nil {
		log.Fatalf("UpdateBoardDescription failed: %v", err)
	}

	fmt.Println("Testing AddBoardTag...")
	err = db.AddBoardTag(fmt.Sprint(boardID), "newTag")
	if err != nil {
		log.Fatalf("AddBoardTag failed: %v", err)
	}

	fmt.Println("Testing RemoveBoardTag...")
	err = db.RemoveBoardTag(fmt.Sprint(boardID), "newTag")
	if err != nil {
		log.Fatalf("RemoveBoardTag failed: %v", err)
	}

	fmt.Println("Testing AddBoardPost...")
	err = db.AddBoardPost(fmt.Sprint(boardID), 123)
	if err != nil {
		log.Fatalf("AddBoardPost failed: %v", err)
	}

	fmt.Println("Testing RemoveBoardPost...")
	err = db.RemoveBoardPost(fmt.Sprint(boardID), 123)
	if err != nil {
		log.Fatalf("RemoveBoardPost failed: %v", err)
	}

	fmt.Println("Testing RemoveBoard...")
	err = db.RemoveBoard(retrievedBoard)
	if err != nil {
		log.Fatalf("RemoveBoard failed: %v", err)
	}

	fmt.Println("All tests completed successfully.")

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/users", db.UserHandler)

	fmt.Println("Starting server on port 8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
