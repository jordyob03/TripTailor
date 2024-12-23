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
	CreationDate time.Time `json:"dateOfCreation"`
	Username     string    `json:"username"`
	Boards       []string  `json:"boards"`
	Likes        int       `json:"likes"`
	Comments     []string  `json:"comments"`
}

func CreatePostTable(DB *sql.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS posts (
		postId SERIAL PRIMARY KEY,
		itineraryId INT,
		creationDate TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		username VARCHAR(255) REFERENCES users(username),
		boards TEXT[],
		likes INT DEFAULT 0,
		comments TEXT[]
	);`

	return CreateTable(DB, createTableSQL)
}

func GetPost(DB *sql.DB, postId int) (Post, error) {
	var post Post

	query := `
	SELECT postId, itineraryId, creationDate, username, boards, likes, comments
	FROM posts
	WHERE postId = $1;
	`

	err := DB.QueryRow(query, postId).Scan(
		&post.PostId,
		&post.ItineraryId,
		&post.CreationDate,
		&post.Username,
		pq.Array(&post.Boards),
		&post.Likes,
		pq.Array(&post.Comments),
	)

	if err != nil {
		log.Printf("Error retrieving post with ID %d: %v\n", postId, err)
		return Post{}, fmt.Errorf("failed to retrieve post: %w", err)
	}

	if post.Boards == nil {
		post.Boards = []string{}
	}

	if post.Comments == nil {
		post.Comments = []string{}
	}

	log.Printf("Post with ID %d successfully retrieved.\n", postId)
	return post, nil
}
