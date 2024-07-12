package handler

import (
	"net/http"

	"github.com/dandevweb/gopportunities/schema"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete  opening
// @Description Delete an opening
// @Tags openings
// @Accept json
// @Produce json
// @Param id path string true "Opening ID"
// @Success 204 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings/{id} [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "id is required")
		return
	}
	opening := schema.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("error finding opening: %v", err.Error())
		sendError(ctx, http.StatusNotFound, err.Error())
		return
	}

	if err := db.Delete(&opening, id).Error; err != nil {
		logger.Errorf("error deleting opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendNoContent(ctx)
}
