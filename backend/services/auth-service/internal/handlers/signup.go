package handlers

import (
	"database/sql"
	"github.com/jordyob03/TripTailor/backend/services/auth-service/internal/db"
	"github.com/jordyob03/TripTailor/backend/services/auth-service/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignUpRequest struct {
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	DateOfBirth string `json:"dateOfBirth" binding:"required"`
}

func SignUp(dbConn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req SignUpRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		dateOfBirth, err := time.Parse("2006-01-02", req.DateOfBirth)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
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

		userId, err := db.AddUser(dbConn, req.Username, req.Email, string(hashedPassword), dateOfBirth)
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
