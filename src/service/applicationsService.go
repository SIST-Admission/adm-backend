package service

import (
	"github.com/SIST-Admission/adm-backend/src/dto"
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
		return nil, &dto.Error{Code: 500, Message: "Internal Server Error"}
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

func (applicationsService *ApplicationsService) SaveBasicDetails(userId int, request dto.SaveBasicDetailsRequest) (dto.SaveBasicDetailsResponse, *dto.Error) {
	logrus.Info("ApplicationsService.SaveBasicDetails")

	logrus.Info("User: ", userId)

	// TODO: check if application exists for user

	// Validate request
	fieldErrors := applicationValidator.ValidateSaveBasicDetailsRequest(&request)
	if len(fieldErrors) > 0 {
		return dto.SaveBasicDetailsResponse{}, &dto.Error{Code: 400, Message: fieldErrors}
	}

	// TODO: Save basic details to database

	// Return response
	return dto.SaveBasicDetailsResponse{}, nil
}
