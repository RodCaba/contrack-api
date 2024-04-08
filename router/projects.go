package router

import (
	"contrack-api/controllers"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	{
		project := router.Group("/projects")
		project.GET("", controllers.FindProjects)
	}

	return router
}