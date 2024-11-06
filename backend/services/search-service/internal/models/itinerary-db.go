package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
)

type Itinerary struct {
	ItineraryId int      `json:"itineraryId"`
	Name        string   `json:"name"`
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
		name TEXT NOT NULL,
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

func GetItinerary(DB *sql.DB, itineraryID int) (Itinerary, error) {
	getItinerarySQL := `
	SELECT name, city, country, title, description, price, languages, tags, events, postId, username
	FROM itineraries
	WHERE itineraryId = $1;`

	var itinerary Itinerary

	err := DB.QueryRow(getItinerarySQL, itineraryID).Scan(
		&itinerary.Name,
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
