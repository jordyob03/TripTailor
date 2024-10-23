package models

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
