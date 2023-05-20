package middlewares

import (
	"github.com/SIST-Admission/adm-backend/src/utils"
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

	claims, err := utils.ParseJwt(cookie)
	if err != nil {
		logrus.Error("JWT Verification Failed: Auth: ", err)
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	logrus.Debug("claims: ", claims)
	c.Next()
}
