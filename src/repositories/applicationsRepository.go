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

func (repo *ApplicationsRepository) SaveBasicDetails(userId, appId int, payload *dto.SaveBasicDetailsRequest) (*models.BasicDetails, error) {
	logrus.Info("ApplicationsRepository.SaveBasicDetails")
	db := db.GetInstance()
	tx := db.Begin()
	defer tx.Rollback()

	// Insert ID proof document to documents table
	identityDocument := models.Document{
		DocumentName: payload.IdentityDocumentKey,
		MimeType:     payload.IdentityDocumentMimeType,
		Key:          payload.IdentityDocumentKey,
		FileUrl:      payload.IdentityDocumentUrl,
		UserID:       userId,
		IsVerified:   false,
	}

	if err := tx.Model(models.Document{}).Save(&identityDocument).Error; err != nil {
		logrus.Error("ApplicationsRepository.SaveBasicDetails: ", err)
		return nil, err
	}

	// insert Passport size photo to documents table
	photoDocument := models.Document{
		DocumentName: payload.PassportPhotoKey,
		MimeType:     payload.PassportPhotoMimeType,
		Key:          payload.PassportPhotoKey,
		FileUrl:      payload.PassportPhotoUrl,
		UserID:       userId,
		IsVerified:   false,
	}

	if err := tx.Model(models.Document{}).Save(&photoDocument).Error; err != nil {
		logrus.Error("ApplicationsRepository.SaveBasicDetails: ", err)
		return nil, err
	}

	// Insert Signature to documents table
	signatureDocument := models.Document{
		DocumentName: payload.SignatureKey,
		MimeType:     payload.SignatureMimeType,
		Key:          payload.SignatureKey,
		FileUrl:      payload.SignatureUrl,
		UserID:       userId,
		IsVerified:   false,
	}

	if err := tx.Model(models.Document{}).Save(&signatureDocument).Error; err != nil {
		logrus.Error("ApplicationsRepository.SaveBasicDetails: ", err)
		return nil, err
	}

	basicDetails := models.BasicDetails{
		Name:                payload.Name,
		DoB:                 payload.DoB, // Accepted format: "YYYY-MM-DD"
		Gender:              payload.Gender,
		Category:            payload.Category,
		IsCoI:               payload.IsCoI,
		IsPwD:               payload.IsPwD,
		FatherName:          payload.FatherName,
		MotherName:          payload.MotherName,
		Nationality:         payload.Nationality,
		IdentityType:        payload.IdentityType,
		IdentityNumber:      payload.IdentityNumber,
		IdentityDocumentId:  identityDocument.Id,
		PhotoDocumentId:     photoDocument.Id,
		SignatureDocumentId: signatureDocument.Id,
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

func (repo *ApplicationsRepository) UpdateBasicDetails(userId, basicDetailsId int, payload *dto.SaveBasicDetailsRequest) (*models.BasicDetails, error) {
	logrus.Info("ApplicationsRepository.UpdateBasicDetails")
	db := db.GetInstance()
	tx := db.Begin()
	defer tx.Rollback()

	// Insert ID proof document to documents table
	identityDocument := models.Document{
		DocumentName: payload.IdentityDocumentKey,
		MimeType:     payload.IdentityDocumentMimeType,
		Key:          payload.IdentityDocumentKey,
		FileUrl:      payload.IdentityDocumentUrl,
		UserID:       userId,
		IsVerified:   false,
	}

	if err := tx.Model(models.Document{}).Save(&identityDocument).Error; err != nil {
		logrus.Error("ApplicationsRepository.SaveBasicDetails: ", err)
		return nil, err
	}

	// insert Passport size photo to documents table
	photoDocument := models.Document{
		DocumentName: payload.PassportPhotoKey,
		MimeType:     payload.PassportPhotoMimeType,
		Key:          payload.PassportPhotoKey,
		FileUrl:      payload.PassportPhotoUrl,
		UserID:       userId,
		IsVerified:   false,
	}

	if err := tx.Model(models.Document{}).Save(&photoDocument).Error; err != nil {
		logrus.Error("ApplicationsRepository.SaveBasicDetails: ", err)
		return nil, err
	}

	// Insert Signature to documents table
	signatureDocument := models.Document{
		DocumentName: payload.SignatureKey,
		MimeType:     payload.SignatureMimeType,
		Key:          payload.SignatureKey,
		FileUrl:      payload.SignatureUrl,
		UserID:       userId,
		IsVerified:   false,
	}

	if err := tx.Model(models.Document{}).Save(&signatureDocument).Error; err != nil {
		logrus.Error("ApplicationsRepository.SaveBasicDetails: ", err)
		return nil, err
	}

	var basicDetails models.BasicDetails = models.BasicDetails{
		Name:                payload.Name,
		DoB:                 payload.DoB, // Accepted format: "YYYY-MM-DD"
		Gender:              payload.Gender,
		Category:            payload.Category,
		IsCoI:               payload.IsCoI,
		IsPwD:               payload.IsPwD,
		FatherName:          payload.FatherName,
		MotherName:          payload.MotherName,
		Nationality:         payload.Nationality,
		IdentityType:        payload.IdentityType,
		IdentityNumber:      payload.IdentityNumber,
		IdentityDocumentId:  identityDocument.Id,
		PhotoDocumentId:     photoDocument.Id,
		SignatureDocumentId: signatureDocument.Id,
	}

	if err := tx.Model(models.BasicDetails{}).Where("id = ?", basicDetailsId).Updates(basicDetails).Error; err != nil {
		logrus.Error("ApplicationsRepository.UpdateBasicDetails: ", err)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		logrus.Error("ApplicationsRepository.UpdateBasicDetails: ", err)
		return nil, err
	}
	return &basicDetails, nil
}

func (repo *ApplicationsRepository) GetApplicationDetails(appId int) (*models.Application, error) {
	logrus.Info("ApplicationsRepository.GetApplicationDetails")

	db := db.GetInstance()
	var application models.Application
	if err := db.Model(models.Application{}).Where("id = ?", appId).Preload("BasicDetails").Preload("BasicDetails.IdentityDocument").Preload("BasicDetails.PhotoDocument").Preload("BasicDetails.SignatureDocument").First(&application).Error; err != nil {
		logrus.Error("ApplicationsRepository.GetApplicationDetails: ", err)
		return nil, err
	}

	// db.Preload("IdentityDocument").First(&application.BasicDetails)

	return &application, nil
}
