package handler

import (
	"net/http"

	"github.com/dandevweb/gopportunities/schema"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schema.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		logger.Errorf("error finding openings: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "listing openings", openings)
}
