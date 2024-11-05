package main

import (
	"database/sql"
	"log"
)

func main() {
	connStr := "postgres://postgres:password@db:5432/database?sslmode=disable"
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer dbConn.Close()
}
