package models

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
