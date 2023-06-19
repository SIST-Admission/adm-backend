package repositories

import (
	"github.com/SIST-Admission/adm-backend/src/db"
	"github.com/SIST-Admission/adm-backend/src/dto"
	"github.com/SIST-Admission/adm-backend/src/models"
	"github.com/sirupsen/logrus"
)

type SubmissionsRepository struct{}

func (repo *SubmissionsRepository) CreateSubmission(userId, appId int, payload *dto.SubmitApplicationRequest) (*models.Submission, error) {
	logrus.Info("SubmissionsRepository.CreateSubmission")
	db := db.GetInstance()

	for _, p := range payload.Submissions {
		s := models.Submission{
			UserId:         userId,
			ApplicationId:  appId,
			DepartmentCode: p.DepartmentCode,
			BatchCode:      p.BatchCode,
			Status:         "created",
			IsVerified:     false,
			IsAdmitted:     false,
		}
		if err := db.Create(&s).Error; err != nil {
			logrus.Error("Failed to create submission: ", err)
			return nil, err
		}
	}

	if err := db.Model(models.Application{}).Where("id = ?", appId).Update("status", "SUBMITTED").Error; err != nil {
		logrus.Error("Failed to update application status: ", err)
		return nil, err
	}

	return nil, nil
}
