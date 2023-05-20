package middlewares

import (
	"github.com/SIST-Admission/adm-backend/src/models"
	"github.com/SIST-Admission/adm-backend/src/repositories"
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
		logrus.Error("Auth: JWT Verification Failed: ", err)
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	// Check user if exists
	userRepository := repositories.UserRepository{}
	user, err := userRepository.GetUserById(claims["userId"].(string))
	if err != nil {
		logrus.Error("Auth: User Not Found", err)
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	if user == nil {
		logrus.Error("Auth: User Not Found")
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	// check if user is active
	if !user.IsActive {
		logrus.Error("Auth: User Not Active")
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	// add user to context
	c.Set("user", user)

	c.Next()
}

func AdminAuth(c *gin.Context) {
	logrus.Info("Middleware:AdminAuth")
	user := c.Keys["user"].(*models.User)
	logrus.Info("User Role", user.Role)
	if user.Role != "ADMIN" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Next()
}
