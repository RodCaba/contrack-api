package router

import (
	"contrack-api/controllers"
	"github.com/gin-gonic/gin"
)

func InitProjectRouter(router *gin.Engine) {
	{
		project := router.Group("/projects")
		project.GET("", controllers.FindProjects)
		project.GET("/:id", controllers.FindProject)
		project.POST("", controllers.CreateProject)
	}
}