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
	Price       int       `json:"price"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	ItineraryId int       `json:"itineraryId"`
	EventImages []string  `json:"eventImages"`
}

func CreateEventTable(DB *sql.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS events (
		eventId SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		price INT NOT NULL,
		location TEXT NOT NULL,
		description TEXT,
		startDate TIMESTAMPTZ NOT NULL,
		endDate TIMESTAMPTZ NOT NULL,
		itineraryId INTEGER,
		eventImages TEXT[]
	);`

	return CreateTable(DB, createTableSQL)
}

func GetEvent(DB *sql.DB, eventID int) (Event, error) {
	query := `SELECT * FROM events WHERE eventId = $1;`

	var event Event
	err := DB.QueryRow(query, eventID).Scan(
		&event.EventId, &event.Name, &event.Price,
		&event.Location, &event.Description,
		&event.StartDate, &event.EndDate,
		&event.ItineraryId, pq.Array(&event.EventImages),
	)

	if err != nil {
		log.Printf("Error retrieving event: %v\n", err)
		return Event{}, fmt.Errorf("failed to retrieve event: %w", err)
	}

	log.Printf("Event with ID %d successfully retrieved.\n", eventID)
	return event, nil
}
