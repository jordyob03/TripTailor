package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/jordyob03/TripTailor/backend/services/auth-service/internal/db"
	utils "github.com/jordyob03/TripTailor/backend/services/auth-service/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignInRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
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
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Cant access user"})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
			return
		}

		token, err := utils.GenerateJWT(user.UserId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		fmt.Print(user.Username)
		c.JSON(http.StatusOK, gin.H{
			"message":  "Signin successful",
			"userId":   user.UserId,
			"token":    token,
			"username": user.Username,
		})
	}
}
