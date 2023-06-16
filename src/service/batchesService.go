package service

import (
	"github.com/SIST-Admission/adm-backend/src/dto"
	"github.com/SIST-Admission/adm-backend/src/models"
	"github.com/SIST-Admission/adm-backend/src/repositories"
	"github.com/sirupsen/logrus"
)

type BatchesService struct{}

var batchesRepository repositories.BatchesRepository = repositories.BatchesRepository{}

func (batchesService *BatchesService) GetBatches() ([]models.Batch, *dto.Error) {
	logrus.Info("BatchesService.GetBatches")

	batches, err := batchesRepository.GetBatches()
	if err != nil {
		logrus.Error("Failed to get batches: ", err)
		return nil, &dto.Error{
			Code:    500,
			Message: "Failed to get batches",
		}
	}

	return batches, nil
}
