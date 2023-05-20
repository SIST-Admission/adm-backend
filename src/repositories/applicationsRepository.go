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

func (repo *ApplicationsRepository) SaveBasicDetails(payload dto.SaveBasicDetailsRequest) (*models.BasicDetails, error) {
	logrus.Info("ApplicationsRepository.SaveBasicDetails")
	// db := db.GetInstance()
	return nil, nil
}
