package controllers

import (
	"net/http"

	"github.com/SIST-Admission/adm-backend/src/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type BatchesController struct{}

var batchesService service.BatchesService = service.BatchesService{}

func (batchesController *BatchesController) GetBatches(c *gin.Context) {
	logrus.Info("ApplicationsController.RegisterRoutes")

	resp, e := batchesService.GetBatches(c.Keys["userId"].(int))
	if e != nil {
		logrus.Error(e.Message)
		c.JSON(e.Code, e)
		return
	}

	c.JSON(http.StatusCreated, resp)
}
