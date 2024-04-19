package router

import (
	"contrack-api/controllers"
	"contrack-api/middleware"

	"github.com/gin-gonic/gin"
)

func InitProjectRouter(router *gin.Engine) {
	{
		project := router.Group("/projects")
		project.GET("", middleware.RequireAuth, controllers.FindProject)
		project.POST("", controllers.CreateProject)
	}
}