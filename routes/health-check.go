package routes

import (
	"github.com/GadSAR/repo-scraper/controllers"
	"github.com/gin-gonic/gin"
)

func HealthCheckRouter(router *gin.Engine) {
	router.GET("/health-check", controllers.GetHealthCheck)
}
