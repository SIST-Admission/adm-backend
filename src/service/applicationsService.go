package service

import (
	"github.com/SIST-Admission/adm-backend/src/dto"
	"github.com/SIST-Admission/adm-backend/src/models"
	"github.com/SIST-Admission/adm-backend/src/repositories"
	"github.com/SIST-Admission/adm-backend/src/validators"
	"github.com/sirupsen/logrus"
)

type ApplicationsService struct{}

var applicationsService ApplicationsService = ApplicationsService{}
var applicationsRepository repositories.ApplicationsRepository = repositories.ApplicationsRepository{}
var applicationValidator validators.ApplicationValidator = validators.ApplicationValidator{}

func (applicationsService *ApplicationsService) StartApplication(userId int, request *dto.StartApplicationRequst) (*dto.StartApplicationResponse, *dto.Error) {
	logrus.Info("ApplicationsService.StartApplication")
	logrus.Info("ApplicationsService.StartApplication: User: ", userId)
	application, err := applicationsRepository.GetApplicationByUserId(userId)
	if err != nil {
		logrus.Error(err)
		return nil, &dto.Error{Code: 500, Message: err.Error()}
	}

	if application != nil {
		logrus.Error("Application already exists for user: ", userId)
		return nil, &dto.Error{Code: 400, Message: "Application already exists"}
	}

	// Create New Application
	newApplication, err := applicationsRepository.CreateNewApplication(userId, request.ApplicationType)
	if err != nil {
		logrus.Error("Error creating new application: ", err)
		return nil, &dto.Error{Code: 500, Message: "Internal Server Error"}
	}

	return &dto.StartApplicationResponse{
		Id:                   newApplication.Id,
		ApplicationType:      newApplication.ApplicationType,
		ApplicationStartDate: newApplication.ApplicationStartDate,
		Status:               newApplication.Status,
	}, nil
}

func (applicationsService *ApplicationsService) SaveBasicDetails(userId int, request *dto.SaveBasicDetailsRequest) (*dto.SaveBasicDetailsResponse, *dto.Error) {
	logrus.Info("ApplicationsService.SaveBasicDetails")
	logrus.Info("User: ", userId)

	// Validate request
	fieldErrors := applicationValidator.ValidateSaveBasicDetailsRequest(request)
	if len(fieldErrors) > 0 {
		return nil, &dto.Error{Code: 400, Message: fieldErrors}
	}

	// TODO: Save basic details to database
	application, err := applicationsRepository.GetApplicationByUserId(userId)
	if err != nil {
		logrus.Error(err)
		return nil, &dto.Error{Code: 500, Message: err.Error()}
	}

	if application == nil {
		logrus.Error("Application does not exist for user: ", userId)
		return nil, &dto.Error{Code: 400, Message: "Application does not exist"}
	}

	var basicDetails *models.BasicDetails
	// TODO: Save basic details to database
	if application.BasicDetailsId == 0 {
		logrus.Info("Creating new basic details")
		basicDetails, err = applicationsRepository.SaveBasicDetails(application.Id, request)
		if err != nil {
			logrus.Error("Error saving basic details: ", err)
			return nil, &dto.Error{Code: 500, Message: "Internal Server Error"}
		}
	} else {
		logrus.Info("Updating existing basic details")
		basicDetails, err = applicationsRepository.UpdateBasicDetails(application.BasicDetailsId, request)
		if err != nil {
			logrus.Error("Error saving basic details: ", err)
			return nil, &dto.Error{Code: 500, Message: "Internal Server Error"}
		}
	}

	return &dto.SaveBasicDetailsResponse{
		Id:                 basicDetails.Id,
		Name:               basicDetails.Name,
		DoB:                basicDetails.DoB,
		Gender:             basicDetails.Gender,
		Category:           basicDetails.Category,
		IsCoI:              basicDetails.IsCoI,
		IsPwD:              basicDetails.IsPwD,
		FatherName:         basicDetails.FatherName,
		MotherName:         basicDetails.MotherName,
		Nationality:        basicDetails.Nationality,
		IdentityType:       basicDetails.IdentityType,
		IdentityNumber:     basicDetails.IdentityNumber,
		IdentityDocumentId: basicDetails.IdentityDocumentId,
	}, nil
}
