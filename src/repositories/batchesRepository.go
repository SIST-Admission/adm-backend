package repositories

import (
	"github.com/SIST-Admission/adm-backend/src/db"
	"github.com/SIST-Admission/adm-backend/src/dto"
	"github.com/SIST-Admission/adm-backend/src/models"
	"github.com/sirupsen/logrus"
)

type BatchesRepository struct{}

func (repo *BatchesRepository) GetBatches() ([]models.Batch, *dto.Error) {
	logrus.Info("BatchesRepository.GetBatches")

	db := db.GetInstance()

	var batches []models.Batch
	if err := db.Preload("Department").Find(&batches).Error; err != nil {
		logrus.Error("Failed to get batches: ", err)
		return nil, &dto.Error{
			Code:    500,
			Message: "Failed to get batches",
		}
	}

	return batches, nil
}
