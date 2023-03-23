package public

import (
	"github.com/gin-gonic/gin"
)

func InitRoute(Router *gin.RouterGroup) gin.IRoutes {
	baseRoute := Router.Group("/api/v1")
	{
		baseRoute.GET("health", Health)
	}

	return baseRoute
}
