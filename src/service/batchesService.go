package service

import (
	"time"

	"github.com/SIST-Admission/adm-backend/src/dto"
	"github.com/SIST-Admission/adm-backend/src/models"
	"github.com/SIST-Admission/adm-backend/src/repositories"
	"github.com/sirupsen/logrus"
)

type BatchesService struct{}

var batchesRepository repositories.BatchesRepository = repositories.BatchesRepository{}

func (batchesService *BatchesService) GetBatches(userId int) ([]models.Batch, *dto.Error) {
	logrus.Info("BatchesService.GetBatches")

	application, err := applicationsRepository.GetApplicationByUserId(userId)
	if err != nil {
		logrus.Error("Failed to get application: ", err)
		return nil, &dto.Error{
			Code:    500,
			Message: "Failed to get application",
		}
	}

	var year int

	if application.ApplicationType == "LATERAL" {
		year = time.Now().Year() - 1
	} else {
		year = time.Now().Year()
	}

	batches, err := batchesRepository.GetBatches(year)

	if err != nil {
		logrus.Error("Failed to get batches: ", err)
		return nil, &dto.Error{
			Code:    500,
			Message: "Failed to get batches",
		}
	}

	return batches, nil
}
