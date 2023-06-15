package controllers

import (
	"net/http"

	"github.com/SIST-Admission/adm-backend/src/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type PaymentsController struct{}

var paymentsService service.PaymentsService = service.PaymentsService{}

func (paymentsController *PaymentsController) GetOrder(c *gin.Context) {
	logrus.Info("PaymentsController.GetOrder")

	resp, e := paymentsService.GetOrder(c.Keys["userId"].(int))
	if e != nil {
		logrus.Error(e.Message)
		c.JSON(e.Code, e)
		return
	}

	c.JSON(http.StatusCreated, resp)
}
