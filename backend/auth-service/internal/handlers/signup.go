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

type SignUpRequest struct {
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

func SignUp(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req SignUpRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if _, err := db.GetUser(dbConn, req.Username); err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		req.Password = string(hashedPassword)

		userId, err := db.AddUser(dbConn, db.User(req))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		token, err := utils.GenerateJWT(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User created successfully",
			"userId":  userId,
			"token":   token,
		})
	}
}
