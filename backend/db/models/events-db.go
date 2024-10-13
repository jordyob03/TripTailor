package DBmodels

import (
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
)

type Event struct {
	EventId      int       `json:"eventId"`
	Name         string    `json:"name"`
	Price        int       `json:"price"`
	Location     string    `json:"location"`
	Description  string    `json:"description"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	ItineraryIds []int     `json:"itineraryIds"`
	PhotoLinks   []string  `json:"photoLinks"`
}

func CreateEventTable() error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS events (
		eventId SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		price INT NOT NULL,
		location TEXT NOT NULL,
		description TEXT,
		startDate TIMESTAMPTZ NOT NULL,
		endDate TIMESTAMPTZ NOT NULL,
		itineraryIds INT[],
		photoLinks TEXT[]
	);`

	return CreateTable(createTableSQL)
}

func AddEvent(event Event) (int, error) {
	insertEventSQL := `
	INSERT INTO events (name, price, location, description, startDate, endDate, itineraryIds, photoLinks)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING eventId;`

	var eventID int
	err := DB.QueryRow(
		insertEventSQL, event.Name, event.Price,
		event.Location, event.Description,
		event.StartDate, event.EndDate,
		event.ItineraryIds, event.PhotoLinks).Scan(&eventID)
	if err != nil {
		log.Printf("Error adding event: %v\n", err)
		return 0, fmt.Errorf("failed to add event: %w", err)
	}

	log.Printf("Event added successfully with ID: %d\n", eventID)
	return eventID, nil
}

func RemoveEvent(eventID int) error {
	var itineraryIds []int

	getItineraryIdsSQL := `SELECT itineraryIds FROM events WHERE eventId = $1;`
	err := DB.QueryRow(getItineraryIdsSQL, eventID).Scan(pq.Array(&itineraryIds))
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

	log.Println("Removal of Event from Associated Itinerary IDs:")
	for _, itineraryID := range itineraryIds {
		err := RemoveItineraryEvent(itineraryID, eventID)
		if err != nil {
			log.Printf("Error removing event from itinerary ID %d: %v\n", itineraryID, err)
			return fmt.Errorf("failed to remove event from itinerary: %w", err)
		}
	}

	log.Printf("Event with ID %d successfully deleted. Itinerary IDs: %v\n", eventID, itineraryIds)
	return nil
}

func GetEvent(eventID int) (Event, error) {
	query := `SELECT * FROM events WHERE eventId = $1;`

	var event Event
	err := DB.QueryRow(query, eventID).Scan(
		&event.EventId, &event.Name, &event.Price,
		&event.Location, &event.Description,
		&event.StartDate, &event.EndDate,
		pq.Array(&event.ItineraryIds), pq.Array(&event.PhotoLinks),
	)

	if err != nil {
		log.Printf("Error retrieving event: %v\n", err)
		return Event{}, fmt.Errorf("failed to retrieve event: %w", err)
	}

	log.Printf("Event with ID %d successfully retrieved.\n", eventID)
	return event, nil
}

func UpdateEventName(eventID int, name string) error {
	err := UpdateAttribute("events", "eventId", eventID, "name", name)
	if err != nil {
		log.Printf("Error updating event name: %v\n", err)
		return fmt.Errorf("failed to update event name: %w", err)
	}

	log.Printf("Event name updated successfully for ID %d.\n", eventID)
	return nil
}

func UpdateEventPrice(eventID int, price int) error {
	err := UpdateAttribute("events", "eventId", eventID, "price", price)
	if err != nil {
		log.Printf("Error updating event price: %v\n", err)
		return fmt.Errorf("failed to update event price: %w", err)
	}

	log.Printf("Event price updated successfully for ID %d.\n", eventID)
	return nil
}

func UpdateEventLocation(eventID int, location string) error {
	err := UpdateAttribute("events", "eventId", eventID, "location", location)
	if err != nil {
		log.Printf("Error updating event location: %v\n", err)
		return fmt.Errorf("failed to update event location: %w", err)
	}

	log.Printf("Event location updated successfully for ID %d.\n", eventID)
	return nil
}

func UpdateEventDescription(eventID int, description string) error {
	err := UpdateAttribute("events", "eventId", eventID, "description", description)
	if err != nil {
		log.Printf("Error updating event description: %v\n", err)
		return fmt.Errorf("failed to update event description: %w", err)
	}

	log.Printf("Event description updated successfully for ID %d.\n", eventID)
	return nil
}

func UpdateEventStartDate(eventID int, startDate time.Time) error {
	err := UpdateAttribute("events", "eventId", eventID, "startDate", startDate)
	if err != nil {
		log.Printf("Error updating event start date: %v\n", err)
		return fmt.Errorf("failed to update event start date: %w", err)
	}

	log.Printf("Event start date updated successfully for ID %d.\n", eventID)
	return nil
}

func UpdateEventEndDate(eventID int, endDate time.Time) error {
	err := UpdateAttribute("events", "eventId", eventID, "endDate", endDate)
	if err != nil {
		log.Printf("Error updating event end date: %v\n", err)
		return fmt.Errorf("failed to update event end date: %w", err)
	}

	log.Printf("Event end date updated successfully for ID %d.\n", eventID)
	return nil
}

func AddEventPhotoLink(eventID int, photoLink string) error {
	err := AddArrayAttribute("events", "eventId", eventID, "photoLinks", []string{photoLink})
	if err != nil {
		log.Printf("Error adding photo link to event: %v\n", err)
		return fmt.Errorf("failed to add photo link to event: %w", err)
	}

	log.Printf("Photo link added successfully for ID %d.\n", eventID)
	return nil
}

func RemoveEventPhotoLink(eventID int, photoLink string) error {
	err := RemoveArrayAttribute("events", "eventId", eventID, "photoLinks", []string{photoLink})
	if err != nil {
		log.Printf("Error removing photo link from event: %v\n", err)
		return fmt.Errorf("failed to remove photo link from event: %w", err)
	}

	log.Printf("Photo link removed successfully for ID %d.\n", eventID)
	return nil
}

func AddEventItinerary(eventID int, itineraryID int) error {
	err := AddArrayAttribute("events", "eventId", eventID, "itineraryIds", []int{itineraryID})
	if err != nil {
		log.Printf("Error adding itinerary to event: %v\n", err)
		return fmt.Errorf("failed to add itinerary to event: %w", err)
	}

	log.Printf("Itinerary added successfully for ID %d.\n", eventID)
	return nil
}

func RemoveEventItinerary(eventID int, itineraryID int) error {
	err := RemoveArrayAttribute("events", "eventId", eventID, "itineraryIds", []int{itineraryID})
	if err != nil {
		log.Printf("Error removing itinerary from event: %v\n", err)
		return fmt.Errorf("failed to remove itinerary from event: %w", err)
	}

	log.Printf("Itinerary removed successfully for ID %d.\n", eventID)
	return nil
}
