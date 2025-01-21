package utills

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Init() {
	initConfig()
	initLogger()
	initGin()
}

func initConfig() {
	if err := LoadConfig(); err != nil {
		log.Fatal("Failed to load config", zap.Error(err))
		panic(err)
	}
}

func initGin() {
	if Config.GIN_MODE == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}

func initLogger() {
	var logger *zap.Logger
	var err error

	if Config.ZAP_MODE == "production" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		panic(err)
	} else {
		zap.ReplaceGlobals(logger)
	}
}
