package dto

import (
	"github.com/SIST-Admission/adm-backend/src/models"
)

type StartApplicationRequst struct {
	ApplicationType string `json:"applicationType"`
}

type StartApplicationResponse struct {
	Id                   int    `json:"id"`
	BasicDetailsId       int    `json:"basicDetailsId"`
	AcademicDetailsId    int    `json:"academicDetailsId"`
	PaymetId             int    `json:"paymetId"`
	ApplicationType      string `json:"applicationType"`
	ApplicationStartDate string `json:"applicationStartDate"`
	Status               string `json:"status"`
}

type GetApplicationResponse struct {
	Id                   int                  `json:"id"`
	ApplicationType      string               `json:"applicationType"`
	Status               string               `json:"status"`
	BasicDetails         *models.BasicDetails `json:"basicDetails"`
	ApplicationStartDate string               `json:"applicationStartDate"`
}
