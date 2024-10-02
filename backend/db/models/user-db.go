package DBmodels

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUserTable() error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`
	return CreateTable(createTableSQL)
}

func AddUser(email, password string) error {
	userData := map[string]interface{}{
		"email":    email,
		"password": password,
	}

	err := AddRow("users", userData)

	if err != nil {
		fmt.Println("Error adding user:", err)
	}

	return nil
}

func GetUser(email string) (map[string]interface{}, error) {
	table := "users"
	condition := "email = $1"

	rows, err := GetRows(table, condition, email)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, fmt.Errorf("no user found with email: %s", email)
	}

	return rows[0], nil
}

func GetAllUsers() ([]User, error) {
	rows, err := DB.Query("SELECT email, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func UpdateUser(email string, data map[string]interface{}) error {
	if len(data) == 0 {
		return fmt.Errorf("no fields to update for user with email: %s", email)
	}

	table := "users"
	condition := "email = $1"

	return UpdateRow(table, data, condition, email)
}

func DeleteUser(email string) error {
	table := "users"
	condition := "email = $1"

	return DeleteRow(table, condition, email)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		email := r.URL.Query().Get("email")
		if email != "" {
			user, err := GetUser(email)
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
		if err := AddUser(user.Email, user.Password); err != nil {
			http.Error(w, "Error adding user", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusCreated)

	case "PUT":
		email := r.URL.Query().Get("email")
		if email == "" {
			http.Error(w, "Email is required", http.StatusBadRequest)
			return
		}

		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		if err := UpdateUser(email, map[string]interface{}{
			"password": user.Password,
		}); err != nil {
			http.Error(w, "Error updating user", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)

	case "DELETE":
		email := r.URL.Query().Get("email")
		if email == "" {
			http.Error(w, "Email is required", http.StatusBadRequest)
			return
		}

		if err := DeleteUser(email); err != nil {
			http.Error(w, "Error deleting user", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
