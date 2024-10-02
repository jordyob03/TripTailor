package DBmodels

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	UserID      int       `json:"userId"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

func CreateUserTable() error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		userId SERIAL PRIMARY KEY,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		dateOfBirth DATE NOT NULL
	);`
	return CreateTable(createTableSQL)
}

func AddUser(username, email, password string, dateOfBirth time.Time) (int, error) {
	insertUserSQL := `
	INSERT INTO users (username, email, password, dateOfBirth)
	VALUES ($1, $2, $3, $4)
	RETURNING userId;`

	var userId int
	err := DB.QueryRow(insertUserSQL, username, email, password, dateOfBirth).Scan(&userId)
	if err != nil {
		fmt.Println("Error adding user:", err)
		return 0, err
	}

	return userId, nil
}

func GetUser(userID string) (User, error) {
	query := `SELECT userId, username, email, password, dateOfBirth FROM users WHERE userId = $1`
	row := DB.QueryRow(query, userID)

	var user User
	err := row.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.DateOfBirth)
	if err != nil {
		return User{}, fmt.Errorf("no user found with userId: %s", userID)
	}

	return user, nil
}

func GetAllUsers() ([]User, error) {
	rows, err := DB.Query("SELECT userId, username, email, password, dateOfBirth FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.DateOfBirth); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func UpdateUser(userID string, data map[string]interface{}) error {
	if len(data) == 0 {
		return fmt.Errorf("no fields to update for user with userId: %s", userID)
	}

	table := "users"
	condition := "userId = $1"

	return UpdateRow(table, data, condition, userID)
}

func DeleteUser(userID string) error {
	table := "users"
	condition := "userId = $1"

	return DeleteRow(table, condition, userID)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		userID := r.URL.Query().Get("userId")

		if userID != "" {
			id, err := strconv.Atoi(userID)
			if err != nil {
				http.Error(w, "Invalid userId", http.StatusBadRequest)
				fmt.Println(err)
				return
			}
			user, err := GetUser(strconv.Itoa(id))
			if err != nil {
				http.Error(w, "User not found", http.StatusNotFound)
				fmt.Println(err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
		} else {
			users, err := GetAllUsers()
			if err != nil {
				http.Error(w, "Error retrieving users", http.StatusInternalServerError)
				fmt.Println(err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(users)
		}

	case "POST":
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		userId, err := AddUser(user.Username, user.Email, user.Password, user.DateOfBirth)
		if err != nil {
			http.Error(w, "Error adding user", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int{"userId": userId})

	case "PUT":
		userID := r.URL.Query().Get("userId")
		if userID == "" {
			http.Error(w, "userId is required", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(userID)
		if err != nil {
			http.Error(w, "Invalid userId", http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		data := map[string]interface{}{}
		if user.Username != "" {
			data["username"] = user.Username
		}
		if user.Password != "" {
			data["password"] = user.Password
		}
		if !user.DateOfBirth.IsZero() {
			data["dateOfBirth"] = user.DateOfBirth
		}

		if err := UpdateUser(strconv.Itoa(id), data); err != nil {
			http.Error(w, "Error updating user", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)

	case "DELETE":
		userID := r.URL.Query().Get("userId")
		if userID == "" {
			http.Error(w, "userId is required", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(userID)
		if err != nil {
			http.Error(w, "Invalid userId", http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		if err := DeleteUser(strconv.Itoa(id)); err != nil {
			http.Error(w, "Error deleting user", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
