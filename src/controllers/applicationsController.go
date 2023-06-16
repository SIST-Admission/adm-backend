package controllers

import (
	"net/http"
	"strconv"

	"github.com/SIST-Admission/adm-backend/src/dto"
	"github.com/SIST-Admission/adm-backend/src/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ApplicationsController struct{}

var applicationsService service.ApplicationsService = service.ApplicationsService{}

func (applicationsController *ApplicationsController) StartApplication(c *gin.Context) {
	logrus.Info("ApplicationsController.RegisterRoutes")
	var request dto.StartApplicationRequst
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, e := applicationsService.StartApplication(c.Keys["userId"].(int), &request)
	if e != nil {
		logrus.Error(e.Message)
		c.JSON(e.Code, e)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (applicationsController *ApplicationsController) SaveBasicDetails(c *gin.Context) {
	logrus.Info("ApplicationController.SaveBasicDetails")
	var request dto.SaveBasicDetailsRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, e := applicationsService.SaveBasicDetails(c.Keys["userId"].(int), &request)
	if e != nil {
		logrus.Error(e.Message)
		c.JSON(e.Code, e)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (applicationsController *ApplicationsController) GetApplication(c *gin.Context) {
	logrus.Info("ApplicationController.GetApplication")

	userId := c.Keys["userId"].(int)

	if strconv.Itoa(userId) != c.Param("userId") && c.Keys["role"] != "ADMIN" {
		c.JSON(http.StatusForbidden, dto.Error{Code: http.StatusForbidden, Message: "Forbidden"})
		return
	}

	application, e := applicationsService.GetApplication(userId)
	if e != nil {
		logrus.Error(e.Message)
		c.JSON(e.Code, e)
		return
	}

	c.JSON(http.StatusOK, application)
}

func (applicationsController *ApplicationsController) SaveAcademicDetails(c *gin.Context) {
	logrus.Info("ApplicationController.SaveAcademicDetails")
	var request dto.SaveAcademicDetailsRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, e := applicationsService.SaveAcademicDetails(c.Keys["userId"].(int), &request)
	if e != nil {
		logrus.Error(e.Message)
		c.JSON(e.Code, e)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (applicationsController *ApplicationsController) SubmitApplication(c *gin.Context) {
	logrus.Info("ApplicationController.SubmitApplication")
	var request dto.SubmitApplicationRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, e := applicationsService.SubmitApplication(c.Keys["userId"].(int), &request)
	if e != nil {
		logrus.Error(e.Message)
		c.JSON(e.Code, e)
		return
	}

	c.JSON(http.StatusCreated, resp)
}
