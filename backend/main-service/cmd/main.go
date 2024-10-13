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
		Boards:      db.IntsToStrings([]int{}),
		Posts:       db.IntsToStrings([]int{}),
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
		Posts:       db.IntsToStrings([]int{1}),
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

	post := db.Post{
		PostId:       1,
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

	newItinerary := db.Itinerary{
		Name:         "Trip to Japan",
		Country:      "Japan",
		Languages:    []string{"Japanese", "English"},
		Tags:         []string{"travel", "culture"},
		Events:       db.IntsToStrings([]int{}),
		PostId:       1,
		Username:     "WyrdWyn4",
		CreationDate: time.Now(),
		LastUpdate:   time.Now(),
	}

	itineraryID, err := db.AddItinerary(newItinerary)
	if err != nil {
		log.Fatalf("Error adding itinerary: %v", err)
	}
	fmt.Printf("Itinerary added with ID: %d\n", itineraryID)

	itinerary, err := db.GetItinerary(itineraryID)
	if err != nil {
		log.Fatalf("Error getting itinerary: %v", err)
	}
	fmt.Printf("Retrieved itinerary: %+v\n", itinerary)

	err = db.UpdateItineraryName(itineraryID, "Trip to Japan (Updated)")
	if err != nil {
		log.Fatalf("Error updating itinerary name: %v", err)
	}

	err = db.UpdateItineraryCountry(itineraryID, "Japan (Updated)")
	if err != nil {
		log.Fatalf("Error updating itinerary country: %v", err)
	}

	err = db.AddItineraryLanguage(itineraryID, "French")
	if err != nil {
		log.Fatalf("Error adding itinerary language: %v", err)
	}

	err = db.RemoveItineraryLanguage(itineraryID, "English")
	if err != nil {
		log.Fatalf("Error removing itinerary language: %v", err)
	}

	err = db.AddItineraryTag(itineraryID, "Adventure")
	if err != nil {
		log.Fatalf("Error adding itinerary tag: %v", err)
	}

	err = db.RemoveItineraryTag(itineraryID, "travel")
	if err != nil {
		log.Fatalf("Error removing itinerary tag: %v", err)
	}

	err = db.AddItineraryEvent(itineraryID, 3)
	if err != nil {
		log.Fatalf("Error adding itinerary event: %v", err)
	}

	err = db.RemoveItineraryEvent(itineraryID, 1)
	if err != nil {
		log.Fatalf("Error removing itinerary event: %v", err)
	}

	// err = db.RemoveItinerary(itineraryID)
	// if err != nil {
	// 	log.Fatalf("Error removing itinerary: %v", err)
	// }

	fmt.Println("All operations completed successfully!")

	event := db.Event{
		Name:         "Sample Event",
		Price:        100,
		Location:     "Sample Location",
		Description:  "This is a sample event.",
		StartDate:    time.Now(),
		EndDate:      time.Now().Add(24 * time.Hour),
		ItineraryIds: db.IntsToStrings([]int{}),
		PhotoLinks:   []string{"http://example.com/photo.jpg"},
	}

	eventID, err := db.AddEvent(event)
	if err != nil {
		log.Fatalf("Failed to add event: %v\n", err)
	}

	retrievedEvent, err := db.GetEvent(eventID)
	if err != nil {
		log.Fatalf("Failed to retrieve event: %v\n", err)
	}
	fmt.Printf("Retrieved Event: %+v\n", retrievedEvent)

	err = db.UpdateEventName(eventID, "Updated Sample Event")
	if err != nil {
		log.Fatalf("Failed to update event name: %v\n", err)
	}

	err = db.UpdateEventPrice(eventID, 150)
	if err != nil {
		log.Fatalf("Failed to update event price: %v\n", err)
	}

	err = db.AddEventPhotoLink(eventID, "http://example.com/newphoto.jpg")
	if err != nil {
		log.Fatalf("Failed to add photo link: %v\n", err)
	}

	err = db.RemoveEventPhotoLink(eventID, "http://example.com/photo.jpg")
	if err != nil {
		log.Fatalf("Failed to remove photo link: %v\n", err)
	}

	err = db.AddEventItinerary(eventID, 1) // Assuming itinerary ID 1 exists
	if err != nil {
		log.Fatalf("Failed to add itinerary: %v\n", err)
	}

	err = db.RemoveEventItinerary(eventID, 1) // Assuming itinerary ID 1 exists
	if err != nil {
		log.Fatalf("Failed to remove itinerary: %v\n", err)
	}

	err = db.RemoveEvent(eventID)
	if err != nil {
		log.Fatalf("Failed to remove event: %v\n", err)
	}

	fmt.Printf("Event with ID %d removed successfully.\n", eventID)

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/users", db.UserHandler)

	fmt.Println("Starting server on port 8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
