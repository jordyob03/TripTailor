package DBmodels

import (
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
)

type Itinerary struct {
	ItineraryId  int       `json:"itineraryId"`
	Name         string    `json:"name"`
	Country      string    `json:"country"`
	Languages    []string  `json:"languages"`
	Tags         []string  `json:"tags"`
	Events       []int     `json:"events"`
	PostId       int       `json:"postId"`
	Username     string    `json:"username"`
	CreationDate time.Time `json:"creationDate"`
	LastUpdate   time.Time `json:"lastUpdate"`
}

func CreateItineraryTable() error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS itineraries (
		itineraryId SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		country TEXT NOT NULL,
		languages TEXT[],
		tags TEXT[],
		events INT[],
		postId INT REFERENCES posts(postId),
		username VARCHAR(255) REFERENCES users(username),
		creationDate TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		lastUpdate TIMESTAMPTZ NOT NULL DEFAULT NOW()
	);`

	return CreateTable(createTableSQL)
}

func AddItinerary(itinerary Itinerary) (int, error) {
	insertItinerarySQL := `
	INSERT INTO itineraries (name, country, languages, tags, events, postId, username)
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING itineraryId;`

	var itineraryID int
	err := DB.QueryRow(
		insertItinerarySQL, itinerary.Name, itinerary.Country,
		itinerary.Languages, itinerary.Tags,
		itinerary.Events, itinerary.PostId,
		itinerary.Username).Scan(&itineraryID)
	if err != nil {
		log.Printf("Error adding itinerary: %v\n", err)
		return 0, fmt.Errorf("failed to add itinerary: %w", err)
	}

	log.Printf("Itinerary added successfully with ID: %d\n", itineraryID)
	return itineraryID, nil
}

func RemoveItinerary(itineraryID int) error {
	var postID int
	var eventIDs []int

	// Retrieve the PostId and EventIds associated with the itinerary
	query := `SELECT postId, eventIds FROM itineraries WHERE itineraryId = $1;`
	err := DB.QueryRow(query, itineraryID).Scan(&postID, pq.Array(&eventIDs))
	if err != nil {
		log.Printf("Error retrieving PostId and EventIds for itinerary ID %d: %v\n", itineraryID, err)
		return fmt.Errorf("failed to retrieve PostId and EventIds: %w", err)
	}

	// Log the Event IDs
	log.Println("Associated Event IDs:")
	for _, eventID := range eventIDs {
		err = RemoveEventItinerary(eventID, itineraryID)
		if err != nil {
			log.Printf("Error removing event from itinerary: %v\n", err)
			return fmt.Errorf("failed to remove event from itinerary: %w", err)
		}
	}

	// Remove the associated post
	err = RemovePost(postID)
	if err != nil {
		log.Printf("Error removing post with ID %d: %v\n", postID, err)
		return fmt.Errorf("failed to remove associated post: %w", err)
	}

	// Now delete the itinerary
	removeItinerarySQL := `DELETE FROM itineraries WHERE itineraryId = $1;`
	_, err = DB.Exec(removeItinerarySQL, itineraryID)
	if err != nil {
		log.Printf("Error removing itinerary: %v\n", err)
		return fmt.Errorf("failed to remove itinerary: %w", err)
	}

	log.Printf("Itinerary with ID %d, associated post ID %d, and related events successfully removed.\n", itineraryID, postID)
	return nil
}

func GetItinerary(itineraryID int) (Itinerary, error) {
	getItinerarySQL := `
	SELECT name, country, languages, tags, events, postId, username, creationDate, lastUpdate
	FROM itineraries WHERE itineraryId = $1;`

	var itinerary Itinerary
	err := DB.QueryRow(getItinerarySQL, itineraryID).Scan(
		&itinerary.Name, &itinerary.Country, &itinerary.Languages,
		&itinerary.Tags, &itinerary.Events, &itinerary.PostId,
		&itinerary.Username, &itinerary.CreationDate, &itinerary.LastUpdate)
	if err != nil {
		log.Printf("Error getting itinerary: %v\n", err)
		return Itinerary{}, fmt.Errorf("failed to get itinerary: %w", err)
	}

	log.Printf("Itinerary retrieved successfully: %+v\n", itinerary)
	return itinerary, nil
}

func UpdateItineraryName(itineraryId int, name string) error {
	err := UpdateAttribute("itineraries", "itineraryId", itineraryId, "name", name)
	if err != nil {
		log.Printf("Error updating itinerary name: %v\n", err)
		return fmt.Errorf("failed to update itinerary name: %w", err)
	}

	log.Printf("Itinerary name updated successfully for ID %d.\n", itineraryId)
	return nil
}

func UpdateItineraryCountry(itineraryId int, country string) error {
	err := UpdateAttribute("itineraries", "itineraryId", itineraryId, "country", country)
	if err != nil {
		log.Printf("Error updating itinerary country: %v\n", err)
		return fmt.Errorf("failed to update itinerary country: %w", err)
	}

	log.Printf("Itinerary country updated successfully for ID %d.\n", itineraryId)
	return nil
}

func AddItineraryLanguage(itineraryId int, language string) error {
	err := AddArrayAttribute("itineraries", "itineraryId", itineraryId, "languages", []string{language})
	if err != nil {
		log.Printf("Error adding itinerary language: %v\n", err)
		return fmt.Errorf("failed to add itinerary language: %w", err)
	}

	log.Printf("Itinerary language added successfully for ID %d.\n", itineraryId)
	return nil
}

func RemoveItineraryLanguage(itineraryId int, language string) error {
	err := RemoveArrayAttribute("itineraries", "itineraryId", itineraryId, "languages", []string{language})
	if err != nil {
		log.Printf("Error removing itinerary language: %v\n", err)
		return fmt.Errorf("failed to remove itinerary language: %w", err)
	}

	log.Printf("Itinerary language removed successfully for ID %d.\n", itineraryId)
	return nil
}

func AddItineraryTag(itineraryId int, tag string) error {
	err := AddArrayAttribute("itineraries", "itineraryId", itineraryId, "tags", []string{tag})
	if err != nil {
		log.Printf("Error adding itinerary tag: %v\n", err)
		return fmt.Errorf("failed to add itinerary tag: %w", err)
	}

	log.Printf("Itinerary tag added successfully for ID %d.\n", itineraryId)
	return nil
}

func RemoveItineraryTag(itineraryId int, tag string) error {
	err := RemoveArrayAttribute("itineraries", "itineraryId", itineraryId, "tags", []string{tag})
	if err != nil {
		log.Printf("Error removing itinerary tag: %v\n", err)
		return fmt.Errorf("failed to remove itinerary tag: %w", err)
	}

	log.Printf("Itinerary tag removed successfully for ID %d.\n", itineraryId)
	return nil
}

func AddItineraryEvent(itineraryId int, eventId int) error {
	err := AddArrayAttribute("itineraries", "itineraryId", itineraryId, "events", []int{eventId})
	if err != nil {
		log.Printf("Error adding itinerary event: %v\n", err)
		return fmt.Errorf("failed to add itinerary event: %w", err)
	}

	err = AddEventItinerary(eventId, itineraryId)
	if err != nil {
		log.Printf("Error adding event to itinerary: %v\n", err)
		return fmt.Errorf("failed to add event to itinerary: %w", err)
	}

	log.Printf("Itinerary event added successfully for ID %d.\n", itineraryId)
	return nil
}

func RemoveItineraryEvent(itineraryId int, eventId int) error {
	err := RemoveArrayAttribute("itineraries", "itineraryId", itineraryId, "events", []int{eventId})
	if err != nil {
		log.Printf("Error removing itinerary event: %v\n", err)
		return fmt.Errorf("failed to remove itinerary event: %w", err)
	}

	err = RemoveEventItinerary(eventId, itineraryId)
	if err != nil {
		log.Printf("Error removing event from itinerary: %v\n", err)
		return fmt.Errorf("failed to remove event from itinerary: %w", err)
	}

	log.Printf("Itinerary event removed successfully for ID %d.\n", itineraryId)
	return nil
}

func UpdateItineraryCreationDate(itineraryId int, creationDate time.Time) error {
	err := UpdateAttribute("itineraries", "itineraryId", itineraryId, "creationDate", creationDate)
	if err != nil {
		log.Printf("Error updating itinerary creation date: %v\n", err)
		return fmt.Errorf("failed to update itinerary creation date: %w", err)
	}

	log.Printf("Itinerary creation date updated successfully for ID %d.\n", itineraryId)
	return nil
}

func UpdateItineraryLastUpdate(itineraryId int, lastUpdate time.Time) error {
	err := UpdateAttribute("itineraries", "itineraryId", itineraryId, "lastUpdate", lastUpdate)
	if err != nil {
		log.Printf("Error updating itinerary last update: %v\n", err)
		return fmt.Errorf("failed to update itinerary last update: %w", err)
	}

	log.Printf("Itinerary last update updated successfully for ID %d.\n", itineraryId)
	return nil
}
