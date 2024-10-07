package api

import (
	"backend/auth-service/internal/handlers"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, dbConn *sql.DB) {
	r.POST("/signup", handlers.SignUp(dbConn))
	r.POST("/signin", handlers.SignIn(dbConn))
}
