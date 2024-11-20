package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	db "github.com/jordyob03/TripTailor/backend/services/main-service/internal/db/models"
)

func PackImagesFromLocal(dirPath string, DB *sql.DB) (image_ids []int, count int, err error) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal("Error reading directory:", err)
		return nil, 0, err
	}

	// Sort files numerically based on the numeric part of the filename
	sort.Slice(files, func(i, j int) bool {
		num1 := extractNumber(files[i].Name())
		num2 := extractNumber(files[j].Name())
		return num1 < num2
	})

	image_ids = make([]int, 0, len(files))

	for _, file := range files {
		filePath := fmt.Sprintf("%s/%s", dirPath, file.Name())

		imageData := db.ImageToByte(filePath)

		if len(imageData) == 0 {
			log.Printf("Failed to convert image to bytes for file %s: empty byte data", file.Name())
			continue
		}

		image := db.Image{
			ImageData: imageData,
			Metadata:  []string{"local"},
			ImageId:   0,
		}

		imageID, err := db.AddImage(DB, image)
		if err != nil {
			log.Printf("Error adding image for file %s: %v", file.Name(), err)
			continue
		}

		image_ids = append(image_ids, imageID)
		count++
	}

	return image_ids, count, nil
}

// Helper function to extract the numeric portion from a filename
func extractNumber(filename string) int {
	nameWithoutExt := strings.TrimSuffix(filename, ".png")
	num, err := strconv.Atoi(nameWithoutExt)
	if err != nil {
		log.Printf("Error extracting number from filename %s: %v", filename, err)
		return -1
	}
	return num
}

func PackUsersFromJSON(fp string, DB *sql.DB) (user_ids []int, count int) {
	data, err := os.ReadFile(fp)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
	var users []db.User
	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
	}
	user_ids_list := make([]int, len(users))

	for i, user := range users {
		if user_id, err := db.AddUser(DB, users[i]); err != nil {
			fmt.Printf("Failed to insert user %s: %v\n", user.Username, err)
		} else {
			user_ids_list[i] = user_id
			fmt.Printf("Successfully inserted user: %s with ID: %d\n", user.Username, user_id)
		}
	}
	return user_ids, count
}
func PackEventFromJSON(fp string, DB *sql.DB) (event_ids []int, count int) {
	data, err := os.ReadFile(fp)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil, 0
	}

	var events []db.Event
	err = json.Unmarshal(data, &events)
	if err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
	}

	event_ids_list := make([]int, len(events))

	for i, event := range events {
		if event_id, err := db.AddEvent(DB, events[i]); err != nil {
			fmt.Printf("Failed to insert event %s: %v\n", event.Name, err)
		} else {
			event_ids_list[i] = event_id
			fmt.Printf("Successfully inserted event: %s with ID: %d\n", event.Name, event_id)
		}
	}

	return event_ids_list, len(events)
}
func PackItinsAndPostFromJSON(fp string, DB *sql.DB) (itins_ids []int, count int) {
	var itins []db.Itinerary
	data, err := os.ReadFile(fp)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil, 0
	}
	err = json.Unmarshal(data, &itins)
	if err != nil {
		fmt.Printf("Failed to decode JSON: %v\n", err)
		return nil, 0
	}

	itins_ids_list := make([]int, len(itins))

	for i, itin := range itins {
		if itin_id, err := db.AddItinerary(DB, itins[i]); err != nil {
			fmt.Printf("Failed to insert itinerary %s: %v\n", itin.Title, err)
		} else {
			post := db.Post{
				ItineraryId:  itin_id,
				CreationDate: time.Now(),
				Username:     itin.Username,
				Boards:       []string{},
				Likes:        0,
				Comments:     []string{},
			}

			if post_id, err := db.AddPost(DB, post); err != nil {
				fmt.Printf("Failed to insert post for itinerary %s: %v\n", itin.Title, err)
			} else {
				itins_ids_list[i] = itin_id
				fmt.Printf("Successfully inserted itinerary: %s with ID: %d and post ID: %d\n", itin.Title, itin_id, post_id)
			}
		}
	}

	return itins_ids_list, len(itins)
}
func PackBoardsFromJSON(fp string, DB *sql.DB) (board_ids []int, count int) {
	var boards []db.Board
	data, err := os.ReadFile(fp)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil, 0
	}
	err = json.Unmarshal(data, &boards)
	if err != nil {
		fmt.Printf("Failed to decode JSON: %v\n", err)
		return nil, 0
	}
	board_ids = make([]int, len(boards))
	for i, board := range boards {
		if board_id, err := db.AddBoard(DB, boards[i]); err != nil {
			fmt.Printf("Failed to insert board %s: %v\n", board.Name, err)
		} else {
			board_ids[i] = board_id
			fmt.Printf("Successfully inserted board: %s with ID: %d\n", board.Name, board_id)
		}
	}
	return board_ids, len(boards)
}
