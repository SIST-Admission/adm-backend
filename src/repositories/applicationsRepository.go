package repositories

import (
	"github.com/SIST-Admission/adm-backend/src/dto"
	"github.com/SIST-Admission/adm-backend/src/models"
	"github.com/sirupsen/logrus"
)

type ApplicationsRepository struct{}

func (repo *ApplicationsRepository) SaveBasicDetails(payload dto.SaveBasicDetailsRequest) (*models.BasicDetails, error) {
	logrus.Info("ApplicationsRepository.SaveBasicDetails")
	// db := db.GetInstance()
	return nil, nil
}
