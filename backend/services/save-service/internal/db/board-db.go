package DBmodels

import (
	"database/sql"
	"fmt"
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

func AddBoard(DB *sql.DB, board Board) (int, error) {
	insertBoardSQL := `
    INSERT INTO boards (name, description, username, posts, tags) 
    VALUES ($1, $2, $3, $4, $5) RETURNING boardId;`

	if board.Posts == nil {
		board.Posts = []string{}
	}

	if board.Tags == nil {
		board.Tags = []string{}
	}

	var boardID int
	err := DB.QueryRow(
		insertBoardSQL, board.Name, board.Description,
		board.Username, pq.Array(board.Posts),
		pq.Array(board.Tags)).Scan(&boardID)
	if err != nil {
		return 0, fmt.Errorf("error adding board: %v", err)
	}

	AddUserBoard(DB, board.Username, boardID)

	log.Printf("Board added successfully with ID: %d\n", boardID)
	return boardID, AddUserBoard(DB, board.Username, boardID)
}

func RemoveBoard(DB *sql.DB, boardId int) error {
	boardData, err := GetBoard(DB, boardId)
	if err != nil {
		log.Printf("Error retrieving board with ID %d: %v\n", boardId, err)
		return err
	}

	boardPosts, err := StringsToInts(boardData.Posts)
	if err != nil {
		log.Printf("Error converting board posts to integers: %v\n", err)
		return err
	}

	for _, postId := range boardPosts {
		err = RemovePostBoard(DB, postId, boardId)
		if err != nil {
			log.Printf("Error removing post %d from board %d: %v\n", postId, boardId, err)
			return err
		}
	}

	err = RemoveUserBoard(DB, boardData.Username, boardId)
	if err != nil {
		log.Printf("Error removing board %d from user %s: %v\n", boardId, boardData.Username, err)
		return err
	}

	log.Println("Board removed successfully.")
	return nil
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

func UpdateBoardName(DB *sql.DB, boardId int, name string) error {
	err := UpdateAttribute(DB, "boards", "boardId", boardId, "name", name)
	if err != nil {
		log.Printf("Error updating board name: %v\n", err)
		return err
	}

	log.Println("Board name updated successfully.")
	return nil
}

func UpdateBoardDescription(DB *sql.DB, boardId int, description string) error {
	err := UpdateAttribute(DB, "boards", "boardId", boardId, "description", description)
	if err != nil {
		log.Printf("Error updating board description: %v\n", err)
		return err
	}

	log.Println("Board description updated successfully.")
	return nil
}

func UpdateBoardCreationDate(DB *sql.DB, boardId int, creationDate time.Time) error {
	err := UpdateAttribute(DB, "boards", "boardId", boardId, "creationDate", creationDate)
	if err != nil {
		log.Printf("Error updating board creation date: %v\n", err)
		return err
	}

	log.Println("Board creation date updated successfully.")
	return nil
}

func AddBoardPost(DB *sql.DB, boardId int, postId int, recursive bool) error {
	err := AddArrayAttribute(DB, "boards", "boardId", boardId, "posts", IntsToStrings([]int{postId}))
	if err != nil {
		log.Printf("Error adding post for board %d: %v\n", boardId, err)
		return err
	}

	log.Printf("Post %d added to board %d successfully.\n", postId, boardId)

	if recursive {
		return AddPostBoard(DB, postId, boardId, false)
	}

	return nil
}

func RemoveBoardPost(DB *sql.DB, boardId int, postId int) error {
	err := RemoveArrayAttribute(DB, "boards", "boardId", boardId, "posts", IntsToStrings([]int{postId}))
	if err != nil {
		log.Printf("Error removing post for board %d: %v\n", boardId, err)
		return err
	}

	return nil
}

func AddBoardTag(DB *sql.DB, boardId int, tag string) error {
	err := AddArrayAttribute(DB, "boards", "boardId", boardId, "tags", []string{tag})
	if err != nil {
		log.Printf("Error adding tag for board %d: %v\n", boardId, err)
		return err
	}

	return nil
}

func RemoveBoardTag(DB *sql.DB, boardId int, tag string) error {
	err := RemoveArrayAttribute(DB, "boards", "boardId", boardId, "tags", []string{tag})
	if err != nil {
		log.Printf("Error removing tag for board %d: %v\n", boardId, err)
		return err
	}

	return nil
}
