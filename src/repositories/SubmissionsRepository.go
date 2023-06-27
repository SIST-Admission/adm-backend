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

func (repo *SubmissionsRepository) GetPaymentBySubmissionId(id int) (*models.Submission, *dto.Error) {
	logrus.Info("SubmissionsRepository.GetPaymentBySubmissionId")
	db := db.GetInstance()
	var submission *models.Submission
	if err := db.Model(models.Submission{}).Preload("Payment").Preload("BatchDetails").Preload("DepartmentDetails").Preload("Application").Preload("Application.BasicDetails").Preload("Application.BasicDetails.PhotoDocument").Where("id = ?", id).First(&submission).Error; err != nil {
		logrus.Error("Failed to get payment id: ", err)
		return nil, &dto.Error{
			Code:    500,
			Message: "Failed to get payment id",
		}
	}

	if submission == nil {
		logrus.Error("Submission not found")
		return nil, &dto.Error{
			Code:    404,
			Message: "Submission not found",
		}
	}

	if submission.PaymentId == nil {
		logrus.Error("Payment not found")
		return nil, nil
	}

	return submission, nil
}

func (repo *SubmissionsRepository) UpdateSubmissionStatus(submissionId int, status string) error {
	logrus.Info("SubmissionsRepository.UpdateSubmissionStatus")
	db := db.GetInstance()
	if status == "captured" {
		if err := db.Model(models.Submission{}).Where("id = ?", submissionId).Update("is_admitted", true).Error; err != nil {
			logrus.Error("Failed to update submission status: ", err)
			return err
		}
	} else {
		logrus.Info("SubmissionsRepository.UpdateSubmissionStatus:", "status is not captured")
	}

	return nil
}
