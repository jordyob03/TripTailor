package DBmodels

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
)

type User struct {
	UserId      int      `json:"userId"`
	Username    string   `json:"username"`
	Email       string   `json:"email"`
	Password    string   `json:"password"`
	DateOfBirth string   `json:"dateOfBirth"` // Changed to string to avoid unnecessary parsing here
	Name        string   `json:"name"`
	Country     string   `json:"country"`
	Languages   []string `json:"languages"`
	Tags        []string `json:"tags"`
	Boards      []string `json:"boards"`
	Posts       []string `json:"posts"`
}

// CreateUserTable creates the user table if it doesn't exist
func CreateUserTable(dbConn *sql.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		userId SERIAL PRIMARY KEY,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		dateOfBirth DATE NOT NULL,
		name TEXT,
		country TEXT,
		languages TEXT[],
		tags TEXT[],
		boards INTEGER[],
		posts INTEGER[]
	);`
	_, err := dbConn.Exec(createTableSQL)
	return err
}

// AddUser adds a new user to the database
func AddUser(dbConn *sql.DB, user User) (int, error) {
	insertUserSQL := `
	INSERT INTO users (username, email, password, dateOfBirth, name, country, languages, tags, boards, posts)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING userId;`

	var userId int
	err := dbConn.QueryRow(
		insertUserSQL, user.Username, user.Email, user.Password,
		user.DateOfBirth, user.Name, user.Country,
		pq.Array(user.Languages), pq.Array(user.Tags),
		pq.Array(user.Boards), pq.Array(user.Posts)).Scan(&userId)
	if err != nil {
		fmt.Println("Error adding user:", err)
		return 0, err
	}

	return userId, nil
}

// GetUser retrieves a user by username from the database
func GetUser(dbConn *sql.DB, username string) (User, error) {
	query := `
    SELECT userId, username, email, password, dateOfBirth, name, country, languages, tags, boards, posts
    FROM users 
    WHERE username = $1`

	var user User
	row := dbConn.QueryRow(query, username)
	err := row.Scan(
		&user.UserId, &user.Username, &user.Email,
		&user.Password, &user.DateOfBirth, &user.Name, &user.Country,
		pq.Array(&user.Languages), pq.Array(&user.Tags),
		pq.Array(&user.Boards), pq.Array(&user.Posts),
	)

	if err == sql.ErrNoRows {
		return User{}, fmt.Errorf("no user found with username: %s", username)
	} else if err != nil {
		log.Printf("Error retrieving user: %v\n", err)
		return User{}, err
	}

	// Ensure slices are not nil
	if user.Languages == nil {
		user.Languages = []string{}
	}
	if user.Tags == nil {
		user.Tags = []string{}
	}
	if user.Boards == nil {
		user.Boards = []string{}
	}
	if user.Posts == nil {
		user.Posts = []string{}
	}

	return user, nil
}

// UpdateUserCountry updates a user's country in the database
func UpdateUserCountry(dbConn *sql.DB, username string, country string) error {
	query := `UPDATE users SET country = $1 WHERE username = $2`
	_, err := dbConn.Exec(query, country, username)
	if err != nil {
		return fmt.Errorf("error updating country for user %s: %v", username, err)
	}
	return nil
}

// AddUserLanguage adds languages to a user
func AddUserLanguage(dbConn *sql.DB, username string, languages []string) error {
	query := `UPDATE users SET languages = array_append(languages, unnest($1::text[])) WHERE username = $2`
	_, err := dbConn.Exec(query, pq.Array(languages), username)
	if err != nil {
		return fmt.Errorf("error adding languages for user %s: %v", username, err)
	}
	return nil
}

// AddUserTag adds tags to a user
func AddUserTag(dbConn *sql.DB, username string, tags []string) error {
	query := `UPDATE users SET tags = array_append(tags, unnest($1::text[])) WHERE username = $2`
	_, err := dbConn.Exec(query, pq.Array(tags), username)
	if err != nil {
		return fmt.Errorf("error adding tags for user %s: %v", username, err)
	}
	return nil
}

// RemoveUser removes a user and their related data (boards, posts) from the database
func RemoveUser(dbConn *sql.DB, username string) error {
	getPostsAndBoardsSQL := `
	SELECT posts, boards 
	FROM users 
	WHERE username = $1;
	`

	var posts, boards []string

	err := dbConn.QueryRow(getPostsAndBoardsSQL, username).Scan(pq.Array(&posts), pq.Array(&boards))
	if err != nil {
		log.Printf("Error retrieving posts and boards for user %s: %v\n", username, err)
		return err
	}

	deleteUserSQL := `DELETE FROM users WHERE username = $1;`

	_, err = dbConn.Exec(deleteUserSQL, username)
	if err != nil {
		log.Printf("Error removing user: %v\n", err)
		return err
	}

	return nil
}
