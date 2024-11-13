package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/lib/pq"
)

type User struct {
	UserId       int       `json:"userId"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	DateOfBirth  time.Time `json:"dateOfBirth"`
	Name         string    `json:"name"`
	Country      string    `json:"country"`
	Languages    []string  `json:"languages"`
	Tags         []string  `json:"tags"`
	Boards       []string  `json:"boards"`
	Posts        []string  `json:"posts"`
	Followers    []string  `json:"followers"`
	Following    []string  `json:"following"`
	ProfileImage int       `json:"profileImage"`
	CoverImage   int       `json:"coverImage"`
}

func CreateUserTable(DB *sql.DB) error {
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
		boards TEXT[],
		posts TEXT[],
		followers TEXT[],
		following TEXT[],
		profileImage INTEGER,
		coverImage INTEGER
	);`
	return CreateTable(DB, createTableSQL)
}

func AddUser(DB *sql.DB, user User) (int, error) {
	insertUserSQL := `
	INSERT INTO users (username, email, password, dateOfBirth, name, country, languages, tags, boards, posts, followers, following, profileImage, coverImage)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
	RETURNING userId;`

	if user.Followers == nil {
		user.Followers = []string{}
	}

	if user.Following == nil {
		user.Following = []string{}
	}

	if user.Languages == nil {
		user.Languages = []string{}
	}

	if user.Tags == nil {
		user.Tags = []string{}
	}

	var userId int
	err := DB.QueryRow(
		insertUserSQL, user.Username, user.Email, user.Password,
		user.DateOfBirth, user.Name, user.Country,
		pq.Array(user.Languages), pq.Array(user.Tags),
		pq.Array(user.Boards), pq.Array(user.Posts),
		pq.Array(user.Followers), pq.Array(user.Following),
		user.ProfileImage, user.CoverImage).Scan(&userId)
	if err != nil {
		fmt.Println("Error adding user:", err)
		return 0, err
	}

	return userId, nil
}

func RemoveUser(DB *sql.DB, username string) error {
	user, err := GetUser(DB, username)
	if err != nil {
		log.Printf("Error retrieving user: %v\n", err)
		return err
	}

	IntPosts, err := StringsToInts(user.Posts)
	if err != nil {
		log.Printf("Error converting post IDs to strings: %v\n", err)
		return err
	}

	IntBoards, err := StringsToInts(user.Boards)
	if err != nil {
		log.Printf("Error converting board IDs to strings: %v\n", err)
		return err
	}

	for _, follower := range user.Followers {
		err := RemoveUserFollowing(DB, follower, username)
		if err != nil {
			log.Printf("Error removing follower %s: %v\n", follower, err)
			return err
		}
	}

	for _, following := range user.Following {
		err := RemoveUserFollower(DB, following, username)
		if err != nil {
			log.Printf("Error removing following %s: %v\n", following, err)
			return err
		}
	}

	for _, post := range IntPosts {
		err := RemovePost(DB, post)
		if err != nil {
			log.Printf("Error removing post %d: %v\n", post, err)
			return err
		}
	}

	for _, board := range IntBoards {
		err := RemoveBoard(DB, board)
		if err != nil {
			log.Printf("Error removing board %d: %v\n", board, err)
			return err
		}
	}

	deleteUserSQL := `DELETE FROM users WHERE username = $1;`
	_, err = DB.Exec(deleteUserSQL, username)
	if err != nil {
		log.Printf("Error removing user: %v\n", err)
		return err
	}

	return nil
}

func GetUser(DB *sql.DB, username string) (User, error) {
	query := `
    SELECT userId, username, email, password, dateOfBirth, name, country, languages, tags, boards, posts, followers, following, profileImage, coverImage
    FROM users 
    WHERE username = $1`

	var user User
	row := DB.QueryRow(query, username)
	err := row.Scan(
		&user.UserId, &user.Username, &user.Email,
		&user.Password, &user.DateOfBirth, &user.Name, &user.Country,
		pq.Array(&user.Languages),
		pq.Array(&user.Tags),
		pq.Array(&user.Boards),
		pq.Array(&user.Posts),
		pq.Array(&user.Followers),
		pq.Array(&user.Following),
		&user.ProfileImage,
		&user.CoverImage,
	)

	if err == sql.ErrNoRows {
		return User{}, fmt.Errorf("no user found with username: %s", username)
	} else if err != nil {
		log.Printf("Error retrieving user: %v\n", err)
		return User{}, err
	}

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

	if user.Followers == nil {
		user.Followers = []string{}
	}

	if user.Following == nil {
		user.Following = []string{}
	}

	return user, nil
}

