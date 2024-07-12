package router

import (
	"github.com/dandevweb/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/openings/:id", handler.ShowOpeningHandler)
		v1.POST("/openings", handler.CreateOpeningHandler)
		v1.DELETE("/openings", handler.DeleteOpeningHandler)
		v1.PUT("/openings", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}

}
