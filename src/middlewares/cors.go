package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func CORSMiddleware() gin.HandlerFunc {
	logrus.Info("Middleware:CORSMiddleware")
	return func(c *gin.Context) {

		allowOrigin := viper.GetString(viper.GetString("env" + "." + "cors.allowOrigin"))
		logrus.Info("Middleware:CORSMiddleware: allowOrigin: ", allowOrigin)

		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