func GetAllUsers(DB *sql.DB) ([]User, error) {
	query := `
    SELECT userId, username, email, password, dateOfBirth, name, country, languages, tags, boards, posts, followers, following, profileImage, coverImage
    FROM users`

	rows, err := DB.Query(query)
	if err != nil {
		log.Printf("Error querying users: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(
			&user.UserId, &user.Username, &user.Email,
			&user.Password, &user.DateOfBirth, &user.Name, &user.Country,
			pq.Array(&user.Languages),
			pq.Array(&user.Tags),
			pq.Array(&user.Boards),
			pq.Array(&user.Posts),
			pq.Array(&user.Followers),
			pq.Array(&user.Following),
			&user.ProfileImage,
			&user.CoverImage,
		); err != nil {
			log.Printf("Error scanning user: %v\n", err)
			return nil, err
		}

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

		if user.Followers == nil {
			user.Followers = []string{}
		}

		if user.Following == nil {
			user.Following = []string{}
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over users: %v\n", err)
		return nil, err
	}

	return users, nil
}

func UpdateName(dbConn *sql.DB, username string, name string) error {
	return UpdateAttribute(dbConn, "users", "username", username, "name", name)
}

func UpdateUserEmail(DB *sql.DB, username, email string) error {
	return UpdateAttribute(DB, "users", "username", username, "email", email)
}

func UpdateUserPassword(DB *sql.DB, username, password string) error {
	return UpdateAttribute(DB, "users", "username", username, "password", password)
}

func UpdateUserDateOfBirth(DB *sql.DB, username string, dateOfBirth time.Time) error {
	return UpdateAttribute(DB, "users", "username", username, "dateOfBirth", dateOfBirth)
}

func UpdateUserCountry(DB *sql.DB, username, country string) error {
	return UpdateAttribute(DB, "users", "username", username, "country", country)
}

func AddUserLanguage(DB *sql.DB, username string, languages []string) error {
	err := AddArrayAttribute(DB, "users", "username", username, "languages", languages)
	if err != nil {
		log.Printf("Error adding languages for user %s: %v\n", username, err)
		return err
	}

	return nil
}

func RemoveUserLanguage(DB *sql.DB, username string, languages []string) error {
	err := RemoveArrayAttribute(DB, "users", "username", username, "languages", languages)
	if err != nil {
		log.Printf("Error removing languages for user %s: %v\n", username, err)
		return err
	}

	return nil
}

func AddUserTag(DB *sql.DB, username string, tags []string) error {
	err := AddArrayAttribute(DB, "users", "username", username, "tags", tags)
	if err != nil {
		log.Printf("Error adding tags for user %s: %v\n", username, err)
		return err
	}

	return nil
}

func RemoveUserTag(DB *sql.DB, username string, tags []string) error {
	err := RemoveArrayAttribute(DB, "users", "username", username, "tags", tags)
	if err != nil {
		log.Printf("Error removing tags for user %s: %v\n", username, err)
		return err
	}

	return nil
}

func AddUserBoard(DB *sql.DB, username string, boardId int) error {
	err := AddArrayAttribute(DB, "users", "username", username, "boards", IntsToStrings([]int{boardId}))
	if err != nil {
		log.Printf("Error adding board for user %s: %v\n", username, err)
		return err
	}

	return nil
}

func RemoveUserBoard(DB *sql.DB, username string, boardId int) error {
	err := RemoveArrayAttribute(DB, "users", "username", username, "boards", IntsToStrings([]int{boardId}))
	if err != nil {
		log.Printf("Error removing board for user %s: %v\n", username, err)
		return err
	}

	return nil
}

func AddUserPost(DB *sql.DB, username string, postId int) error {
	err := AddArrayAttribute(DB, "users", "username", username, "posts", IntsToStrings([]int{postId}))
	if err != nil {
		log.Printf("Error adding post for user %s: %v\n", username, err)
		return err
	}

	return nil
}

func RemoveUserPost(DB *sql.DB, username string, postId int) error {
	err := RemoveArrayAttribute(DB, "users", "username", username, "posts", IntsToStrings([]int{postId}))
	if err != nil {
		log.Printf("Error removing post for user %s: %v\n", username, err)
		return err
	}

	return nil
}

func AddUserFollower(DB *sql.DB, username string, follower string) error {
	return AddArrayAttribute(DB, "users", "username", username, "followers", []string{follower})
}

func RemoveUserFollower(DB *sql.DB, username string, follower string) error {
	return RemoveArrayAttribute(DB, "users", "username", username, "followers", []string{follower})
}

func AddUserFollowing(DB *sql.DB, username string, following string) error {
	return AddArrayAttribute(DB, "users", "username", username, "following", []string{following})
}

func RemoveUserFollowing(DB *sql.DB, username string, following string) error {
	return RemoveArrayAttribute(DB, "users", "username", username, "following", []string{following})
}

func UpdateUserProfileImage(DB *sql.DB, username string, imageId int) error {
	image, err := GetImage(DB, imageId)
	if err != nil {
		log.Printf("Error retrieving image: %v\n", err)
		return err
	}

	for _, metadata := range image.Metadata {
		if metadata == "profile" {
			fmt.Printf("Image %d is already a profile image", imageId)
			return nil
		}
	}

	err = AddImageMetaData(DB, imageId, "profile")
	if err != nil {
		log.Printf("Error adding image metadata: %v\n", err)
		return err
	}

	return UpdateAttribute(DB, "users", "username", username, "profileImage", imageId)
}

func RemoveUserProfileImage(DB *sql.DB, username string) error {
	return UpdateAttribute(DB, "users", "username", username, "profileImage", '0')
}

func UpdateUserCoverImage(DB *sql.DB, username string, imageId int) error {
	image, err := GetImage(DB, imageId)
	if err != nil {
		log.Printf("Error retrieving image: %v\n", err)
		return err
	}

	for _, metadata := range image.Metadata {
		if metadata == "cover" {
			fmt.Printf("Image %d is already a cover image", imageId)
			return nil
		}
	}

	err = AddImageMetaData(DB, imageId, "cover")
	if err != nil {
		log.Printf("Error adding image metadata: %v\n", err)
		return err
	}

	return UpdateAttribute(DB, "users", "username", username, "coverImage", imageId)
}

func RemoveUserCoverImage(DB *sql.DB, username string) error {
	return UpdateAttribute(DB, "users", "username", username, "coverImage", '1')
}

func UserHandler(DB *sql.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		username := r.URL.Query().Get("username")

		if username != "" {
			user, err := GetUser(DB, username)
			if err != nil {
				http.Error(w, "User not found", http.StatusNotFound)
				log.Printf("Error retrieving user: %v\n", err)
				return
			}
			json.NewEncoder(w).Encode(user)
		} else {
			users, err := GetAllUsers(DB)
			if err != nil {
				http.Error(w, "Error retrieving users", http.StatusInternalServerError)
				log.Printf("Error retrieving users: %v\n", err)
				return
			}
			json.NewEncoder(w).Encode(users)
		}

	case "POST":
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		userId, err := AddUser(DB, user)
		if err != nil {
			http.Error(w, "Error adding user", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int{"userId": userId})

	case "PUT":
		username := r.URL.Query().Get("username")
		if username == "" {
			http.Error(w, "Username is required", http.StatusBadRequest)
			return
		}

		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			log.Printf("Error decoding request body: %v\n", err)
			return
		}

		data := make(map[string]interface{})
		if user.Username != "" {
			data["username"] = user.Username
		}
		if user.Password != "" {
			data["password"] = user.Password
		}
		if !user.DateOfBirth.IsZero() {
			data["dateOfBirth"] = user.DateOfBirth
		}
		if user.Name != "" {
			data["name"] = user.Name
		}
		if user.Country != "" {
			data["country"] = user.Country
		}

		w.WriteHeader(http.StatusOK)
		return

	case "DELETE":
		username := r.URL.Query().Get("username")
		if username == "" {
			http.Error(w, "Username is required", http.StatusBadRequest)
			return
		}

		if err := RemoveUser(DB, username); err != nil {
			http.Error(w, "Error deleting user", http.StatusInternalServerError)
			log.Printf("Error deleting user: %v\n", err)
			return
		}

		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func UserHandlerWrapper(DB *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		UserHandler(DB, w, r)
	}
}
