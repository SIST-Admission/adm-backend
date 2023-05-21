package controllers

import (
	"net/http"
	"time"

	"github.com/SIST-Admission/adm-backend/src/dto"
	"github.com/SIST-Admission/adm-backend/src/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type UserController struct{}

var userService service.UserService = service.UserService{}

func (userController *UserController) RegisterUser(c *gin.Context) {
	logrus.Info("UserController.RegisterUser")
	var request dto.RegisterUserRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, e := userService.RegisterUser(request)
	if e != nil {
		logrus.Error(e.Message)
		c.JSON(e.Code, e)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (userController *UserController) LoginUser(c *gin.Context) {
	logrus.Info("UserController.RegisterUser")
	var request dto.LoginUserRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, e := userService.LoginUser(request)
	if e != nil {
		logrus.Error(e.Message)
		c.JSON(e.Code, e)
		return
	}
	exp := int(time.Second) * 60 * 60 * 24
	// set HTTPOnly Secure Cookie
	cookieHost := viper.GetString(viper.GetString("env") + "." + "server.host")
	c.SetCookie("auth", resp.JwtToken, exp, "/", cookieHost, false, false)

	c.JSON(http.StatusOK, resp)
}

func (userController *UserController) LogoutUser(c *gin.Context) {
	logrus.Info("UserController.LogoutUser")
	cookieHost := viper.GetString(viper.GetString("env") + "." + "server.host")
	c.SetCookie("auth", "", -1, "/", cookieHost, false, false)
	c.JSON(http.StatusOK, gin.H{"message": "Logout Successful"})
}

func (userController *UserController) LoggedInUser(c *gin.Context) {

	resp, e := userService.GetUser(c.Keys["userId"].(int))
	if e != nil {
		logrus.Error(e.Message)
		c.JSON(e.Code, e)
		return
	}

	c.JSON(http.StatusCreated, resp)
}
