package main

import (
	"github.com/GadSAR/repo-scraper/utills"
	"go.uber.org/zap"
)

func init() {
	utills.Init()
}

func main() {
	router := utills.SetupRouter()
	config := utills.Config
	zap.L().Info("Starting server", zap.String("host", config.HOST))

	if err := router.Run(config.HOST); err != nil {
		zap.L().Fatal("Failed to start server", zap.Error(err))
	}
}
