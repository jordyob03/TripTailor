package handlers

import (
	db "backend/auth-service/internal/db"
	"backend/auth-service/utils"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignInRequest struct {
	UserId      int       `json:"userId"`
	Username    string    `json:"username" binding:"required"`
	Email       string    `json:"email" binding:"required,email"`
	Password    string    `json:"password" binding:"required,min=6"`
	DateOfBirth time.Time `json:"dateOfBirth" binding:"required"`
	Name        string    `json:"name"`
	Country     string    `json:"country"`
	Languages   []string  `json:"languages"`
	Tags        []string  `json:"tags"`
	Boards      []string  `json:"boards"`
	Posts       []string  `json:"posts"`
}

func SignIn(DB *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req SignInRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := db.GetUser(DB, req.Username)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		token, err := utils.GenerateJWT(user.UserId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Signin successful",
			"userId":  user.UserId,
			"token":   token,
		})
	}
}
