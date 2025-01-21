package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := zap.L()

		logger.Info("Request received", zap.String("method", c.Request.Method), zap.String("path", c.Request.URL.Path))
		c.Next()
		logger.Info("Request completed", zap.Int("status", c.Writer.Status()))
	}
}
