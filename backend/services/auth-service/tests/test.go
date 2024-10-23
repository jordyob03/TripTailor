package handlers_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jordyob03/TripTailor/backend/services/auth-service/internal/db"
	"github.com/jordyob03/TripTailor/backend/services/auth-service/internal/handlers"
	"github.com/jordyob03/TripTailor/backend/services/auth-service/internal/models"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestSignIn(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockDB := &sql.DB{} // You can replace this with a mock if needed
	username := "testuser"
	password := "password123"

	// Mock GetUser to return a user with a hashed password
	db.GetUser = func(dbConn *sql.DB, username string) (models.User, error) {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		return models.User{UserId: 1, Username: username, Password: string(hashedPassword)}, nil
	}

	// Create the test context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Create a SignInRequest payload
	reqBody, _ := json.Marshal(handlers.SignInRequest{
		Username: username,
		Password: password,
	})

	// Simulate an HTTP request
	c.Request, _ = http.NewRequest("POST", "/signin", bytes.NewBuffer(reqBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Call the SignIn handler
	handlers.SignIn(mockDB)(c)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Signin successful", response["message"])
	assert.Equal(t, float64(1), response["userId"])
	assert.NotEmpty(t, response["token"])
}

func TestSignUp(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockDB := &sql.DB{} // You can replace this with a mock if needed
	username := "newuser"
	password := "newpassword123"
	email := "test@example.com"
	dateOfBirth := "1990-01-01"

	// Mock AddUser to return a valid user ID
	db.AddUser = func(dbConn *sql.DB, user models.User) (int64, error) {
		return 1, nil
	}

	// Mock GetUser to return an error, simulating a non-existing user
	db.GetUser = func(dbConn *sql.DB, username string) (models.User, error) {
		return models.User{}, sql.ErrNoRows
	}

	// Create the test context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Create a SignUpRequest payload
	reqBody, _ := json.Marshal(handlers.SignUpRequest{
		Username:    username,
		Email:       email,
		Password:    password,
		DateOfBirth: dateOfBirth,
	})

	// Simulate an HTTP request
	c.Request, _ = http.NewRequest("POST", "/signup", bytes.NewBuffer(reqBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Call the SignUp handler
	handlers.SignUp(mockDB)(c)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "User created successfully", response["message"])
	assert.Equal(t, float64(1), response["userId"])
	assert.NotEmpty(t, response["token"])
}
