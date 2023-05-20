package service

import (
	"github.com/SIST-Admission/adm-backend/src/dto"
	"github.com/SIST-Admission/adm-backend/src/validators"
	"github.com/sirupsen/logrus"
)

type ApplicationsService struct{}

var applicationsService ApplicationsService = ApplicationsService{}
var applicationValidator validators.ApplicationValidator = validators.ApplicationValidator{}

func (applicationsService *ApplicationsService) SaveBasicDetails(request dto.SaveBasicDetailsRequest) (dto.SaveBasicDetailsResponse, *dto.Error) {
	logrus.Info("ApplicationsService.SaveBasicDetails")

	// Validate request
	fieldErrors := applicationValidator.ValidateSaveBasicDetailsRequest(&request)
	if len(fieldErrors) > 0 {
		return dto.SaveBasicDetailsResponse{}, &dto.Error{Code: 400, Message: fieldErrors}
	}

	// TODO: Save basic details to database

	// Return response
	return dto.SaveBasicDetailsResponse{}, nil
}
