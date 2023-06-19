package repositories

import (
	"github.com/SIST-Admission/adm-backend/src/db"
	"github.com/SIST-Admission/adm-backend/src/models"
	"github.com/sirupsen/logrus"
)

type BatchesRepository struct{}

func (repo *BatchesRepository) GetBatches(year int) ([]models.Batch, error) {
	logrus.Info("BatchesRepository.GetBatches")

	db := db.GetInstance()

	var batches []models.Batch
	if err := db.Model(models.Batch{}).Preload("Department").Where("start_year = ?", year).Find(&batches).Error; err != nil {
		logrus.Error("Failed to get batches: ", err)
		return nil, err
	}

	return batches, nil
}
