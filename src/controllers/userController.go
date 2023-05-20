package controllers

import (
	"net/http"

	"github.com/SIST-Admission/adm-backend/src/dto"
	"github.com/SIST-Admission/adm-backend/src/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
