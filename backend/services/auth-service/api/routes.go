package api

import (
	"database/sql"
	"github.com/jordyob03/TripTailor/backend/services/auth-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, dbConn *sql.DB) {
	r.POST("/signup", handlers.SignUp(dbConn))
	r.POST("/signin", handlers.SignIn(dbConn))
}
