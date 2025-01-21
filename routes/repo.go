package routes

import (
	"github.com/GadSAR/repo-scraper/controllers"
	"github.com/gin-gonic/gin"
)

func RepoRouter(router *gin.Engine) {
	router.GET("/repo/check", controllers.GetRepoCheck)
}
