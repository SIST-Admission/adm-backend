package controllers

import (
	"net/http"

	"github.com/SIST-Admission/adm-backend/src/dto"
	"github.com/SIST-Admission/adm-backend/src/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MeritListsController struct{}

var meritListService service.MeritListService = service.MeritListService{}

func (meritListsController *MeritListsController) CreateMeritList(c *gin.Context) {
	logrus.Info("MeritListsController.CreateMeritList")
	var request dto.CreateMeritListRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, e := meritListService.CreateMeritList(&request)
	if e != nil {
		logrus.Error(e.Message)
		c.JSON(e.Code, e)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (meritListsController *MeritListsController) AddStudents(c *gin.Context) {
	logrus.Info("MeritListsController.AddStudents")
	var request dto.AddStudentsToMeritListRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, e := meritListService.AddStudents(&request)
	if e != nil {
		logrus.Error(e.Message)
		c.JSON(e.Code, e)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (MeritListsController *MeritListsController) GetAllMeritLists(c *gin.Context) {
	logrus.Info("MeritListsController.GetAllMeritLists")
	var request dto.GetAllMeritListsRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, e := meritListService.GetAllMeritLists(&request)
	if e != nil {
		logrus.Error(e.Message)
		c.JSON(e.Code, e)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (MeritListsController *MeritListsController) GetUnListedCandidates(c *gin.Context) {
	logrus.Info("MeritListsController.GetAllMeritLists")
	var request dto.GetUnListedCandidatesRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, e := meritListService.GetUnListedCandidates(&request)
	if e != nil {
		logrus.Error(e.Message)
		c.JSON(e.Code, e)
		return
	}

	c.JSON(http.StatusOK, resp)
}
