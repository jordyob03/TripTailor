package DBmodels

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
)

type Post struct {
	PostId       int       `json:"postId"`
	ItineraryId  int       `json:"itineraryId"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CreationDate time.Time `json:"dateOfCreation"`
	Username     string    `json:"username"`
	Tags         []string  `json:"tags"`
	Boards       []string  `json:"boards"`
	PostImages   []string  `json:"postImages"`
}

func CreatePostTable(DB *sql.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS posts (
		postId SERIAL PRIMARY KEY,
		itineraryId INT,
		title TEXT NOT NULL,
		imageLink TEXT,
		description TEXT,
		creationDate TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		username VARCHAR(255) REFERENCES users(username),
		tags TEXT[],
		boards TEXT[],
		postImages TEXT[]
	);`

	return CreateTable(DB, createTableSQL)
}

func AddPost(DB *sql.DB, post Post) (int, error) {
	insertSQL := `
	INSERT INTO posts (itineraryId, title, description, creationDate, username, tags, boards, postImages)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING postId;
	`

	var postId int
	err := DB.QueryRow(
		insertSQL,
		post.ItineraryId,
		post.Title,
		post.Description,
		post.CreationDate,
		post.Username,
		pq.Array(post.Tags),
		pq.Array(post.Boards),
		pq.Array(post.PostImages),
	).Scan(&postId)

	if err != nil {
		log.Printf("Error adding post: %v\n", err)
		return 0, fmt.Errorf("failed to add post: %w", err)
	}

	log.Printf("Post with ID %d successfully added.\n", postId)
	return postId, nil
}

