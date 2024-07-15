package handler

import (
	"net/http"

	"github.com/dandevweb/gopportunities/model"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Update opening
// @Description Update an opening
// @Tags openings
// @Accept json
// @Produce json
// @Param id path string true "Opening ID"
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} UpdateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings/{id} [put]
// @Security BearerAuth
func UpdateOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "id is required")
		return
	}
	opening := model.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("error finding opening: %v", err.Error())
		sendError(ctx, http.StatusNotFound, err.Error())
		return
	}

	if err := ctx.BindJSON(&opening); err != nil {
		logger.Errorf("error binding opening: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error saving opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "Opening updated", opening)
}
