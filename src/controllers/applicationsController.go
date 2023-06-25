package controllers

import (
	"net/http"
	"strconv"

	"github.com/SIST-Admission/adm-backend/src/dto"
	"github.com/SIST-Admission/adm-backend/src/repositories"
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

	// Get UserID from application ID
	userRepo := repositories.UserRepository{}
	user, err := userRepo.GetUserById(strconv.Itoa(userId))

	if err != nil {
		logrus.Error("Failed to Get User by ID", err)
		c.JSON(http.StatusInternalServerError, dto.Error{Code: http.StatusInternalServerError, Message: "Failed to fetch Application"})
		return
	}

	if strconv.Itoa(user.ApplicationId) != c.Param("appId") && c.Keys["role"] != "ADMIN" {
		c.JSON(http.StatusForbidden, dto.Error{Code: http.StatusForbidden, Message: "Forbidden"})
		return
	}

	application, e := applicationsService.GetApplication(c.Param("appId"))
	if e != nil {
		logrus.Error(e.Message)
		c.JSON(e.Code, e)
		return
	}

	c.JSON(http.StatusOK, application)
}

func (applicationsController *ApplicationsController) GetApplicationByUser(c *gin.Context) {
	logrus.Info("ApplicationController.GetApplication")

	userId := c.Keys["userId"].(int)

	// Get UserID from application ID
	userRepo := repositories.UserRepository{}
	user, err := userRepo.GetUserById(strconv.Itoa(userId))

	if err != nil {
		logrus.Error("Failed to Get User by ID", err)
		c.JSON(http.StatusInternalServerError, dto.Error{Code: http.StatusInternalServerError, Message: "Failed to fetch Application"})
		return
	}

	application, e := applicationsService.GetApplication(strconv.Itoa(user.ApplicationId))
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

func (applicationsController *ApplicationsController) GetAppApplications(c *gin.Context) {
	logrus.Info("ApplicationController.GetAppApplications")

	var request dto.GetAllApplicationsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	applications, e := applicationsService.GetAllApplications(&request)
	if e != nil {
		logrus.Error(e.Message)
		c.JSON(e.Code, e)
		return
	}

	c.JSON(http.StatusOK, applications)
}

func (applicationsController *ApplicationsController) UpdateDocumentStatus(c *gin.Context) {
	logrus.Info("ApplicationController.updateDocumentStatus")
	var request dto.UpdateDocumentStatusRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, e := applicationsService.UpdateDocumentStatus(&request)
	if e != nil {
		logrus.Error(e.Message)
		c.JSON(e.Code, e)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (applicationsController *ApplicationsController) UpdateApplicationStatus(c *gin.Context) {
	logrus.Info("ApplicationController.updateApplicationStatus")
	var request dto.UpdateApplicationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, e := applicationsService.UpdateApplicationStatus(&request)
	if e != nil {
		logrus.Error(e.Message)
		c.JSON(e.Code, e)
		return
	}

	c.JSON(http.StatusOK, resp)
}
