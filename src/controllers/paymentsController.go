package controllers

import (
	"bytes"
	"net/http"

	"github.com/SIST-Admission/adm-backend/src/dto"
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
func (paymentsController *PaymentsController) GetAdmissionOrder(c *gin.Context) {
	logrus.Info("PaymentsController.GetAdmissionOrder")
	var request dto.GetAdmissionOrderRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, e := paymentsService.GetAdmissionOrder(&request)
	if e != nil {
		logrus.Error(e)
		c.JSON(e.Code, e)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (paymentsController *PaymentsController) VerifyPayment(c *gin.Context) {
	logrus.Info("UserController.RegisterUser")
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	payload := buf.String()
	signature := c.Request.Header.Get("x-razorpay-signature")
	logrus.Info("Payload: ", payload)
	logrus.Info("Signature: ", signature)
	go paymentsService.VerifyPayment(payload, signature)
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}

func (paymentsController *PaymentsController) GetTransactions(c *gin.Context) {
	logrus.Info("PaymentsController.GetPaymentDetails")

	resp, e := paymentsService.GetTransactions(c.Keys["userId"].(int))
	if e != nil {
		logrus.Error(e.Message)
		c.JSON(e.Code, e)
		return
	}

	c.JSON(http.StatusOK, resp)
}
