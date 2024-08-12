package router

import (
	"github.com/RenatoFa/gopportunities.git/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	route := router.Group("/api/v1")
	{
		route.GET("/opening", handler.ShowOpeningHandler)

		route.POST("/opening", handler.CreateOpeningHandler)

		route.DELETE("/opening", handler.DeleteOpeningHandler)

		route.PUT("/opening", handler.UpdateOpeningHandler)

		route.GET("/openings", handler.ListOpeningsHandler)
	}
}
