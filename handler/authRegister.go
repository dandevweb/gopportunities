package handler

import (
	"net/http"

	"github.com/dandevweb/gopportunities/helpers"
	"github.com/dandevweb/gopportunities/model"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Register
// @Description Register
// @Tags auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Request body"
// @Success 200 {object} RegisterResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /register [post]
func RegisterHandler(ctx *gin.Context) {
	request := RegisterRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("error validating opening: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := helpers.HashPassword(request.Password)

	if err != nil {
		logger.Errorf("error hashing password: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error hashing password")
		return
	}

	user := model.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("error creating user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "user created", user)
}
