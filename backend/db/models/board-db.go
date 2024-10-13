package DBmodels

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

func CreateBoardTable(db *sql.DB) error {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS boards (
        board_id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        description TEXT,
        username VARCHAR(255) REFERENCES users(username),
        posts TEXT[],
        tags TEXT[]
    );`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Printf("Error creating boards table: %v\n", err)
		return err
	}

	log.Println("Boards table created successfully.")
	return nil
}

func AddBoard(db *sql.DB, board Board) error {
	insertBoardSQL := `
    INSERT INTO boards (name, description, username, posts, tags) 
    VALUES ($1, $2, $3, $4, $5) RETURNING board_id;`

	var boardID int
	err := db.QueryRow(insertBoardSQL, board.Name, board.Description, board.Username, pq.Array(board.Posts), pq.Array(board.Tags)).Scan(&boardID)
	if err != nil {
		log.Printf("Error adding board: %v\n", err)
		return err
	}

	AddUserBoard(board.Username, boardID)

	log.Printf("Board added successfully with ID: %d\n", boardID)
	return nil
}

func RemoveBoard(db *sql.DB, board Board) error {
	deleteBoardSQL := `DELETE FROM boards WHERE board_id = $1;`

	_, err := db.Exec(deleteBoardSQL, board.BoardId)
	if err != nil {
		log.Printf("Error removing board: %v\n", err)
		return err
	}

	RemoveUserBoard(board.Username, board.BoardId)

	log.Println("Board removed successfully.")
	return nil
}
