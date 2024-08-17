package router

import (
	docs "github.com/RenatoFa/gopportunities.git/docs"
	"github.com/RenatoFa/gopportunities.git/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	docs.SwaggerInfo.BasePath = "/api/v1"
	route := router.Group("/api/v1")
	{
		route.GET("/opening", handler.ShowOpeningHandler)

		route.POST("/opening", handler.CreateOpeningHandler)

		route.DELETE("/opening", handler.DeleteOpeningHandler)

		route.PUT("/opening", handler.UpdateOpeningHandler)

		route.GET("/openings", handler.ListOpeningsHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
