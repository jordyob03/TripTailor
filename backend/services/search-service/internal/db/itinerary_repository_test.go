package db

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	assert "github.com/stretchr/testify/assert"
)

func TestQueryItinerariesByLocation(t *testing.T) {
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer database.Close()

	country := "Japan"
	city := "Tokyo"

	query := `
        SELECT itineraryid, name, city, country, title, description, price, languages, tags, events, postid, username
        FROM itineraries
        WHERE country = \$1 AND city = \$2`

	rows := sqlmock.NewRows([]string{"itineraryid", "name", "city", "country", "title", "description", "price", "languages", "tags", "events", "postid", "username"}).
		AddRow(1, "Tokyo Forest Exploration", "Tokyo", "Japan", "Forest Adventure", "Explore the forests of Tokyo", 150.00, "English,Japanese", "Nature,Adventure", "1,2", 1, "johndoe").
		AddRow(2, "Tokyo Nerd Convention", "Tokyo", "Japan", "Anime Expo", "A convention for anime lovers", 50.00, "English", "Anime,Convention", "3", 2, "janedoe")

	mock.ExpectQuery(query).WithArgs(country, city).WillReturnRows(rows)

	itineraries, err := QueryItinerariesByLocation(database, country, city)

	assert.NoError(t, err)
	assert.Len(t, itineraries, 2)

	// Assert for the first itinerary
	assert.Equal(t, 1, itineraries[0].ItineraryId)
	assert.Equal(t, "Tokyo Forest Exploration", itineraries[0].Name)
	assert.Equal(t, "Tokyo", itineraries[0].City)
	assert.Equal(t, "Japan", itineraries[0].Country)
	assert.Equal(t, "Forest Adventure", itineraries[0].Title)
	assert.Equal(t, "Explore the forests of Tokyo", itineraries[0].Description)
	assert.Equal(t, 150.00, itineraries[0].Price)
	assert.Equal(t, []string{"English", "Japanese"}, itineraries[0].Languages)
	assert.Equal(t, []string{"Nature", "Adventure"}, itineraries[0].Tags)
	assert.Equal(t, []string{"1", "2"}, itineraries[0].Events)
	assert.Equal(t, 1, itineraries[0].PostId)
	assert.Equal(t, "johndoe", itineraries[0].Username)

	// Assert for the second itinerary
	assert.Equal(t, 2, itineraries[1].ItineraryId)
	assert.Equal(t, "Tokyo Nerd Convention", itineraries[1].Name)
	assert.Equal(t, "Tokyo", itineraries[1].City)
	assert.Equal(t, "Japan", itineraries[1].Country)
	assert.Equal(t, "Anime Expo", itineraries[1].Title)
	assert.Equal(t, "A convention for anime lovers", itineraries[1].Description)
	assert.Equal(t, 50.00, itineraries[1].Price)
	assert.Equal(t, []string{"English"}, itineraries[1].Languages)
	assert.Equal(t, []string{"Anime", "Convention"}, itineraries[1].Tags)
	assert.Equal(t, []string{"3"}, itineraries[1].Events)
	assert.Equal(t, 2, itineraries[1].PostId)
	assert.Equal(t, "janedoe", itineraries[1].Username)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
