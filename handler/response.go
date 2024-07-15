package handler

import (
	"fmt"
	"net/http"

	"github.com/dandevweb/gopportunities/model"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	headerJson(ctx)
	ctx.JSON(code, gin.H{"message": msg, "errorCode": code})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	headerJson(ctx)
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfully", op),
		"data":    data,
	})
}

func sendNoContent(ctx *gin.Context) {
	headerJson(ctx)
	ctx.JSON(http.StatusNoContent, nil)
}

func headerJson(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}

type CreateOpeningResponse struct {
	Message string                `json:"message"`
	Data    model.OpeningResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string `json:"message"`
}

type ShowOpeningResponse struct {
	Message string                `json:"message"`
	Data    model.OpeningResponse `json:"data"`
}

type ListOpeningsResponse struct {
	Message string                  `json:"message"`
	Data    []model.OpeningResponse `json:"data"`
}

type UpdateOpeningResponse struct {
	Data    model.OpeningResponse `json:"data"`
	Message string                `json:"message"`
}

type RegisterResponse struct {
	Message string             `json:"message"`
	Data    model.UserResponse `json:"data"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
