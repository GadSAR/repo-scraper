package controllers

import (
	"net/http"

	"github.com/GadSAR/repo-scraper/models"
	"github.com/GadSAR/repo-scraper/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetRepoCheck(c *gin.Context) {
	var request models.CheckRepoRequest

	if err := c.BindJSON(&request); err != nil {
		zap.L().Error("Error binding JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	zap.L().Info("Repo checkin", zap.String("repo", request.CloneURL), zap.Int("size", request.Size))

	var striped []models.File
	if fs, err := services.GetRepoFileSystem(request.CloneURL); err != nil {
		services.RecursivePrintAllFiles(fs)
		// 	zap.L().Error("Error getting repo files", zap.Error(err))
		// } else {
		// 	striped = services.StripFiles(files)
	}

	c.JSON(http.StatusOK, gin.H{"files": striped})
}
