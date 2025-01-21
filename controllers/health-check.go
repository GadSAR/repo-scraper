package controllers

import (
	"net/http"

	"github.com/GadSAR/repo-scraper/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetHealthCheck(c *gin.Context) {
	zap.L().Info("Checking if server is healthy")
	c.String(http.StatusOK, services.HealthMessage())
}
