package handlers

import (
	"strings"
)

// Shared Profiles map used across both create and update handlers.
var Profiles = make(map[string]CreateProfileRequest)

// CreateProfileRequest struct shared between create and update operations.
type CreateProfileRequest struct {
	Language string `json:"language" binding:"required"`
	Country  string `json:"country" binding:"required"`
	Tags     string `json:"tags" binding:"required"`
	Name     string `json:"name"`
	Username string `json:"username" binding:"required"`
}

// ParseTags trims and splits comma-separated tags into a slice.
func ParseTags(tags string) []string {
	parsedTags := strings.Split(tags, ",")
	for i, tag := range parsedTags {
		parsedTags[i] = strings.TrimSpace(tag)
	}
	return parsedTags
}
