package router

import (
	"github.com/dandevweb/gopportunities/config"
	"github.com/dandevweb/gopportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	config.SwaggerInit()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/openings/:id", handler.ShowOpeningHandler)
		v1.POST("/openings", handler.CreateOpeningHandler)
		v1.DELETE("/openings/:id", handler.DeleteOpeningHandler)
		v1.PUT("/openings/:id", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
