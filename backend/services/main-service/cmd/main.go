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

		// Board 1 - "Travel Board"
		board1 := db.Board{
			Name:         "Travel Board",
			CreationDate: time.Now(),
			Description:  "A board for all things travel!",
			Username:     "wmksherwani",
			Posts:        []string{},
			Tags:         []string{"Travel", "Adventure"},
		}
		if boardId1, err := db.AddBoard(DB, board1); err != nil {
			log.Fatal("Error creating board:", err)
		} else {
			fmt.Printf("Board %d created successfully!\n", boardId1)

			// Post 1
			post1 := db.Post{
				CreationDate: time.Now(),
				Username:     "wmksherwani",
				Boards:       []string{"Travel Board"},
				Likes:        0,
				Comments:     []string{},
			}
			if postid1, err := db.AddPost(DB, post1); err != nil {
				log.Fatal("Error creating post:", err)
			} else {
				fmt.Printf("Post %d created successfully!\n", postid1)
				if err := db.AddBoardPost(DB, boardId1, postid1, true); err != nil {
					log.Fatal("Error adding post to board:", err)
				}

				// Itinerary 1
				itinerary1 := db.Itinerary{
					Name:        "Beach Vacation",
					City:        "St. John's",
					Country:     "Canada",
					Title:       "Beach Trip",
					Description: "Enjoy a relaxing day at the beach with friends and family!",
					Price:       150.00,
					Languages:   []string{"English", "Arabic"},
					Tags:        []string{"Travel", "Vacation"},
					Events:      []string{},
					PostId:      postid1,
					Username:    "wmksherwani",
				}
				if itineraryid1, err := db.AddItinerary(DB, itinerary1); err != nil {
					log.Fatal("Error creating itinerary:", err)
				} else {
					fmt.Printf("Itinerary %d created successfully!\n", itineraryid1)

					// Event 1
					event1 := db.Event{
						Name:        "Beach Party",
						Cost:        150.00,
						Address:     "123 Beach St.",
						Description: "A fun beach party with music, games, and snacks!",
						StartTime:   time.Now(),
						EndTime:     time.Now().Add(time.Hour * 4),
						ItineraryId: itineraryid1,
						EventImages: []string{},
					}
					if eventid1, err := db.AddEvent(DB, event1); err != nil {
						log.Fatal("Error creating event:", err)
					} else {
						fmt.Printf("Event %d created successfully!\n", eventid1)
						// Image for Event 1
						image1 := db.Image{
							ImageData: db.WebImageToByte("https://wallpapercave.com/wp/NjGW245.jpg"),
						}
						id1, err := db.AddImage(DB, image1)
						if err != nil {
							log.Fatal("Error adding image:", err)
						} else {
							fmt.Printf("Image %d added successfully!\n", id1)
						}

						// Add image to event
						if err := db.AddEventImage(DB, eventid1, id1); err != nil {
							log.Fatal("Error adding event image:", err)
						}
					}
				}
			}

			// Post 2
			post2 := db.Post{
				CreationDate: time.Now(),
				Username:     "wmksherwani",
				Boards:       []string{"Travel Board"},
				Likes:        0,
				Comments:     []string{},
			}
			if postid2, err := db.AddPost(DB, post2); err != nil {
				log.Fatal("Error creating post:", err)
			} else {
				fmt.Printf("Post %d created successfully!\n", postid2)
				if err := db.AddBoardPost(DB, boardId1, postid2, true); err != nil {
					log.Fatal("Error adding post to board:", err)
				}

				// Itinerary 2
				itinerary2 := db.Itinerary{
					Name:        "Mountain Hike",
					City:        "Banff",
					Country:     "Canada",
					Title:       "Mountain Adventure",
					Description: "A challenging but rewarding hike to the top of the mountain!",
					Price:       200.00,
					Languages:   []string{"English", "Punjabi"},
					Tags:        []string{"Adventure", "Hiking"},
					Events:      []string{},
					PostId:      postid2,
					Username:    "wmksherwani",
				}
				if itineraryid2, err := db.AddItinerary(DB, itinerary2); err != nil {
					log.Fatal("Error creating itinerary:", err)
				} else {
					fmt.Printf("Itinerary %d created successfully!\n", itineraryid2)

					// Event 2
					event2 := db.Event{
						Name:        "Mountain Hike",
						Cost:        200.00,
						Address:     "Banff National Park",
						Description: "A guided hike through Banff's beautiful trails.",
						StartTime:   time.Now(),
						EndTime:     time.Now().Add(time.Hour * 6),
						ItineraryId: itineraryid2,
						EventImages: []string{},
					}
					if eventid2, err := db.AddEvent(DB, event2); err != nil {
						log.Fatal("Error creating event:", err)
					} else {
						fmt.Printf("Event %d created successfully!\n", eventid2)
						// Image for Event 2
						image2 := db.Image{
							ImageData: db.WebImageToByte("https://wallpapercave.com/wp/wp6976401.png"),
						}
						id2, err := db.AddImage(DB, image2)
						if err != nil {
							log.Fatal("Error adding image:", err)
						} else {
							fmt.Printf("Image %d added successfully!\n", id2)
						}

						// Add image to event
						if err := db.AddEventImage(DB, eventid2, id2); err != nil {
							log.Fatal("Error adding event image:", err)
						}
					}
				}
			}
		}

		// Board 2 - "Adventure Board"
		board2 := db.Board{
			Name:         "Adventure Board",
			CreationDate: time.Now(),
			Description:  "A board for adventurous activities!",
			Username:     "wmksherwani",
			Posts:        []string{},
			Tags:         []string{"Adventure", "Outdoor"},
		}
		if boardId2, err := db.AddBoard(DB, board2); err != nil {
			log.Fatal("Error creating board:", err)
		} else {
			fmt.Printf("Board %d created successfully!\n", boardId2)

			// Post 3
			post3 := db.Post{
				CreationDate: time.Now(),
				Username:     "wmksherwani",
				Boards:       []string{"Adventure Board"},
				Likes:        0,
				Comments:     []string{},
			}
			if postid3, err := db.AddPost(DB, post3); err != nil {
				log.Fatal("Error creating post:", err)
			} else {
				fmt.Printf("Post %d created successfully!\n", postid3)
				if err := db.AddBoardPost(DB, boardId2, postid3, true); err != nil {
					log.Fatal("Error adding post to board:", err)
				}

				// Itinerary 3
				itinerary3 := db.Itinerary{
					Name:        "Scuba Diving",
					City:        "Vancouver",
					Country:     "Canada",
					Title:       "Ocean Adventure",
					Description: "Explore the underwater world with a guided scuba diving tour.",
					Price:       300.00,
					Languages:   []string{"English", "Sindhi"},
					Tags:        []string{"Adventure", "Scuba Diving"},
					Events:      []string{},
					PostId:      postid3,
					Username:    "wmksherwani",
				}
				if itineraryid3, err := db.AddItinerary(DB, itinerary3); err != nil {
					log.Fatal("Error creating itinerary:", err)
				} else {
					fmt.Printf("Itinerary %d created successfully!\n", itineraryid3)

					// Event 3
					event3 := db.Event{
						Name:        "Scuba Diving Tour",
						Cost:        300.00,
						Address:     "Vancouver Aquarium",
						Description: "A thrilling scuba diving adventure to explore marine life.",
						StartTime:   time.Now(),
						EndTime:     time.Now().Add(time.Hour * 5),
						ItineraryId: itineraryid3,
						EventImages: []string{},
					}
					if eventid3, err := db.AddEvent(DB, event3); err != nil {
						log.Fatal("Error creating event:", err)
					} else {
						fmt.Printf("Event %d created successfully!\n", eventid3)
						// Image for Event 3
						image3 := db.Image{
							ImageData: db.WebImageToByte("https://wallpapercave.com/wp/wp8484607.jpg"),
						}
						id3, err := db.AddImage(DB, image3)
						if err != nil {
							log.Fatal("Error adding image:", err)
						} else {
							fmt.Printf("Image %d added successfully!\n", id3)
						}

						// Add image to event
						if err := db.AddEventImage(DB, eventid3, id3); err != nil {
							log.Fatal("Error adding event image:", err)
						}
					}
				}
			}

			// Post 4
			post4 := db.Post{
				CreationDate: time.Now(),
				Username:     "wmksherwani",
				Boards:       []string{"Adventure Board"},
				Likes:        0,
				Comments:     []string{},
			}

			if postid4, err := db.AddPost(DB, post4); err != nil {
				log.Fatal("Error creating post:", err)
			} else {
				fmt.Printf("Post %d created successfully!\n", postid4)
				if err := db.AddBoardPost(DB, boardId2, postid4, true); err != nil {
					log.Fatal("Error adding post to board:", err)
				}

				// Itinerary 4
				itinerary4 := db.Itinerary{
					Name:        "Skydiving",
					City:        "Toronto",
					Country:     "Canada",
					Title:       "Skydiving Adventure",
					Description: "Experience the thrill of skydiving with a tandem jump!",
					Price:       400.00,
					Languages:   []string{"English", "Pashto"},
					Tags:        []string{"Adventure", "Skydiving"},
					Events:      []string{},
					PostId:      postid4,
					Username:    "wmksherwani",
				}
				if itineraryid4, err := db.AddItinerary(DB, itinerary4); err != nil {
					log.Fatal("Error creating itinerary:", err)
				} else {
					fmt.Printf("Itinerary %d created successfully!\n", itineraryid4)

					// Event 4
					event4 := db.Event{
						Name:        "Tandem Skydiving",
						Cost:        400.00,
						Address:     "Toronto Skydiving Center",
						Description: "A thrilling tandem skydiving experience with a professional instructor.",
						StartTime:   time.Now(),
						EndTime:     time.Now().Add(time.Hour * 3),
						ItineraryId: itineraryid4,
						EventImages: []string{},
					}
					if eventid4, err := db.AddEvent(DB, event4); err != nil {
						log.Fatal("Error creating event:", err)
					} else {
						fmt.Printf("Event %d created successfully!\n", eventid4)
						// Image for Event 4
						image4 := db.Image{
							ImageData: db.WebImageToByte("https://wallpapercave.com/wp/wp8484603.jpg"),
						}
						id4, err := db.AddImage(DB, image4)
						if err != nil {
							log.Fatal("Error adding image:", err)
						} else {
							fmt.Printf("Image %d added successfully!\n", id4)
						}
					}
				}
			}

			// Post 5
			post5 := db.Post{
				CreationDate: time.Now(),
				Username:     "wmksherwani",
				Boards:       []string{"Adventure Board"},
				Likes:        0,
				Comments:     []string{},
			}

			if postid5, err := db.AddPost(DB, post5); err != nil {
				log.Fatal("Error creating post:", err)
			} else {
				fmt.Printf("Post %d created successfully!\n", postid5)
				if err := db.AddBoardPost(DB, boardId2, postid5, true); err != nil {
					log.Fatal("Error adding post to board:", err)
				}

				// Itinerary 5
				itinerary5 := db.Itinerary{
					Name:        "Bungee Jumping",
					City:        "Niagara Falls",
					Country:     "Canada",
					Title:       "Bungee Jumping Adventure",
					Description: "Take the plunge with a bungee jump off the Niagara Falls!",
					Price:       500.00,
					Languages:   []string{"English", "Balochi"},
					Tags:        []string{"Adventure", "Bungee Jumping"},
					Events:      []string{},
					PostId:      postid5,
					Username:    "wmksherwani",
				}
				if itineraryid5, err := db.AddItinerary(DB, itinerary5); err != nil {
					log.Fatal("Error creating itinerary:", err)
				} else {
					fmt.Printf("Itinerary %d created successfully!\n", itineraryid5)

					// Event 5
					event5 := db.Event{
						Name:        "Bungee Jump",
						Cost:        500.00,
						Address:     "Niagara Falls Bungee Jumping Center",
						Description: "Experience the ultimate thrill with a bungee jump off the Niagara Falls!",
						StartTime:   time.Now(),
						EndTime:     time.Now().Add(time.Hour * 2),
						ItineraryId: itineraryid5,
						EventImages: []string{},
					}
					if eventid5, err := db.AddEvent(DB, event5); err != nil {
						log.Fatal("Error creating event:", err)
					} else {
						fmt.Printf("Event %d created successfully!\n", eventid5)
						// Image for Event 5
						image5 := db.Image{
							ImageData: db.WebImageToByte("https://wallpapercave.com/wp/wp8484601.jpg"),
						}
						id5, err := db.AddImage(DB, image5)
						if err != nil {
							log.Fatal("Error adding image:", err)
						} else {
							fmt.Printf("Image %d added successfully!\n", id5)
						}
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
