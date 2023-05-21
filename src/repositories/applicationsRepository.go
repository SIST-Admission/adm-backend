package repositories

import (
	"time"

	"github.com/SIST-Admission/adm-backend/src/db"
	"github.com/SIST-Admission/adm-backend/src/dto"
	"github.com/SIST-Admission/adm-backend/src/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ApplicationsRepository struct{}

func (repo *ApplicationsRepository) GetApplicationByUserId(userId int) (*models.Application, error) {
	logrus.Info("ApplicationsRepository.GetApplicationByUserId")
	db := db.GetInstance()
	var application models.Application
	var user models.User
	if err := db.Model(models.User{}).Where("id = ?", userId).First(&user).Error; err != nil {
		logrus.Error("ApplicationsRepository.GetApplicationByUserId: ", err)
		return nil, err
	}
	logrus.Info("ApplicationsRepository.GetApplicationByUserId: Reterived Application ID ", user.ApplicationId)
	if user.ApplicationId == 0 {
		return nil, nil
	}

	if err := db.Model(models.Application{}).Where("id = ?", user.ApplicationId).First(&application).Error; err != nil {
		logrus.Error("ApplicationsRepository.GetApplicationByUserId: ", err)
		return nil, err
	}
	return &application, nil
}

func (repo *ApplicationsRepository) CreateNewApplication(userId int, applicationType string) (*models.Application, error) {
	logrus.Info("ApplicationsRepository.CreateNewApplication")
	db := db.GetInstance()
	var application models.Application
	if err := db.Transaction(func(tx *gorm.DB) error {
		application = models.Application{
			ApplicationType:      applicationType,
			ApplicationStartDate: time.Now().Format("2006-01-02 15:04:05"),
			Status:               "DRAFT",
		}
		if err := db.Create(&application).Error; err != nil {
			logrus.Error("ApplicationsRepository.CreateNewApplication: ", err)
			return err
		}

		if err := db.Model(models.User{}).Where("id = ?", userId).Update("application_id", application.Id).Error; err != nil {
			logrus.Error("ApplicationsRepository.CreateNewApplication: ", err)
			return err
		}

		return nil
	}); err != nil {
		logrus.Error("ApplicationsRepository.CreateNewApplication: ", err)
		return nil, err
	}
	return &application, nil
}

func (repo *ApplicationsRepository) SaveBasicDetails(appId int, payload *dto.SaveBasicDetailsRequest) (*models.BasicDetails, error) {
	logrus.Info("ApplicationsRepository.SaveBasicDetails")
	db := db.GetInstance()
	tx := db.Begin()
	defer tx.Rollback()

	basicDetails := models.BasicDetails{
		Name:               payload.Name,
		DoB:                time.Now().Format("2006-01-02 15:04:05"),
		Gender:             payload.Gender,
		Category:           payload.Category,
		IsCoI:              payload.IsCoI,
		IsPwD:              payload.IsPwD,
		FatherName:         payload.FatherName,
		MotherName:         payload.MotherName,
		Nationality:        payload.Nationality,
		IdentityType:       payload.IdentityType,
		IdentityNumber:     payload.IdentityNumber,
		IdentityDocumentId: payload.IdentityDocumentId,
	}

	if err := tx.Model(models.BasicDetails{}).Save(&basicDetails).Error; err != nil {
		logrus.Error("ApplicationsRepository.SaveBasicDetails: ", err)
		return nil, err
	}

	if err := tx.Model(models.Application{}).Where("id = ?", appId).Update("basic_details_id", basicDetails.Id).Error; err != nil {
		logrus.Error("ApplicationsRepository.SaveBasicDetails: ", err)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		logrus.Error("ApplicationsRepository.SaveBasicDetails: ", err)
		return nil, err
	}
	return &basicDetails, nil
}

func (repo *ApplicationsRepository) UpdateBasicDetails(basicDetailsId int, payload *dto.SaveBasicDetailsRequest) (*models.BasicDetails, error) {
	logrus.Info("ApplicationsRepository.UpdateBasicDetails")
	db := db.GetInstance()

	var basicDetails models.BasicDetails = models.BasicDetails{
		Id:                 basicDetailsId,
		Name:               payload.Name,
		DoB:                time.Now().Format("2006-01-02 15:04:05"),
		Gender:             payload.Gender,
		Category:           payload.Category,
		IsCoI:              payload.IsCoI,
		IsPwD:              payload.IsPwD,
		FatherName:         payload.FatherName,
		MotherName:         payload.MotherName,
		Nationality:        payload.Nationality,
		IdentityType:       payload.IdentityType,
		IdentityNumber:     payload.IdentityNumber,
		IdentityDocumentId: payload.IdentityDocumentId,
	}

	if err := db.Model(models.BasicDetails{}).Where("id = ?", basicDetailsId).Updates(basicDetails).Error; err != nil {
		logrus.Error("ApplicationsRepository.UpdateBasicDetails: ", err)
		return nil, err
	}

	return &basicDetails, nil
}
