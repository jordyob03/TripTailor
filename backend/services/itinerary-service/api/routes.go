package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	handlers "github.com/jordyob03/TripTailor/backend/services/itinerary-service/internal/handlers"
)

func RegisterRoutes(r *gin.Engine, dbConn *sql.DB) {
	r.POST("/itin-creation", handlers.CreateItin(dbConn))
}
