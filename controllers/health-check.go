package controllers

import (
	"net/http"

	"github.com/GadSAR/repo-scraper/services"
	"github.com/gin-gonic/gin"
)

func GetHealthCheck(c *gin.Context) {
	c.String(http.StatusOK, services.HealthMessage())
}
