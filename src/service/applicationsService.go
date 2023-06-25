package service

import (
	"strconv"

	"github.com/SIST-Admission/adm-backend/src/dto"
	"github.com/SIST-Admission/adm-backend/src/models"
	"github.com/SIST-Admission/adm-backend/src/repositories"
	"github.com/SIST-Admission/adm-backend/src/validators"
	"github.com/sirupsen/logrus"
)

type ApplicationsService struct{}

var applicationsRepository repositories.ApplicationsRepository = repositories.ApplicationsRepository{}
var submissionsRepository repositories.SubmissionsRepository = repositories.SubmissionsRepository{}
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
		basicDetails, err = applicationsRepository.SaveBasicDetails(userId, application.Id, request)
		if err != nil {
			logrus.Error("Error saving basic details: ", err)
			return nil, &dto.Error{Code: 500, Message: "Internal Server Error"}
		}
	} else {
		logrus.Info("Updating existing basic details")
		basicDetails, err = applicationsRepository.UpdateBasicDetails(userId, application.BasicDetailsId, request)
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

func (applicationsService *ApplicationsService) GetApplication(appId string) (*dto.GetApplicationResponse, *dto.Error) {
	logrus.Info("ApplicationsService.GetApplication")
	logrus.Info("Application: ", appId)

	applicationId, err := strconv.Atoi(appId)
	if err != nil {
		logrus.Error(err)
		return nil, &dto.Error{Code: 400, Message: "Invalid Application Id"}
	}

	// Get Application Details and Basic Details from database
	applicationDetails, err := applicationsRepository.GetApplicationDetails(applicationId)
	if err != nil {
		logrus.Error(err)
		return nil, &dto.Error{Code: 500, Message: err.Error()}
	}

	return &dto.GetApplicationResponse{
		Id:                   applicationDetails.Id,
		ApplicationType:      applicationDetails.ApplicationType,
		Status:               applicationDetails.Status,
		BasicDetails:         applicationDetails.BasicDetails,
		ApplicationStartDate: applicationDetails.ApplicationStartDate,
		AcademicDetails:      applicationDetails.AcademicDetails,
		PaymentDetails:       applicationDetails.PaymentDetails,
		Submissions:          &applicationDetails.Submissions,
	}, nil
}

func (applicationsService *ApplicationsService) SaveAcademicDetails(userId int, request *dto.SaveAcademicDetailsRequest) (map[string]interface{}, *dto.Error) {
	logrus.Info("ApplicationsService.SaveAcademicDetails")
	logrus.Info("User: ", userId)

	application, err := applicationsRepository.GetApplicationByUserId(userId)
	if err != nil {
		logrus.Error(err)
		return nil, &dto.Error{Code: 500, Message: err.Error()}
	}

	if application == nil {
		logrus.Error("Application does not exist for user: ", userId)
		return nil, &dto.Error{Code: 400, Message: "Application does not exist"}
	}

	// Save Academic Details to database
	err = applicationsRepository.SaveAcademicDetails(userId, application.Id, request)
	if err != nil {
		logrus.Error("Error saving academic details: ", err)
		return nil, &dto.Error{Code: 500, Message: "Internal Server Error"}
	}

	return map[string]interface{}{
		"code":    201,
		"success": true,
		"message": "Academic Details saved successfully",
	}, nil
}

func (applicationsService *ApplicationsService) SubmitApplication(userId int, payload *dto.SubmitApplicationRequest) (map[string]interface{}, *dto.Error) {
	logrus.Info("ApplicationsService.SubmitApplication")
	logrus.Info("User: ", userId)

	application, err := applicationsRepository.GetApplicationByUserId(userId)
	if err != nil {
		logrus.Error(err)
		return nil, &dto.Error{Code: 500, Message: err.Error()}
	}

	if application == nil {
		logrus.Error("Application does not exist for user: ", userId)
		return nil, &dto.Error{Code: 400, Message: "Application does not exist"}
	}

	// Save Academic Details to database
	_, err = submissionsRepository.CreateSubmission(userId, application.Id, payload)
	if err != nil {
		logrus.Error("Error creating submission: ", err)
		return nil, &dto.Error{Code: 500, Message: "Internal Server Error"}
	}
	if err != nil {
		logrus.Error("Error submitting application: ", err)
		return nil, &dto.Error{Code: 500, Message: "Internal Server Error"}
	}

	return map[string]interface{}{
		"code":    201,
		"success": true,
		"message": "Application submitted successfully",
	}, nil
}

func (applicationsService *ApplicationsService) GetAllApplications(req *dto.GetAllApplicationsRequest) ([]*models.Application, *dto.Error) {
	logrus.Info("ApplicationsService.GetAllApplications:", req.Status)
	applications, err := applicationsRepository.GetAllApplications(req)
	if err != nil {
		logrus.Error("Error getting all applications: ", err)
		return nil, &dto.Error{Code: 500, Message: err.Error()}
	}
	if applications == nil {
		applications = []*models.Application{}
	}
	return applications, nil
}

func (applicationsService *ApplicationsService) UpdateDocumentStatus(doc *dto.UpdateDocumentStatusRequest) (*map[string]interface{}, *dto.Error) {
	logrus.Info("ApplicationsService.UpdateDocumentStatus")
	err := applicationsRepository.UpdateDocumentStatus(doc)
	if err != nil {
		logrus.Error("Error updating document status: ", err)
		return nil, err
	}

	return &map[string]interface{}{
		"code":       200,
		"success":    true,
		"message":    "Document status updated successfully",
		"documentId": doc.DocumentId,
		"status":     doc.Status,
		"isVerified": doc.IsVerified,
	}, nil
}

func (applicationsService *ApplicationsService) UpdateApplicationStatus(req *dto.UpdateApplicationRequest) (*map[string]interface{}, *dto.Error) {
	logrus.Info("ApplicationsService.UpdateApplicationStatus")
	err := applicationsRepository.UpdateApplicationStatus(req)
	if err != nil {
		logrus.Error("Error updating application status: ", err)
		return nil, err
	}

	return &map[string]interface{}{
		"code":    200,
		"success": true,
		"message": "Application status updated successfully",
		"status":  req.Status,
		"id":      req.Id,
	}, nil
}
