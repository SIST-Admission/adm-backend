package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Auth(c *gin.Context) {
	logrus.Info("Middleware:Auth")
	cookie, err := c.Cookie("auth")
	if err != nil {
		logrus.Error("Auth: ", err)
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	log.Default().Println(cookie)
	logrus.Info("Auth: ", cookie)
	c.Next()
}
