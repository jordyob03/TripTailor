// package createprofile

// import (
// 	"backend/profile-service/internal/handlers"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// // Handler for creating a profile.
// func CreateProfile() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var req handlers.CreateProfileRequest

// 		// Bind and validate the request body.
// 		if err := c.ShouldBindJSON(&req); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// Ensure at least 3 tags are provided.
// 		tags := handlers.ParseTags(req.Tags)
// 		if len(tags) < 3 {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "At least 3 tags must be selected"})
// 			return
// 		}

// 		// Check if the profile already exists.
// 		if _, exists := handlers.Profiles[req.Username]; exists {
// 			c.JSON(http.StatusConflict, gin.H{"error": "Profile already exists"})
// 			return
// 		}

// 		// Store the new profile.
// 		handlers.Profiles[req.Username] = req

// 		// Respond with the created profile.
// 		c.JSON(http.StatusCreated, gin.H{
// 			"message": "Profile created successfully",
// 			"profile": req,
// 		})
// 	}
// }
