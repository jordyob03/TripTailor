package db

import (
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	assert "github.com/stretchr/testify/assert"
	testing "testing"
	"time"
)

func TestQueryItinerariesByLocation(t *testing.T) {
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer database.Close()

	country := "Tokyo"
	city := "Japan"

	query := `
        SELECT itineraryid, name, city, country, languages, tags, events, postid, username, creationdate, lastupdate
        FROM itineraries
        WHERE country = \$1 AND city = \$2`

	creationDate, _ := time.Parse("2006-01-02", "2021-01-01")
	lastUpdate, _ := time.Parse("2006-01-02", "2021-01-02")

	rows := sqlmock.NewRows([]string{"itineraryid", "name", "city", "country", "languages", "tags", "events", "postid", "username", "creationdate", "lastupdate"}).
		AddRow(1, "Tokyo Forest Exploration", "Tokyo", "Japan", "English,Japanese", "Food,History", "1,2", 1, "johndoe", creationDate, lastUpdate).
		AddRow(2, "Tokyo Furry Convention", "Tokyo", "Japan", "English", "Food", "1", 2, "janedoe", creationDate, lastUpdate)

	mock.ExpectQuery(query).WithArgs(country, city).WillReturnRows(rows)

	itineraries, err := QueryItinerariesByLocation(database, country, city)

	assert.NoError(t, err)
	assert.Len(t, itineraries, 2)

	assert.Equal(t, 1, itineraries[0].PostId)
	assert.Equal(t, "Tokyo Forest Exploration", itineraries[0].Name)
	assert.Equal(t, "Tokyo", itineraries[0].City)
	assert.Equal(t, "Japan", itineraries[0].Country)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}
