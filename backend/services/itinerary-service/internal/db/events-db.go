package DBmodels

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

func AddEvent(DB *sql.DB, event Event) (int, error) {
	insertEventSQL := `
	INSERT INTO events (name, cost, address, description, startTime, endTime, itineraryId, eventImages)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING eventId;`

	var eventID int
	err := DB.QueryRow(
		insertEventSQL, event.Name, event.Cost,
		event.Address, event.Description,
		event.StartTime, event.EndTime,
		event.ItineraryId, pq.Array(event.EventImages)).Scan(&eventID)
	if err != nil {
		log.Printf("Error adding event: %v\n", err)
		return 0, fmt.Errorf("failed to add event: %w", err)
	}

	err = AddItineraryEvent(DB, event.ItineraryId, eventID)
	if err != nil {
		log.Printf("Error adding event to itinerary: %v\n", err)
		return 0, fmt.Errorf("failed to add event to itinerary: %w", err)
	}

	log.Printf("Event added successfully with ID: %d\n", eventID)
	return eventID, nil
}

func RemoveEvent(DB *sql.DB, eventID int) error {

	getImages := `SELECT eventImages FROM events WHERE eventId = $1;`
	var imageStringIDs []string
	err := DB.QueryRow(getImages, eventID).Scan(pq.Array(&imageStringIDs))
	if err != nil {
		log.Printf("Error retrieving image IDs for event ID %d: %v\n", eventID, err)
		return fmt.Errorf("failed to retrieve image IDs: %w", err)
	}

	imageIDs, err := StringsToInts(imageStringIDs)
	if err != nil {
		log.Printf("Error converting image IDs to integers: %v\n", err)
		return fmt.Errorf("failed to convert image IDs: %w", err)
	}

	for _, imageID := range imageIDs {
		err := RemoveEventImage(DB, eventID, imageID)
		if err != nil {
			log.Printf("Error removing image metadata: %v\n", err)
			return fmt.Errorf("failed to remove image metadata: %w", err)
		}
	}

	err = RemoveItineraryEvent(DB, eventID, eventID)
	if err != nil {
		log.Printf("Error retrieving itinerary IDs for event ID %d: %v\n", eventID, err)
		return fmt.Errorf("failed to retrieve itinerary IDs: %w", err)
	}

	deleteEventSQL := `DELETE FROM events WHERE eventId = $1;`
	_, err = DB.Exec(deleteEventSQL, eventID)
	if err != nil {
		log.Printf("Error deleting event: %v\n", err)
		return fmt.Errorf("failed to delete event: %w", err)
	}

	log.Printf("Event with ID %d successfully deleted.\n", eventID)
	return nil
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

	log.Printf("Event with ID %d successfully retrieved.\n", eventID)
	return event, nil
}

func UpdateEventName(DB *sql.DB, eventID int, name string) error {
	err := UpdateAttribute(DB, "events", "eventId", eventID, "name", name)
	if err != nil {
		log.Printf("Error updating event name: %v\n", err)
		return fmt.Errorf("failed to update event name: %w", err)
	}

	log.Printf("Event name updated successfully for ID %d.\n", eventID)
	return nil
}

func UpdateEventPrice(DB *sql.DB, eventID int, cost int) error {
	err := UpdateAttribute(DB, "events", "eventId", eventID, "cost", cost)
	if err != nil {
		log.Printf("Error updating event cost: %v\n", err)
		return fmt.Errorf("failed to update event cost: %w", err)
	}

	log.Printf("Event cost updated successfully for ID %d.\n", eventID)
	return nil
}

func UpdateEventLocation(DB *sql.DB, eventID int, address string) error {
	err := UpdateAttribute(DB, "events", "eventId", eventID, "address", address)
	if err != nil {
		log.Printf("Error updating event address: %v\n", err)
		return fmt.Errorf("failed to update event address: %w", err)
	}

	log.Printf("Event address updated successfully for ID %d.\n", eventID)
	return nil
}

func UpdateEventDescription(DB *sql.DB, eventID int, description string) error {
	err := UpdateAttribute(DB, "events", "eventId", eventID, "description", description)
	if err != nil {
		log.Printf("Error updating event description: %v\n", err)
		return fmt.Errorf("failed to update event description: %w", err)
	}

	log.Printf("Event description updated successfully for ID %d.\n", eventID)
	return nil
}

func UpdateEventStartDate(DB *sql.DB, eventID int, startTime time.Time) error {
	err := UpdateAttribute(DB, "events", "eventId", eventID, "startTime", startTime)
	if err != nil {
		log.Printf("Error updating event start date: %v\n", err)
		return fmt.Errorf("failed to update event start date: %w", err)
	}

	log.Printf("Event start date updated successfully for ID %d.\n", eventID)
	return nil
}

func UpdateEventEndDate(DB *sql.DB, eventID int, endTime time.Time) error {
	err := UpdateAttribute(DB, "events", "eventId", eventID, "endTime", endTime)
	if err != nil {
		log.Printf("Error updating event end date: %v\n", err)
		return fmt.Errorf("failed to update event end date: %w", err)
	}

	log.Printf("Event end date updated successfully for ID %d.\n", eventID)
	return nil
}

func AddEventImage(DB *sql.DB, eventID int, imageID int) error {
	err := AddImageMetaData(DB, imageID, "event")
	if err != nil {
		log.Printf("Error adding image metadata: %v\n", err)
		return fmt.Errorf("failed to add image metadata: %w", err)
	}

	err = AddArrayAttribute(DB, "events", "eventId", eventID, "eventImages", IntsToStrings([]int{imageID}))
	if err != nil {
		log.Printf("Error adding image to event: %v\n", err)
		return fmt.Errorf("failed to add image to event: %w", err)
	}

	log.Printf("Image added successfully for ID %d.\n", eventID)
	return nil
}

func RemoveEventImage(DB *sql.DB, eventID int, imageID int) error {
	query := `SELECT eventId FROM events WHERE $1 = ANY(eventImages);`
	rows, err := DB.Query(query, imageID)
	if err != nil {
		log.Printf("Error retrieving events with image ID %d: %v\n", imageID, err)
		return fmt.Errorf("failed to retrieve events with image ID: %w", err)
	}

	var eventIDs []int
	for rows.Next() {
		var eventID int
		err := rows.Scan(&eventID)
		if err != nil {
			log.Printf("Error scanning event ID: %v\n", err)
			return fmt.Errorf("failed to scan event ID: %w", err)
		}

		eventIDs = append(eventIDs, eventID)
	}

	if len(eventIDs) == 1 {
		err := RemoveImageMetaData(DB, imageID, "event")
		if err != nil {
			log.Printf("Error removing image metadata: %v\n", err)
			return fmt.Errorf("failed to remove image metadata: %w", err)
		}
	}

	err = RemoveArrayAttribute(DB, "events", "eventId", eventID, "eventImages", IntsToStrings([]int{imageID}))
	if err != nil {
		log.Printf("Error removing image from event: %v\n", err)
		return fmt.Errorf("failed to remove image from event: %w", err)
	}

	log.Printf("Image removed successfully for ID %d.\n", eventID)
	return nil
}
