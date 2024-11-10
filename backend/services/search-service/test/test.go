package handlers_test

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	db "github.com/jordyob03/TripTailor/backend/services/search-service/internal/db"
	"github.com/jordyob03/TripTailor/backend/services/search-service/internal/handlers"
	"github.com/stretchr/testify/assert"
)

func TestSearchItineraries(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockDB := &sql.DB{}

	expectedItineraries := []db.Itinerary{
		{ItineraryId: 1, Country: "France", City: "Paris", Title: "Eiffel Tower Tour"},
		{ItineraryId: 2, Country: "France", City: "Paris", Title: "Louvre Museum Visit"},
	}

	db.QueryItinerariesByLocation = func(dbConn *sql.DB, country, city string) ([]db.Itinerary, error) {
		return expectedItineraries, nil
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("GET", "/search?country=France&city=Paris", nil)

	handlers.SearchItineraries(mockDB)(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []db.Itinerary
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedItineraries, response)
}

func TestSearchItineraries_MissingParams(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockDB := &sql.DB{}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("GET", "/search?country=France", nil)

	handlers.SearchItineraries(mockDB)(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Missing query parameters", response["error"])
}

func TestSearchItineraries_DatabaseError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockDB := &sql.DB{}

	db.QueryItinerariesByLocation = func(dbConn *sql.DB, country, city string) ([]db.Itinerary, error) {
		return nil, sql.ErrConnDone
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("GET", "/search?country=France&city=Paris", nil)

	handlers.SearchItineraries(mockDB)(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, sql.ErrConnDone.Error(), response["error"])
}
