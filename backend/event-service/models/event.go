package models

// Event struct definition
type Event struct {
	EventId          int
	EventName        string
	EventDate        string
	EventTime        string
	EventLocation    string
	EventPrice       float64
	EventDescription string
	EventPhotos      []string
}
