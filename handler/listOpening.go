package handler

import (
	"net/http"

	"github.com/dandevweb/gopportunities/model"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List openings
// @Description List all openings
// @Tags openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
// @Security BearerAuth
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []model.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		logger.Errorf("error finding openings: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "listing openings", openings)
}
