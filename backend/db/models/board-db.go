package DBmodels

import (
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

func CreateBoardTable() error {
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

	_, err := DB.Exec(createTableSQL)
	if err != nil {
		log.Printf("Error creating boards table: %v\n", err)
		return err
	}

	log.Println("Boards table created successfully.")
	return nil
}

func AddBoard(board Board) error {
	insertBoardSQL := `
    INSERT INTO boards (name, description, username, posts, tags) 
    VALUES ($1, $2, $3, $4, $5) RETURNING board_id;`

	var boardID int
	err := DB.QueryRow(
		insertBoardSQL, board.Name, board.Description,
		board.Username, pq.Array(board.Posts),
		pq.Array(board.Tags)).Scan(&boardID)
	if err != nil {
		log.Printf("Error adding board: %v\n", err)
		return err
	}

	AddUserBoard(board.Username, boardID)

	log.Printf("Board added successfully with ID: %d\n", boardID)
	return nil
}

func RemoveBoard(board Board) error {
	deleteBoardSQL := `DELETE FROM boards WHERE board_id = $1;`

	_, err := DB.Exec(deleteBoardSQL, board.BoardId)
	if err != nil {
		log.Printf("Error removing board: %v\n", err)
		return err
	}

	RemoveUserBoard(board.Username, board.BoardId)

	log.Println("Board removed successfully.")
	return nil
}

func GetBoard(boardId int) (Board, error) {
	var board Board

	query := `
	SELECT board_id, name, creation_date, description, username, posts, tags 
	FROM boards 
	WHERE board_id = $1;
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

	log.Printf("Board with ID %d retrieved successfully.\n", boardId)
	return board, nil
}

func AddBoardPost(boardId string, postId int) error {
	err := AddArrayAttribute("boards", "boardId", boardId, "posts", []int{postId})
	if err != nil {
		log.Printf("Error adding post for board %s: %v\n", boardId, err)
		return err
	}

	return nil
}

func RemoveBoardPost(boardId string, postId int) error {
	err := RemoveArrayAttribute("boards", "boardId", boardId, "posts", []int{postId})
	if err != nil {
		log.Printf("Error removing post for board %s: %v\n", boardId, err)
		return err
	}

	return nil
}

func AddBoardTag(boardId string, tag string) error {
	err := AddArrayAttribute("boards", "boardId", boardId, "tags", []string{tag})
	if err != nil {
		log.Printf("Error adding tag for board %s: %v\n", boardId, err)
		return err
	}

	return nil
}

func RemoveBoardTag(boardId string, tag string) error {
	err := RemoveArrayAttribute("boards", "boardId", boardId, "tags", []string{tag})
	if err != nil {
		log.Printf("Error removing tag for board %s: %v\n", boardId, err)
		return err
	}

	return nil
}
