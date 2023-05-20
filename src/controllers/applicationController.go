package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ApplicationController struct{}

func (app *ApplicationController) Ping(c *gin.Context) {
	logrus.Info("ApplicationController.Ping")
	c.String(http.StatusOK, "Pong")
}
