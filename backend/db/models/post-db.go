package DBmodels

import (
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
)

type Post struct {
	PostId       int       `json:"postId"`
	ItineraryId  int       `json:"itineraryId"`
	Title        string    `json:"title"`
	ImageLink    string    `json:"imageLink"`
	Description  string    `json:"description"`
	CreationDate time.Time `json:"dateOfCreation"`
	Username     string    `json:"username"`
	Tags         []string  `json:"tags"`
}

func CreatePostTable() error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS posts (
		postId SERIAL PRIMARY KEY,
		itineraryId INT REFERENCES itineraries(itineraryId),
		title TEXT NOT NULL,
		imageLink TEXT,
		description TEXT,
		creationDate TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		username VARCHAR(255) REFERENCES users(username),,
		tags TEXT[]
	);`

	return CreateTable(createTableSQL)
}

func AddPost(post Post) (int, error) {
	insertSQL := `
	INSERT INTO posts (itineraryId, title, imageLink, description, creationDate, username, tags)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING postId;
	`

	var postId int
	err := DB.QueryRow(
		insertSQL,
		post.ItineraryId,
		post.Title,
		post.ImageLink,
		post.Description,
		post.CreationDate,
		post.Username,
		pq.Array(post.Tags),
	).Scan(&postId)

	if err != nil {
		log.Printf("Error adding post: %v\n", err)
		return 0, fmt.Errorf("failed to add post: %w", err)
	}

	log.Printf("Post with ID %d successfully added.\n", postId)
	return postId, nil
}

func RemovePost(postId int) error {
	deleteSQL := `
	DELETE FROM posts
	WHERE postId = $1;
	`

	result, err := DB.Exec(deleteSQL, postId)
	if err != nil {
		log.Printf("Error removing post with ID %d: %v\n", postId, err)
		return fmt.Errorf("failed to remove post: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		log.Printf("No post found with ID %d.\n", postId)
		return fmt.Errorf("post with ID %d not found", postId)
	}

	log.Printf("Post with ID %d successfully removed.\n", postId)
	return nil
}

func GetPost(postId int) (Post, error) {
	var post Post

	query := `
	SELECT postId, itineraryId, title, imageLink, description, creationDate, username, tags
	FROM posts
	WHERE postId = $1;
	`

	err := DB.QueryRow(query, postId).Scan(
		&post.PostId,
		&post.ItineraryId,
		&post.Title,
		&post.ImageLink,
		&post.Description,
		&post.CreationDate,
		&post.Username,
		pq.Array(&post.Tags),
	)

	if err != nil {
		log.Printf("Error retrieving post with ID %d: %v\n", postId, err)
		return Post{}, fmt.Errorf("failed to retrieve post: %w", err)
	}

	log.Printf("Post with ID %d successfully retrieved.\n", postId)
	return post, nil
}

func AddPostTag(postId int, tag string) error {
	err := AddArrayAttribute("posts", "postId", postId, "tags", []string{tag})
	if err != nil {
		log.Printf("Error adding tag to post %d: %v\n", postId, err)
		return fmt.Errorf("failed to add tag to post: %w", err)
	}

	log.Printf("Tag added to post %d successfully.\n", postId)
	return nil
}

func RemovePostTag(postId int, tag string) error {
	err := RemoveArrayAttribute("posts", "postId", postId, "tags", []string{tag})
	if err != nil {
		log.Printf("Error removing tag from post %d: %v\n", postId, err)
		return fmt.Errorf("failed to remove tag from post: %w", err)
	}

	log.Printf("Tag removed from post %d successfully.\n", postId)
	return nil
}

func AddPostImageLink(postId int, imageLink string) error {
	err := UpdateAttribute("posts", "postId", postId, "imageLink", imageLink)
	if err != nil {
		log.Printf("Error adding image link to post %d: %v\n", postId, err)
		return fmt.Errorf("failed to add image link to post: %w", err)
	}

	log.Printf("Image link added to post %d successfully.\n", postId)
	return nil
}

func UpdatePostDescription(postId int, description string) error {
	err := UpdateAttribute("posts", "postId", postId, "description", description)
	if err != nil {
		log.Printf("Error updating description for post %d: %v\n", postId, err)
		return fmt.Errorf("failed to update description for post: %w", err)
	}

	log.Printf("Description updated for post %d successfully.\n", postId)
	return nil
}

func UpdatePostTitle(postId int, title string) error {
	err := UpdateAttribute("posts", "postId", postId, "title", title)
	if err != nil {
		log.Printf("Error updating title for post %d: %v\n", postId, err)
		return fmt.Errorf("failed to update title for post: %w", err)
	}

	log.Printf("Title updated for post %d successfully.\n", postId)
	return nil
}

func UpdatePostCreationDate(postId int, creationDate time.Time) error {
	err := UpdateAttribute("posts", "postId", postId, "creationDate", creationDate)
	if err != nil {
		log.Printf("Error updating creation date for post %d: %v\n", postId, err)
		return fmt.Errorf("failed to update creation date for post: %w", err)
	}

	log.Printf("Creation date updated for post %d successfully.\n", postId)
	return nil
}
