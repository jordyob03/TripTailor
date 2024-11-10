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

func AddPost(DB *sql.DB, post Post) (int, error) {
	insertSQL := `
	INSERT INTO posts (itineraryId, creationDate, username, boards, likes, comments)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING postId;
	`

	var postId int
	err := DB.QueryRow(
		insertSQL,
		post.ItineraryId,
		post.CreationDate,
		post.Username,
		pq.Array(post.Boards),
		post.Likes,
		pq.Array(post.Comments),
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
	SELECT username, boards, itineraryId
	FROM posts
	WHERE postId = $1;
	`

	var username string
	var boardStringIds []string
	var itineraryId int
	err := DB.QueryRow(getBoardsSQL, postId).Scan(&username, pq.Array(&boardStringIds), &itineraryId)
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

	err = RemoveItinerary(DB, itineraryId)
	if err != nil {
		log.Printf("Error removing itinerary %d: %v\n", itineraryId, err)
		return err
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

func UpdatePostItineraryId(DB *sql.DB, postId int, itineraryId int) error {
	return UpdateAttribute(DB, "posts", "postId", postId, "itineraryId", itineraryId)
}

func AddPostBoard(DB *sql.DB, postId int, board int, recursive bool) error {
	return AddArrayAttribute(DB, "posts", "postId", postId, "boards", IntsToStrings([]int{board}))
}

func RemovePostBoard(DB *sql.DB, postId int, board int) error {
	return RemoveArrayAttribute(DB, "posts", "postId", postId, "boards", IntsToStrings([]int{board}))
}

func UpdatePostCreationDate(DB *sql.DB, postId int, creationDate time.Time) error {
	return UpdateAttribute(DB, "posts", "postId", postId, "creationDate", creationDate)
}

func UpdatePostLikes(DB *sql.DB, postId int, likes int) error {
	return UpdateAttribute(DB, "posts", "postId", postId, "likes", likes)
}

func AddPostComment(DB *sql.DB, postId int, commentId int) error {
	return AddArrayAttribute(DB, "posts", "postId", postId, "comments", IntsToStrings([]int{commentId}))
}

func RemovePostComment(DB *sql.DB, postId int, commentId int) error {
	return RemoveArrayAttribute(DB, "posts", "postId", postId, "comments", IntsToStrings([]int{commentId}))
}
