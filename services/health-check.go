package services

import (
	"go.uber.org/zap"
)

func HealthMessage() string {
	zap.L().Info("Checking if server is healthy")
	return "The server is healthy"
}
