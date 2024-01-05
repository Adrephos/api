package routes_v1

import (
	handlers_v1 "github.com/Adrephos/api/handlers/v1"
	"github.com/gin-gonic/gin"
)

func GameEndpoints(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/games/cover/:query", handlers_v1.GetCover)
		v1.GET("/games/thumbnail/:query", handlers_v1.GetThumbnail)
	}
	router.Run()
}
