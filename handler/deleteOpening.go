package handler

import (
	"net/http"

	"github.com/dandevweb/gopportunities/schema"
	"github.com/gin-gonic/gin"
)

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
