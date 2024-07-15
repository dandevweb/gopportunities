package handler

import (
	"net/http"

	"github.com/dandevweb/gopportunities/config"
	"github.com/dandevweb/gopportunities/helpers"
	"github.com/dandevweb/gopportunities/model"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Login
// @Description Login
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Request body"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /login [post]
func LoginHandler(ctx *gin.Context) {
	request := LoginRequest{}

	ctx.BindJSON(&request)

	var user model.User
	if err := db.Where("email = ?", request.Email).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	match := helpers.CheckPasswordHash(request.Password, user.Password)
	if !match {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := config.GenerateToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})

}
