package controllers

import (
	"net/http"
	"os"

	"github.com/GadSAR/repo-scraper/models"
	"github.com/GadSAR/repo-scraper/services"
	"github.com/gin-gonic/gin"
	"github.com/go-git/go-billy/v5"
	"go.uber.org/zap"
)

func GetRepoCheck(c *gin.Context) {
	var request models.CheckRepoRequest
	var files []os.FileInfo
	var errMsg string
	var err error

	if err = c.BindJSON(&request); err != nil {
		errMsg = "Error binding JSON"

	} else {
		var fs billy.Filesystem
		zap.L().Info("Repo checkin", zap.String("repo", request.CloneURL), zap.Float64("size(MB)", request.Size))

		if fs, err = services.GetRepoFileSystem(request.CloneURL); err != nil {
			errMsg = "Error getting repo"

		} else {
			if files, err = services.TraverseFiles(fs, nil); err != nil {
				errMsg = "Error traversing files"

			} else {
				files = services.ExceedThreshold(files, int64(request.Size*1000000))
			}
		}
	}

	if err != nil {
		zap.L().Error(errMsg, zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": errMsg, "error": err.Error()})

	} else {
		zap.L().Info("Repo checked", zap.Int("total", len(files)))
		c.JSON(http.StatusOK, gin.H{"total": len(files), "files": services.FormatFiles(files)})
	}
}
