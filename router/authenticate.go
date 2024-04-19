package router
import (
	"contrack-api/controllers"
	"github.com/gin-gonic/gin"
)

func InitAuthenticateRouter(router *gin.Engine) {
	{
		authenticate := router.Group("/authenticate")
		authenticate.POST("", controllers.Authenticate)
	}
}