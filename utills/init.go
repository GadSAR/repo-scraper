package utills

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Init() {
	initLogger()
	initConfig()
	gin.SetMode(Config.GIN_MODE)
}

func initLogger() {
	if logger, err := zap.NewDevelopment(); err != nil {
		panic(err)
	} else {
		zap.ReplaceGlobals(logger)
	}
}

func initConfig() {
	if err := LoadConfig(); err != nil {
		zap.L().Fatal("Failed to load config", zap.Error(err))
		panic(err)
	}
}
