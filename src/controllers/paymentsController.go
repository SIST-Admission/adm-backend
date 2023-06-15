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

func (paymentsController *PaymentsController) VerifyPayment(c *gin.Context) {
	logrus.Info("UserController.RegisterUser")
	var request map[string]interface{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Info(request)
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}
