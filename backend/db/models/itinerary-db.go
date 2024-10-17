package DBmodels

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
)

type Itinerary struct {
	ItineraryId  int       `json:"itineraryId"`
	Name         string    `json:"name"`
	City         string    `json:"city"`
	Country      string    `json:"country"`
	Languages    []string  `json:"languages"`
	Tags         []string  `json:"tags"`
	Events       []string  `json:"events"`
	PostId       int       `json:"postId"`
	Username     string    `json:"username"`
	CreationDate time.Time `json:"creationDate"`
	LastUpdate   time.Time `json:"lastUpdate"`
}

func CreateItineraryTable(DB *sql.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS itineraries (
		itineraryId SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		city TEXT NOT NULL,
		country TEXT NOT NULL,
		languages TEXT[],
		tags TEXT[],
		events TEXT[],
		postId INT NOT NULL,
		username VARCHAR(255) REFERENCES users(username),
		creationDate TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		lastUpdate TIMESTAMPTZ NOT NULL DEFAULT NOW()
	);`

	return CreateTable(DB, createTableSQL)
}

func AddItinerary(DB *sql.DB, itinerary Itinerary) (int, error) {
	insertItinerarySQL := `
	INSERT INTO itineraries (name, city, country, languages, tags, events, postId, username, creationDate, lastUpdate)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING itineraryId;`

	var itineraryID int
	err := DB.QueryRow(
		insertItinerarySQL, itinerary.Name, itinerary.City, itinerary.Country,
		pq.Array(itinerary.Languages), pq.Array(itinerary.Tags),
		pq.Array(itinerary.Events), itinerary.PostId,
		itinerary.Username, itinerary.CreationDate, itinerary.LastUpdate).Scan(&itineraryID)
	if err != nil {
		log.Printf("Error adding itinerary: %v\n", err)
		return 0, fmt.Errorf("failed to add itinerary: %w", err)
	}

	EventIDs, err := StringsToInts(itinerary.Events)
	if err != nil {
		log.Printf("Error converting event IDs to integers: %v\n", err)
		return 0, fmt.Errorf("failed to convert event IDs to integers: %w", err)
	}

	for _, eventID := range EventIDs {
		err = AddEventItinerary(DB, eventID, itineraryID, false)
		if err != nil {
			log.Printf("Error adding event to itinerary: %v\n", err)
			return 0, fmt.Errorf("failed to add event to itinerary: %w", err)
		}
	}

	log.Printf("Itinerary added successfully with ID: %d\n", itineraryID)
	return itineraryID, nil
}

func RemoveItinerary(DB *sql.DB, itineraryID int) error {
	var postID int
	var StringEventIDs []string

	query := `SELECT postId, events FROM itineraries WHERE itineraryId = $1;`
	err := DB.QueryRow(query, itineraryID).Scan(&postID, pq.Array(&StringEventIDs))
	if err != nil {
		log.Printf("Error retrieving PostId and EventId for itinerary ID %d: %v\n", itineraryID, err)
		return fmt.Errorf("failed to retrieve PostId and EventIds: %w", err)
	}

	eventIDs, err := StringsToInts(StringEventIDs)
	if err != nil {
		log.Printf("Error converting event IDs to integers: %v\n", err)
		return fmt.Errorf("failed to convert event IDs to integers: %w", err)
	}

	log.Println("Associated Event IDs:")
	for _, eventID := range eventIDs {
		err = RemoveEventItinerary(DB, eventID, itineraryID)
		if err != nil {
			log.Printf("Error removing event from itinerary: %v\n", err)
			return fmt.Errorf("failed to remove event from itinerary: %w", err)
		}
	}

	err = RemovePost(DB, postID)
	if err != nil {
		log.Printf("Error removing post with ID %d: %v\n", postID, err)
		return fmt.Errorf("failed to remove associated post: %w", err)
	}

	removeItinerarySQL := `DELETE FROM itineraries WHERE itineraryId = $1;`
	_, err = DB.Exec(removeItinerarySQL, itineraryID)
	if err != nil {
		log.Printf("Error removing itinerary: %v\n", err)
		return fmt.Errorf("failed to remove itinerary: %w", err)
	}

	log.Printf("Itinerary with ID %d, associated post ID %d, and related events successfully removed.\n", itineraryID, postID)
	return nil
}

func GetItinerary(DB *sql.DB, itineraryID int) (Itinerary, error) {
	getItinerarySQL := `
	SELECT itineraryId, name, city, country, languages, tags, events, postId, username, creationDate, lastUpdate
	FROM itineraries
	WHERE itineraryId = $1;`

	var itinerary Itinerary

	err := DB.QueryRow(getItinerarySQL, itineraryID).Scan(
		&itinerary.ItineraryId,
		&itinerary.Name,
		&itinerary.City,
		&itinerary.Country,
		pq.Array(&itinerary.Languages),
		pq.Array(&itinerary.Tags),
		pq.Array(&itinerary.Events),
		&itinerary.PostId,
		&itinerary.Username,
		&itinerary.CreationDate,
		&itinerary.LastUpdate,
	)
	if err == sql.ErrNoRows {
		log.Printf("No itinerary found with ID %d\n", itineraryID)
		return Itinerary{}, fmt.Errorf("no itinerary found with ID: %w", err)
	} else if err != nil {
		log.Printf("Error getting itinerary: %v\n", err)
		log.Println(err)
		return Itinerary{}, fmt.Errorf("failed to get itinerary: %w", err)
	}

	if itinerary.Languages == nil {
		itinerary.Languages = []string{}
	}

	if itinerary.Tags == nil {
		itinerary.Tags = []string{}
	}

	if itinerary.Events == nil {
		itinerary.Events = []string{}
	}

	log.Printf("Itinerary retrieved successfully: %+v\n", itinerary)
	return itinerary, nil
}

func UpdateItineraryName(DB *sql.DB, itineraryId int, name string) error {
	err := UpdateAttribute(DB, "itineraries", "itineraryId", itineraryId, "name", name)
	if err != nil {
		log.Printf("Error updating itinerary name: %v\n", err)
		return fmt.Errorf("failed to update itinerary name: %w", err)
	}

	log.Printf("Itinerary name updated successfully for ID %d.\n", itineraryId)
	return nil
}

func UpdateItineraryCity(DB *sql.DB, itineraryId int, city string) error {
	err := UpdateAttribute(DB, "itineraries", "itineraryId", itineraryId, "city", city)
	if err != nil {
		log.Printf("Error updating itinerary city: %v\n", err)
		return fmt.Errorf("failed to update itinerary city: %w", err)
	}

	log.Printf("Itinerary city updated successfully for ID %d.\n", itineraryId)
	return nil
}

func UpdateItineraryCountry(DB *sql.DB, itineraryId int, country string) error {
	err := UpdateAttribute(DB, "itineraries", "itineraryId", itineraryId, "country", country)
	if err != nil {
		log.Printf("Error updating itinerary country: %v\n", err)
		return fmt.Errorf("failed to update itinerary country: %w", err)
	}

	log.Printf("Itinerary country updated successfully for ID %d.\n", itineraryId)
	return nil
}

func AddItineraryLanguage(DB *sql.DB, itineraryId int, language string) error {
	err := AddArrayAttribute(DB, "itineraries", "itineraryId", itineraryId, "languages", []string{language})
	if err != nil {
		log.Printf("Error adding itinerary language: %v\n", err)
		return fmt.Errorf("failed to add itinerary language: %w", err)
	}

	log.Printf("Itinerary language added successfully for ID %d.\n", itineraryId)
	return nil
}

func RemoveItineraryLanguage(DB *sql.DB, itineraryId int, language string) error {
	err := RemoveArrayAttribute(DB, "itineraries", "itineraryId", itineraryId, "languages", []string{language})
	if err != nil {
		log.Printf("Error removing itinerary language: %v\n", err)
		return fmt.Errorf("failed to remove itinerary language: %w", err)
	}

	log.Printf("Itinerary language removed successfully for ID %d.\n", itineraryId)
	return nil
}

func AddItineraryTag(DB *sql.DB, itineraryId int, tag string) error {
	err := AddArrayAttribute(DB, "itineraries", "itineraryId", itineraryId, "tags", []string{tag})
	if err != nil {
		log.Printf("Error adding itinerary tag: %v\n", err)
		return fmt.Errorf("failed to add itinerary tag: %w", err)
	}

	log.Printf("Itinerary tag added successfully for ID %d.\n", itineraryId)
	return nil
}

func RemoveItineraryTag(DB *sql.DB, itineraryId int, tag string) error {
	err := RemoveArrayAttribute(DB, "itineraries", "itineraryId", itineraryId, "tags", []string{tag})
	if err != nil {
		log.Printf("Error removing itinerary tag: %v\n", err)
		return fmt.Errorf("failed to remove itinerary tag: %w", err)
	}

	log.Printf("Itinerary tag removed successfully for ID %d.\n", itineraryId)
	return nil
}

func AddItineraryEvent(DB *sql.DB, itineraryId int, eventId int, recursive bool) error {
	if !recursive {
		return nil
	}

	err := AddArrayAttribute(DB, "itineraries", "itineraryId", itineraryId, "events", IntsToStrings([]int{eventId}))
	if err != nil {
		log.Printf("Error adding itinerary event: %v\n", err)
		return fmt.Errorf("failed to add itinerary event: %w", err)
	}

	err = AddEventItinerary(DB, eventId, itineraryId, false)
	if err != nil {
		log.Printf("Error adding event to itinerary: %v\n", err)
		return fmt.Errorf("failed to add event to itinerary: %w", err)
	}

	log.Printf("Itinerary event added successfully for ID %d.\n", itineraryId)
	return nil
}

func RemoveItineraryEvent(DB *sql.DB, itineraryId int, eventId int) error {
	err := RemoveArrayAttribute(DB, "itineraries", "itineraryId", itineraryId, "events", IntsToStrings([]int{eventId}))
	if err != nil {
		log.Printf("Error removing itinerary event: %v\n", err)
		return fmt.Errorf("failed to remove itinerary event: %w", err)
	}

	err = RemoveEventItinerary(DB, eventId, itineraryId)
	if err != nil {
		log.Printf("Error removing event from itinerary: %v\n", err)
		return fmt.Errorf("failed to remove event from itinerary: %w", err)
	}

	log.Printf("Itinerary event removed successfully for ID %d.\n", itineraryId)
	return nil
}

func UpdateItineraryCreationDate(DB *sql.DB, itineraryId int, creationDate time.Time) error {
	err := UpdateAttribute(DB, "itineraries", "itineraryId", itineraryId, "creationDate", creationDate)
	if err != nil {
		log.Printf("Error updating itinerary creation date: %v\n", err)
		return fmt.Errorf("failed to update itinerary creation date: %w", err)
	}

	log.Printf("Itinerary creation date updated successfully for ID %d.\n", itineraryId)
	return nil
}

func UpdateItineraryLastUpdate(DB *sql.DB, itineraryId int, lastUpdate time.Time) error {
	err := UpdateAttribute(DB, "itineraries", "itineraryId", itineraryId, "lastUpdate", lastUpdate)
	if err != nil {
		log.Printf("Error updating itinerary last update: %v\n", err)
		return fmt.Errorf("failed to update itinerary last update: %w", err)
	}

	log.Printf("Itinerary last update updated successfully for ID %d.\n", itineraryId)
	return nil
}
