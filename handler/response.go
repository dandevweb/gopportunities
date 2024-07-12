package handler

import (
	"fmt"
	"net/http"

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
