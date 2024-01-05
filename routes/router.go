package routes

import (
	"os"

	"github.com/Adrephos/api/middlewares"
	routes_v1 "github.com/Adrephos/api/routes/v1"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	environment := os.Getenv("DEBUG")
	if environment == "true" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.CORSMiddleware())

	routes_v1.GameEndpoints(router)

	return router
}
