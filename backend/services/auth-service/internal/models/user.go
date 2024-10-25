package models

import "time"

type User struct {
	UserId       int       `json:"userId"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	DateOfBirth  time.Time `json:"dateOfBirth"`
	Name         string    `json:"name"`
	Country      string    `json:"country"`
	Languages    []string  `json:"languages"`
	Tags         []string  `json:"tags"`
	Boards       []string  `json:"boards"`
	Posts        []string  `json:"posts"`
	ProfileImage int       `json:"profileImage"`
	CoverImage   int       `json:"coverImage"`
}

func NewUser(
	userId int,
	username string,
	email string,
	password string,
	dateOfBirth time.Time,
	name string,
	country string,
	languages []string,
	tags []string,
	boards []string,
	posts []string,
	profileImage int,
	coverImage int,
) User {
	return User{
		UserId:       userId,
		Username:     username,
		Email:        email,
		Password:     password,
		DateOfBirth:  dateOfBirth,
		Name:         name,
		Country:      country,
		Languages:    languages,
		Tags:         tags,
		Boards:       boards,
		Posts:        posts,
		ProfileImage: profileImage,
		CoverImage:   coverImage,
	}
}
