package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	db "github.com/jordyob03/TripTailor/backend/services/main-service/internal/db/models"
	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123/"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Error hashing password:", err)
	}

	user := db.User{
		Username:     "wmksherwani",
		Email:        "wmksherwani@mun.ca",
		Password:     string(hashedPassword),
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
		board := db.Board{
			Name:         "Travel Board",
			CreationDate: time.Now(),
			Description:  "A board for all things travel!",
			Username:     "wmksherwani",
			Posts:        []string{},
			Tags:         []string{"Travel", "Adventure"},
		}

		if boardId, err := db.AddBoard(DB, board); err != nil {
			log.Fatal("Error creating board:", err)
		} else {
			fmt.Printf("Board %d created successfully!\n", boardId)
			post1 := db.Post{
				CreationDate: time.Now(),
				Username:     "wmksherwani",
				Boards:       []string{},
				Likes:        0,
				Comments:     []string{},
			}

			if postid1, err := db.AddPost(DB, post1); err != nil {
				log.Fatal("Error creating post:", err)
			} else {
				fmt.Printf("Post %d created successfully!\n", postid1)
				if err := db.AddBoardPost(DB, boardId, postid1, true); err != nil {
					log.Fatal("Error adding post to board:", err)
				}
				itinerary1 := db.Itinerary{
					Name:        "Trip to the Beach",
					City:        "St. John's",
					Country:     "Canada",
					Title:       "Beach Trip",
					Description: "Enjoy a day at the beach with friends and family!",
					Price:       100.00,
					Languages:   []string{"English", "Urdu"},
					Tags:        []string{"Travel", "Adventure"},
					Events:      []string{},
					PostId:      postid1,
					Username:    "wmksherwani",
				}

				if itineraryid1, err := db.AddItinerary(DB, itinerary1); err != nil {
					log.Fatal("Error creating itinerary:", err)
				} else {
					fmt.Printf("Itinerary %d created successfully!\n", itineraryid1)

					event := db.Event{
						Name:        "Trip to the Beach",
						Cost:        100.00,
						Address:     "1234 Beach St.",
						Description: "Enjoy a day at the beach with friends and family!",
						StartTime:   time.Now(),
						EndTime:     time.Now(),
						ItineraryId: itineraryid1,
						EventImages: []string{},
					}

					if eventid, err := db.AddEvent(DB, event); err != nil {
						log.Fatal("Error creating event:", err)
					} else {
						fmt.Printf("Event %d created successfully!\n", eventid)
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

					if err := db.AddEventImage(DB, event.EventId, 1); err != nil {
						log.Fatal("Error adding event image:", err)
					}

					if err := db.AddEventImage(DB, event.EventId, 2); err != nil {
						log.Fatal("Error adding event image:", err)
					}
				}
			}

			post2 := db.Post{
				CreationDate: time.Now(),
				Username:     "wmksherwani",
				Boards:       []string{},
				Likes:        0,
				Comments:     []string{},
			}

			if postid2, err := db.AddPost(DB, post2); err != nil {
				log.Fatal("Error creating post:", err)
			} else {
				fmt.Printf("Post %d created successfully!\n", postid2)
				if err := db.AddBoardPost(DB, boardId, postid2, true); err != nil {
					log.Fatal("Error adding post to board:", err)
				}
				itinerary2 := db.Itinerary{
					Name:        "Trip to the Beach",
					City:        "St. John's",
					Country:     "Canada",
					Title:       "Beach Trip",
					Description: "Enjoy a day at the beach with friends and family!",
					Price:       100.00,
					Languages:   []string{"English", "Urdu"},
					Tags:        []string{"Travel", "Adventure"},
					Events:      []string{},
					PostId:      postid2,
					Username:    "wmksherwani",
				}

				if itineraryid2, err := db.AddItinerary(DB, itinerary2); err != nil {
					log.Fatal("Error creating itinerary:", err)
				} else {
					fmt.Printf("Itinerary %d created successfully!\n", itineraryid2)

					event := db.Event{
						Name:        "Trip to the Beach",
						Cost:        100.00,
						Address:     "1234 Beach St.",
						Description: "Enjoy a day at the beach with friends and family!",
						StartTime:   time.Now(),
						EndTime:     time.Now(),
						ItineraryId: itineraryid2,
						EventImages: []string{},
					}

					if eventid, err := db.AddEvent(DB, event); err != nil {
						log.Fatal("Error creating event:", err)
					} else {
						fmt.Printf("Event %d created successfully!\n", eventid)
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

					if err := db.AddEventImage(DB, event.EventId, 1); err != nil {
						log.Fatal("Error adding event image:", err)
					}

					if err := db.AddEventImage(DB, event.EventId, 2); err != nil {
						log.Fatal("Error adding event image:", err)
					}
				}
			}
		}
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
