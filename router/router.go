package router

import (
	"contrack-api/middleware"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	InitAuthenticateRouter(router)
	InitProjectRouter(router)

	return router
}