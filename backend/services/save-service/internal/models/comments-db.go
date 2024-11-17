package models

import (
	"database/sql"
)

type Comment struct {
	CommentId   int    `json:"commentId"`
	PostId      int    `json:"postId"`
	ItineraryId int    `json:"itineraryId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Likes       int    `json:"likes"`
	Username    string `json:"username"`
}

func CreateCommentTable(DB *sql.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS comments (
		commentId SERIAL PRIMARY KEY,
		postId INT,
		itineraryId INT,
		title TEXT NOT NULL,
		description TEXT,
		likes INT,
		username VARCHAR(255) REFERENCES users(username)
	);`

	return CreateTable(DB, createTableSQL)
}

func AddComment(DB *sql.DB, comment Comment) (int, error) {
	insertSQL := `
	INSERT INTO comments (commentId, postId, itineraryId, title, description, likes, username)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING commentId;
	`

	var commentId int
	err := DB.QueryRow(
		insertSQL,
		comment.CommentId,
		comment.PostId,
		comment.ItineraryId,
		comment.Title,
		comment.Description,
		comment.Likes,
		comment.Username,
	).Scan(&commentId)

	return commentId, err
}

func GetComment(DB *sql.DB, commentId int) (Comment, error) {
	getCommentSQL := `
	SELECT * FROM comments
	WHERE commentId = $1;`

	var comment Comment
	err := DB.QueryRow(getCommentSQL, commentId).Scan(
		&comment.CommentId,
		&comment.PostId,
		&comment.ItineraryId,
		&comment.Title,
		&comment.Description,
		&comment.Likes,
		&comment.Username,
	)

	return comment, err
}

func RemoveComment(DB *sql.DB, commentId int) error {
	removeCommentSQL := `
	DELETE FROM comments
	WHERE commentId = $1;`

	_, err := DB.Exec(removeCommentSQL, commentId)

	return err
}

func UpdateCommentTitle(DB *sql.DB, commentId int, title string) error {
	return UpdateAttribute(DB, "comments", "commentId", commentId, "title", title)
}

func UpdateCommentDescription(DB *sql.DB, commentId int, description string) error {
	return UpdateAttribute(DB, "comments", "commentId", commentId, "description", description)
}

func UpdateCommentLikes(DB *sql.DB, commentId int, likes int) error {
	return UpdateAttribute(DB, "comments", "commentId", commentId, "likes", likes)
}
