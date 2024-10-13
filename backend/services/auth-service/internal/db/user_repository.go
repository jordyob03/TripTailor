package db

import (
	"database/sql"
	"fmt"
	"github.com/jordyob03/TripTailor/backend/services/auth-service/internal/models"
	"time"
)

func AddUser(db *sql.DB, username, email, password string, dateOfBirth time.Time) (int, error) {
	insertUserSQL := `
    INSERT INTO users (username, email, password, dateOfBirth)
    VALUES ($1, $2, $3, $4)
    RETURNING userId;`

	var userId int
	err := db.QueryRow(insertUserSQL, username, email, password, dateOfBirth).Scan(&userId)
	if err != nil {
		fmt.Println("Error adding user:", err)
		return 0, err
	}
	return userId, nil
}

func GetUser(db *sql.DB, username string) (models.User, error) {
	query := `SELECT userId, username, email, password, dateOfBirth FROM users WHERE username = $1`
	row := db.QueryRow(query, username)

	var user models.User
	err := row.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.DateOfBirth)
	if err != nil {
		return models.User{}, fmt.Errorf("no user found with username: %s", username)
	}
	return user, nil
}
