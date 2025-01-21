package utills

import (
	"github.com/GadSAR/repo-scraper/middlewares"
	"github.com/GadSAR/repo-scraper/routes"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	initMiddlewares(router)
	initRoutes(router)

	return router
}

func initRoutes(router *gin.Engine) {
	routes.HealthCheckRouter(router)
	routes.RepoRouter(router)
}

func initMiddlewares(router *gin.Engine) {
	router.Use(middlewares.Logger())
	router.Use(gin.Recovery())
}
