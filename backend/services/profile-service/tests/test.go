package handlers_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jordyob03/TripTailor/backend/services/profile-service/internal/db"
	"github.com/jordyob03/TripTailor/backend/services/profile-service/internal/handlers"
	"github.com/stretchr/testify/assert"
)

func TestCreateProfile(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockDB := &sql.DB{}

	db.GetUser = func(dbConn *sql.DB, username string) (db.User, error) {
		return db.User{
			Username:  username,
			Country:   "USA",
			Tags:      []string{"tag1", "tag2"},
			Languages: []string{"English", "Spanish"},
		}, nil
	}

	db.UpdateUserCountry = func(dbConn *sql.DB, username string, country string) error {
		return nil
	}

	// Mock language and tag update
	db.AddUserLanguage = func(dbConn *sql.DB, username string, languages []string) error {
		return nil
	}

	db.AddUserTag = func(dbConn *sql.DB, username string, tags []string) error {
		return nil
	}

	// Create HTTP test context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Define request payload
	profileReq := db.User{
		Username:  "testuser",
		Country:   "USA",
		Tags:      []string{"tag1", "tag2"},
		Languages: []string{"English", "Spanish"},
	}

	// Create request body
	reqBody, _ := json.Marshal(profileReq)
	c.Request, _ = http.NewRequest("POST", "/profile/create", bytes.NewBuffer(reqBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Call CreateProfile handler
	handlers.CreateProfile(mockDB)(c)

	// Assert that the response code is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert response body
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Profile updated successfully", response["message"])
}

func TestUpdateProfile(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Mock DB
	mockDB := &sql.DB{}

	// Mock User
	db.GetUser = func(dbConn *sql.DB, username string) (db.User, error) {
		return db.User{
			Username:  username,
			Country:   "USA",
			Tags:      []string{"tag1", "tag2"},
			Languages: []string{"English", "Spanish"},
		}, nil
	}

	// Mock successful country update
	db.UpdateUserCountry = func(dbConn *sql.DB, username string, country string) error {
		return nil
	}

	// Mock language and tag update
	db.AddUserLanguage = func(dbConn *sql.DB, username string, languages []string) error {
		return nil
	}

	db.AddUserTag = func(dbConn *sql.DB, username string, tags []string) error {
		return nil
	}

	// Create HTTP test context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Define request payload
	profileReq := handlers.CreateProfileRequest{
		Username: "testuser",
		Country:  "Canada",
		Language: "French",
		Tags:     "tag1,tag2,tag3",
	}

	// Create request body
	reqBody, _ := json.Marshal(profileReq)
	c.Request, _ = http.NewRequest("POST", "/profile/update", bytes.NewBuffer(reqBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Call UpdateProfile handler
	handlers.UpdateProfile(mockDB)(c)

	// Assert that the response code is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert response body
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Profile updated successfully", response["message"])
}
