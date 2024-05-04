package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/AndrewFilkin/for-any-test/pkg/services"
)

func InitRoutes() *gin.Engine {

	router := gin.New()
	api := router.Group("/api")
	{
		lists := api.Group("/hello")
		{
			lists.GET("/", services.HelloWorld)
			// lists.GET("/:id")
			// lists.PUT("/:id")
			// lists.DELETE("/:id")
		}
	}
	return router

}
