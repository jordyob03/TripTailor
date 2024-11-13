package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
)

type Event struct {
	EventId     int       `json:"eventId"`
	Name        string    `json:"name"`
	Cost        float64   `json:"cost"`
	Address     string    `json:"address"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	ItineraryId int       `json:"itineraryId"`
	EventImages []string  `json:"eventImages"`
}

func CreateEventTable(DB *sql.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS events (
		eventId SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		cost FLOAT NOT NULL,
		address TEXT NOT NULL,
		description TEXT,
		startTime TIMESTAMPTZ NOT NULL,
		endTime TIMESTAMPTZ NOT NULL,
		itineraryId INTEGER,
		eventImages TEXT[]
	);`

	return CreateTable(DB, createTableSQL)
}

func GetEvent(DB *sql.DB, eventID int) (Event, error) {
	query := `SELECT * FROM events WHERE eventId = $1;`

	var event Event
	err := DB.QueryRow(query, eventID).Scan(
		&event.EventId, &event.Name, &event.Cost,
		&event.Address, &event.Description,
		&event.StartTime, &event.EndTime,
		&event.ItineraryId, pq.Array(&event.EventImages),
	)

	if err != nil {
		log.Printf("Error retrieving event: %v\n", err)
		return Event{}, fmt.Errorf("failed to retrieve event: %w", err)
	}

	if event.EventImages == nil {
		event.EventImages = []string{}
	}

	log.Printf("Event with ID %d successfully retrieved.\n", eventID)
	return event, nil
}
