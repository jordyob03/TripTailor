package DBmodels

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
)

type Itinerary struct {
	ItineraryId int      `json:"itineraryId"`
	City        string   `json:"city"`
	Country     string   `json:"country"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Languages   []string `json:"languages"`
	Tags        []string `json:"tags"`
	Events      []string `json:"events"`
	PostId      int      `json:"postId"`
	Username    string   `json:"username"`
}

func CreateItineraryTable(DB *sql.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS itineraries (
		itineraryId SERIAL PRIMARY KEY,
		city TEXT NOT NULL,
		country TEXT NOT NULL,
		title TEXT,
		description TEXT,
		price FLOAT,
		languages TEXT[],
		tags TEXT[],
		events TEXT[],
		postId INT NOT NULL,
		username VARCHAR(255) REFERENCES users(username)
	);`

	return CreateTable(DB, createTableSQL)
}

func AddItinerary(DB *sql.DB, itinerary Itinerary) (int, error) {
	insertItinerarySQL := `
	INSERT INTO itineraries (city, country, title, description, price, languages, tags, events, postId, username)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING itineraryId;`

	if itinerary.Languages == nil {
		itinerary.Languages = []string{}
	}

	if itinerary.Tags == nil {
		itinerary.Tags = []string{}
	}

	if itinerary.Events == nil {
		itinerary.Events = []string{}
	}

	var itineraryID int
	err := DB.QueryRow(
		insertItinerarySQL, itinerary.City, itinerary.Country,
		itinerary.Title, itinerary.Description, itinerary.Price,
		pq.Array(itinerary.Languages), pq.Array(itinerary.Tags),
		pq.Array(itinerary.Events), itinerary.PostId,
		itinerary.Username).Scan(&itineraryID)
	if err != nil {
		log.Printf("Error adding itinerary: %v\n", err)
		return 0, fmt.Errorf("failed to add itinerary: %w", err)
	}

	err = UpdatePostItineraryId(DB, itinerary.PostId, itineraryID)
	if err != nil {
		log.Printf("Error updating post with itinerary ID: %v\n", err)
		return 0, fmt.Errorf("failed to update post with itinerary ID: %w", err)
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
		err = RemoveEvent(DB, eventID)
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
	SELECT *
	FROM itineraries
	WHERE itineraryId = $1;`

	var itinerary Itinerary

	err := DB.QueryRow(getItinerarySQL, itineraryID).Scan(
		&itinerary.ItineraryId,
		&itinerary.City,
		&itinerary.Country,
		&itinerary.Title,
		&itinerary.Description,
		&itinerary.Price,
		pq.Array(&itinerary.Languages),
		pq.Array(&itinerary.Tags),
		pq.Array(&itinerary.Events),
		&itinerary.PostId,
		&itinerary.Username,
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

func UpdateItineraryCity(DB *sql.DB, itineraryId int, city string) error {
	return UpdateAttribute(DB, "itineraries", "itineraryId", itineraryId, "city", city)
}

func UpdateItineraryCountry(DB *sql.DB, itineraryId int, country string) error {
	return UpdateAttribute(DB, "itineraries", "itineraryId", itineraryId, "country", country)
}

func UpdateItineraryTitle(DB *sql.DB, itineraryId int, title string) error {
	return UpdateAttribute(DB, "itineraries", "itineraryId", itineraryId, "title", title)
}

func AddItineraryLanguage(DB *sql.DB, itineraryId int, language string) error {
	return AddArrayAttribute(DB, "itineraries", "itineraryId", itineraryId, "languages", []string{language})
}

func RemoveItineraryLanguage(DB *sql.DB, itineraryId int, language string) error {
	return RemoveArrayAttribute(DB, "itineraries", "itineraryId", itineraryId, "languages", []string{language})
}

func AddItineraryTag(DB *sql.DB, itineraryId int, tag string) error {
	return AddArrayAttribute(DB, "itineraries", "itineraryId", itineraryId, "tags", []string{tag})
}

func RemoveItineraryTag(DB *sql.DB, itineraryId int, tag string) error {
	return RemoveArrayAttribute(DB, "itineraries", "itineraryId", itineraryId, "tags", []string{tag})
}

func AddItineraryEvent(DB *sql.DB, itineraryId int, eventId int) error {
	return AddArrayAttribute(DB, "itineraries", "itineraryId", itineraryId, "events", IntsToStrings([]int{eventId}))
}

func RemoveItineraryEvent(DB *sql.DB, itineraryId int, eventId int) error {
	return RemoveArrayAttribute(DB, "itineraries", "itineraryId", itineraryId, "events", IntsToStrings([]int{eventId}))
}

func UpdateItineraryPrice(DB *sql.DB, itineraryId int) error {
	var price = 0.0
	itinerary, err := GetItinerary(DB, itineraryId)
	if err != nil {
		log.Printf("Error getting events for itinerary ID %d: %v\n", itineraryId, err)
		return fmt.Errorf("failed to get events for itinerary: %w", err)
	}

	itineraryEvents, err := StringsToInts(itinerary.Events)
	if err != nil {
		log.Printf("Error converting event IDs to integers: %v\n", err)
		return fmt.Errorf("failed to convert event IDs to integers: %w", err)
	}

	for _, eventID := range itineraryEvents {
		temp, err := GetEvent(DB, eventID)
		if err != nil {
			log.Printf("Error getting event ID %d: %v\n", eventID, err)
			return fmt.Errorf("failed to get event ID: %w", err)
		}

		price += temp.Cost
	}

	fmt.Println("Price Updated to:", price)
	return UpdateAttribute(DB, "itineraries", "itineraryId", itineraryId, "price", price)
}
