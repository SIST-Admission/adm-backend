package validators

import (
	"github.com/SIST-Admission/adm-backend/src/dto"
)

type ApplicationValidator struct{}

func (applicationValidator ApplicationValidator) ValidateSaveBasicDetailsRequest(request *dto.SaveBasicDetailsRequest) []string {
	var fieldErrors []string = make([]string, 0)
	// Validate request
	if request.Name == "" {
		fieldErrors = append(fieldErrors, "name should not be empty")
	}

	if request.DoB == "" {
		fieldErrors = append(fieldErrors, "date of birth should not be empty")
	}

	if request.Gender == "" {
		fieldErrors = append(fieldErrors, "please select a gender")
	}

	if request.Category == "" {
		fieldErrors = append(fieldErrors, "please select a category")
	} else if request.Category != "GEN" && request.Category != "OBC" && request.Category != "SC" && request.Category != "ST" && request.Category != "EWS" {
		fieldErrors = append(fieldErrors, "invalid category")
	}

	if request.FatherName == "" {
		fieldErrors = append(fieldErrors, "father's name should not be empty")
	}

	if request.MotherName == "" {
		fieldErrors = append(fieldErrors, "mother's name should not be empty")
	}

	if request.IdentityType == "" {
		fieldErrors = append(fieldErrors, "please select an identity type")
	} else if request.IdentityType != "Aadhaar" && request.IdentityType != "PAN" && request.IdentityType != "Voter ID" && request.IdentityType != "Passport" {
		fieldErrors = append(fieldErrors, "invalid identity type")
	} else if request.IdentityNumber == "" {
		fieldErrors = append(fieldErrors, "identity number should not be empty")
	}

	if request.Nationality == "" {
		fieldErrors = append(fieldErrors, "please select your nationality")
	}

	return fieldErrors
}
