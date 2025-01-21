package main

import (
	"github.com/GadSAR/repo-scraper/middlewares"
	"github.com/GadSAR/repo-scraper/models"
	"github.com/GadSAR/repo-scraper/routes"
	"github.com/GadSAR/repo-scraper/utills"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var config *models.Config

func init() {
	initLogger()
	loadConfig()
	gin.SetMode(config.GIN_MODE)
}

func initLogger() {
	zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))
}

func loadConfig() {
	if err := utills.LoadConfig(); err != nil {
		zap.L().Fatal("Failed to load config", zap.Error(err))
		panic(err)
	}

	config = utills.Config
}

func main() {
	router := setupRouter()

	zap.L().Info("Starting server", zap.String("host", config.HOST))

	if err := router.Run(config.HOST); err != nil {
		zap.L().Fatal("Failed to start server", zap.Error(err))
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	InitMiddlewares(router)
	InitRoutes(router)

	return router
}

func InitRoutes(router *gin.Engine) {
	routes.HealthCheckRouter(router)
}

func InitMiddlewares(router *gin.Engine) {
	router.Use(middlewares.Logger())
	router.Use(gin.Recovery())
}
