package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	db "github.com/jordyob03/TripTailor/backend/services/main-service/internal/db/models"
	"log"
	"os"
)

func PackImagesFromLocal(dirPath string, DB *sql.DB) (image_ids []int, count int, err error) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal("Error reading directory:", err)
		return nil, 0, err
	}

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
func PackItinsFromJSON(fp string, DB *sql.DB) (itins_ids []int, count int) {
	var itins []db.Itinerary
	data, err := os.ReadFile(fp)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	err = json.Unmarshal(data, &itins)
	if err != nil {
		fmt.Printf("Failed to decode JSON: %v\n", err)
	}
	itins_ids_list := make([]int, len(itins))

	for i, itin := range itins {
		if itin_id, err := db.AddItinerary(DB, itins[i]); err != nil {
			fmt.Printf("Failed to insert itinerary %s: %v\n", itin.Title, err)
		} else {
			itins_ids_list[i] = itin_id
			fmt.Printf("Successfully inserted itinerary: %s with ID: %d\n", itin.Title, itin_id)
		}
	}
	return itins_ids_list, len(itins)

}