func RemovePost(DB *sql.DB, postId int) error {
	getBoardsSQL := `
	SELECT username, boards
	FROM posts
	WHERE postId = $1;
	`

	var username string
	var boardStringIds []string
	err := DB.QueryRow(getBoardsSQL, postId).Scan(&username, pq.Array(&boardStringIds))
	if err == sql.ErrNoRows {
		log.Printf("No boards found with postID %d.\n", postId)
		return nil
	} else if err != nil {
		log.Printf("Error retrieving boards for post ID %d: %v\n", postId, err)
		return err
	}

	boardIds, err := StringsToInts(boardStringIds)
	if err != nil {
		log.Printf("Error converting board IDs to integers: %v\n", err)
		return err
	}

	for _, board := range boardIds {
		err = RemoveBoardPost(DB, board, postId)
		if err != nil {
			log.Printf("Error removing post %d from board %d: %v\n", postId, board, err)
			return err
		}
	}

	RemoveUserPost(DB, username, postId)

	deleteSQL := `
	DELETE FROM posts
	WHERE postId = $1;
	`

	result, err := DB.Exec(deleteSQL, postId)
	if err != nil {
		log.Printf("Error removing post with ID %d: %v\n", postId, err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		log.Printf("No post found with ID %d.\n", postId)
		return err
	}

	log.Printf("Post with ID %d successfully removed.\n", postId)
	return nil
}

func GetPost(DB *sql.DB, postId int) (Post, error) {
	var post Post

	query := `
	SELECT postId, itineraryId, title, description, creationDate, username, tags, boards, postImages
	FROM posts
	WHERE postId = $1;
	`

	err := DB.QueryRow(query, postId).Scan(
		&post.PostId,
		&post.ItineraryId,
		&post.Title,
		&post.Description,
		&post.CreationDate,
		&post.Username,
		pq.Array(&post.Tags),
		pq.Array(&post.Boards),
		pq.Array(&post.PostImages),
	)

	if err != nil {
		log.Printf("Error retrieving post with ID %d: %v\n", postId, err)
		return Post{}, fmt.Errorf("failed to retrieve post: %w", err)
	}

	log.Printf("Post with ID %d successfully retrieved.\n", postId)
	return post, nil
}

func AddPostTag(DB *sql.DB, postId int, tag string) error {
	err := AddArrayAttribute(DB, "posts", "postId", postId, "tags", []string{tag})
	if err != nil {
		log.Printf("Error adding tag to post %d: %v\n", postId, err)
		return fmt.Errorf("failed to add tag to post: %w", err)
	}

	log.Printf("Tag added to post %d successfully.\n", postId)
	return nil
}

func RemovePostTag(DB *sql.DB, postId int, tag string) error {
	err := RemoveArrayAttribute(DB, "posts", "postId", postId, "tags", []string{tag})
	if err != nil {
		log.Printf("Error removing tag from post %d: %v\n", postId, err)
		return fmt.Errorf("failed to remove tag from post: %w", err)
	}

	log.Printf("Tag removed from post %d successfully.\n", postId)
	return nil
}

func AddPostBoard(DB *sql.DB, postId int, board int, recursive bool) error {
	if !recursive {
		return nil
	}

	err := AddArrayAttribute(DB, "posts", "postId", postId, "boards", IntsToStrings([]int{board}))
	if err != nil {
		log.Printf("Error adding board to post %d: %v\n", postId, err)
		return fmt.Errorf("failed to add board to post: %w", err)
	}

	err = AddBoardPost(DB, board, postId, false)
	if err != nil {
		log.Printf("Error adding post %d to board %d: %v\n", postId, board, err)
		return fmt.Errorf("failed to add post to board: %w", err)
	}

	log.Printf("Board added to post %d successfully.\n", postId)
	return nil
}

func RemovePostBoard(DB *sql.DB, postId int, board int) error {
	err := RemoveArrayAttribute(DB, "posts", "postId", postId, "boards", IntsToStrings([]int{board}))
	if err != nil {
		log.Printf("Error removing board from post %d: %v\n", postId, err)
		return fmt.Errorf("failed to remove board from post: %w", err)
	}

	log.Printf("Board removed from post %d successfully.\n", postId)
	return nil
}

func UpdatePostDescription(DB *sql.DB, postId int, description string) error {
	err := UpdateAttribute(DB, "posts", "postId", postId, "description", description)
	if err != nil {
		log.Printf("Error updating description for post %d: %v\n", postId, err)
		return fmt.Errorf("failed to update description for post: %w", err)
	}

	log.Printf("Description updated for post %d successfully.\n", postId)
	return nil
}

func UpdatePostTitle(DB *sql.DB, postId int, title string) error {
	err := UpdateAttribute(DB, "posts", "postId", postId, "title", title)
	if err != nil {
		log.Printf("Error updating title for post %d: %v\n", postId, err)
		return fmt.Errorf("failed to update title for post: %w", err)
	}

	log.Printf("Title updated for post %d successfully.\n", postId)
	return nil
}

func UpdatePostCreationDate(DB *sql.DB, postId int, creationDate time.Time) error {
	err := UpdateAttribute(DB, "posts", "postId", postId, "creationDate", creationDate)
	if err != nil {
		log.Printf("Error updating creation date for post %d: %v\n", postId, err)
		return fmt.Errorf("failed to update creation date for post: %w", err)
	}

	log.Printf("Creation date updated for post %d successfully.\n", postId)
	return nil
}

func AddPostImage(DB *sql.DB, postId int, image string) error {
	err := AddArrayAttribute(DB, "posts", "postId", postId, "postImages", []string{image})
	if err != nil {
		log.Printf("Error adding image to post %d: %v\n", postId, err)
		return fmt.Errorf("failed to add image to post: %w", err)
	}

	log.Printf("Image added to post %d successfully.\n", postId)
	return nil
}

func RemovePostImage(DB *sql.DB, postId int, image string) error {
	err := RemoveArrayAttribute(DB, "posts", "postId", postId, "postImages", []string{image})
	if err != nil {
		log.Printf("Error removing image from post %d: %v\n", postId, err)
		return fmt.Errorf("failed to remove image from post: %w", err)
	}

	log.Printf("Image removed from post %d successfully.\n", postId)
	return nil
}

func GetAllPostImages(DB *sql.DB, postId int) ([]string, error) {
	query := `SELECT postImages FROM posts WHERE postId = $1;`

	var postImages []string
	err := DB.QueryRow(query, postId).Scan(pq.Array(&postImages))
	if err != nil {
		log.Printf("Error retrieving images for post %d: %v\n", postId, err)
		return nil, fmt.Errorf("failed to retrieve images for post: %w", err)
	}

	log.Printf("Images for post %d successfully retrieved.\n", postId)
	return postImages, nil
}
