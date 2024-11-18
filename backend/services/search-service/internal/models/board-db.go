package models

import (
	"database/sql"
	"log"
	"time"

	"github.com/lib/pq"
)

type Board struct {
	BoardId      int       `json:"boardId"`
	Name         string    `json:"name"`
	CreationDate time.Time `json:"dateOfCreation"`
	Description  string    `json:"description"`
	Username     string    `json:"username"`
	Posts        []string  `json:"posts"`
	Tags         []string  `json:"tags"`
}

func CreateBoardTable(DB *sql.DB) error {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS boards (
        boardId SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        creationDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        description TEXT,
        username VARCHAR(255) REFERENCES users(username),
        posts TEXT[],
        tags TEXT[]
    );`

	return CreateTable(DB, createTableSQL)
}

func GetBoard(DB *sql.DB, boardId int) (Board, error) {
	var board Board

	query := `
	SELECT boardId, name, creationDate, description, username, posts, tags 
	FROM boards 
	WHERE boardId = $1;
	`

	err := DB.QueryRow(query, boardId).Scan(
		&board.BoardId,
		&board.Name,
		&board.CreationDate,
		&board.Description,
		&board.Username,
		pq.Array(&board.Posts),
		pq.Array(&board.Tags),
	)
	if err != nil {
		log.Printf("Error fetching board with ID %d: %v\n", boardId, err)
		return Board{}, err
	}

	if board.Posts == nil {
		board.Posts = []string{}
	}

	if board.Tags == nil {
		board.Tags = []string{}
	}

	log.Printf("Board with ID %d retrieved successfully.\n", boardId)
	return board, nil
}
