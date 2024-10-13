package DBmodels

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
	UserId      int       `json:"userId"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Name        string    `json:"name"`
	Country     string    `json:"country"`
	Languages   []string  `json:"languages"`
	Tags        []string  `json:"tags"`
	Boards      []string  `json:"boards"`
	Posts       []string  `json:"posts"`
}

func CreateUserTable() error {
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
	return CreateTable(createTableSQL)
}

func AddUser(user User) (int, error) {
	insertUserSQL := `
	INSERT INTO users (username, email, password, dateOfBirth, name, country, languages, tags, boards, posts)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING userId;`

	var userId int
	err := DB.QueryRow(
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

func GetUser(username string) (User, error) {
	query := `
    SELECT userId, username, email, password, dateOfBirth, name, country, languages, tags, boards, posts
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

	return user, nil
}

func GetAllUsers() ([]User, error) {
	query := `
    SELECT userId, username, email, password, dateOfBirth, name, country, languages, tags, boards, posts
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

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over users: %v\n", err)
		return nil, err
	}

	return users, nil
}

func UpdateUserEmail(username, email string) error {
	return UpdateAttribute("users", "username", username, "email", email)
}

func UpdateUserPassword(username, password string) error {
	return UpdateAttribute("users", "username", username, "password", password)
}

func UpdateUserDateOfBirth(username string, dateOfBirth time.Time) error {
	return UpdateAttribute("users", "username", username, "dateOfBirth", dateOfBirth)
}

func UpdateUserCountry(username, country string) error {
	return UpdateAttribute("users", "username", username, "country", country)
}

func DeleteUser(username string) error {
	table := "users"
	condition := "username = $1"

	return DeleteRow(table, condition, username)
}

func AddUserLanguage(username string, languages []string) error {
	err := AddArrayAttribute("users", "username", username, "languages", languages)
	if err != nil {
		log.Printf("Error adding languages for user %s: %v\n", username, err)
		return err
	}

	return nil
}

func RemoveUserLanguage(username string, languages []string) error {
	err := RemoveArrayAttribute("users", "username", username, "languages", languages)
	if err != nil {
		log.Printf("Error removing languages for user %s: %v\n", username, err)
		return err
	}

	return nil
}

func AddUserTag(username string, tags []string) error {
	err := AddArrayAttribute("users", "username", username, "tags", tags)
	if err != nil {
		log.Printf("Error adding tags for user %s: %v\n", username, err)
		return err
	}

	return nil
}

func RemoveUserTag(username string, tags []string) error {
	err := RemoveArrayAttribute("users", "username", username, "tags", tags)
	if err != nil {
		log.Printf("Error removing tags for user %s: %v\n", username, err)
		return err
	}

	return nil
}

func AddUserBoard(username string, boardId int) error {
	err := AddArrayAttribute("users", "username", username, "boards", IntsToStrings([]int{boardId}))
	if err != nil {
		log.Printf("Error adding board for user %s: %v\n", username, err)
		return err
	}

	return nil
}

func RemoveUserBoard(username string, boardId int) error {
	err := RemoveArrayAttribute("users", "username", username, "boards", IntsToStrings([]int{boardId}))
	if err != nil {
		log.Printf("Error removing board for user %s: %v\n", username, err)
		return err
	}

	return nil
}

func AddUserPost(username string, postId int) error {
	err := AddArrayAttribute("users", "username", username, "posts", IntsToStrings([]int{postId}))
	if err != nil {
		log.Printf("Error adding post for user %s: %v\n", username, err)
		return err
	}

	return nil
}

func RemoveUserPost(username string, postId int) error {
	err := RemoveArrayAttribute("users", "username", username, "posts", IntsToStrings([]int{postId}))
	if err != nil {
		log.Printf("Error removing post for user %s: %v\n", username, err)
		return err
	}

	return nil
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		username := r.URL.Query().Get("username")

		if username != "" {
			user, err := GetUser(username)
			if err != nil {
				http.Error(w, "User not found", http.StatusNotFound)
				log.Printf("Error retrieving user: %v\n", err)
				return
			}
			json.NewEncoder(w).Encode(user)
		} else {
			users, err := GetAllUsers()
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
		userId, err := AddUser(user)
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

		if err := DeleteUser(username); err != nil {
			http.Error(w, "Error deleting user", http.StatusInternalServerError)
			log.Printf("Error deleting user: %v\n", err)
			return
		}

		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
