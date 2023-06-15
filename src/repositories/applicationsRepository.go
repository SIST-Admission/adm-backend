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
	if err := db.Model(models.Application{}).Where("id = ?", appId).
		Preload("BasicDetails").Preload("BasicDetails.IdentityDocument").
		Preload("BasicDetails.PhotoDocument").
		Preload("BasicDetails.SignatureDocument").
		Preload("AcademicDetails").
		Preload("AcademicDetails.ClassXDetails").Preload("AcademicDetails.ClassXDetails.MarksheetDocument").
		Preload("AcademicDetails.ClassXIIDetails").Preload("AcademicDetails.ClassXIIDetails.MarksheetDocument").
		Preload("AcademicDetails.DiplomaDetails").Preload("AcademicDetails.DiplomaDetails.MarksheetDocument").
		Preload("PaymentDetails").
		First(&application).Error; err != nil {
		logrus.Error("ApplicationsRepository.GetApplicationDetails: ", err)
		return nil, err
	}

	// db.Preload("IdentityDocument").First(&application.BasicDetails)

	return &application, nil
}

func (repo *ApplicationsRepository) SaveAcademicDetails(userId, appId int, request *dto.SaveAcademicDetailsRequest) error {
	logrus.Info("ApplicationsRepository.SaveAcademicDetails")

	db := db.GetInstance()
	tx := db.Begin()
	defer tx.Rollback()

	// Save Class X Details
	classXMarksheetDocument := models.Document{
		UserID:       userId,
		Key:          request.Class10Details.Marksheet.Key,
		DocumentName: request.Class10Details.Marksheet.Key,
		MimeType:     request.Class10Details.Marksheet.MimeType,
		FileUrl:      request.Class10Details.Marksheet.Url,
		IsVerified:   false,
	}
	if err := tx.Model(models.Document{}).Save(&classXMarksheetDocument).Error; err != nil {
		logrus.Error("ApplicationsRepository.SaveAcademicDetails: ", err)
		return err
	}

	classXDetails := models.School{
		Board:               request.Class10Details.BoardName,
		YearOfPassing:       request.Class10Details.YearOfPass,
		RollNumber:          request.Class10Details.RollNumber,
		Percentage:          request.Class10Details.Percentage,
		TotalMarks:          0, // Not Storing Total Marks for Class X
		MarksObtained:       0, // Not Storing Marks Obtained for Class X``
		MarksheetDocumentId: classXMarksheetDocument.Id,
		SchoolName:          request.Class10Details.SchoolName,
	}

	if err := tx.Model(models.School{}).Save(&classXDetails).Error; err != nil {
		logrus.Error("ApplicationsRepository.SaveAcademicDetails: ", err)
		return err
	}

	var diplomaDetailsId *int = nil

	if request.DiplomaDetails != nil {
		// Save Diploma Details
		diplomaMarksheetDocument := models.Document{
			UserID:       userId,
			Key:          request.DiplomaDetails.Marksheet.Key,
			DocumentName: request.DiplomaDetails.Marksheet.Key,
			MimeType:     request.DiplomaDetails.Marksheet.MimeType,
			FileUrl:      request.DiplomaDetails.Marksheet.Url,
			IsVerified:   false,
		}

		if err := tx.Model(models.Document{}).Save(&diplomaMarksheetDocument).Error; err != nil {
			logrus.Error("ApplicationsRepository.SaveAcademicDetails: ", err)
			return err
		}

		diplomaDetails := models.Diploma{
			CollegeName:         request.DiplomaDetails.CollegeName,
			Department:          request.DiplomaDetails.Department,
			YearOfPassing:       request.DiplomaDetails.YearOfPass,
			Cgpa:                request.DiplomaDetails.Cgpa,
			MarksheetDocumentId: diplomaMarksheetDocument.Id,
		}

		if err := tx.Model(models.Diploma{}).Save(&diplomaDetails).Error; err != nil {
			logrus.Error("ApplicationsRepository.SaveAcademicDetails: ", err)
			return err
		}

		diplomaDetailsId = &diplomaDetails.Id
	}
	var classXIIDetailsId *int = nil
	if request.Class12Details != nil {
		// Save Class XII Details
		classXIIMarksheetDocument := models.Document{
			UserID:       userId,
			Key:          request.Class12Details.Marksheet.Key,
			DocumentName: request.Class12Details.Marksheet.Key,
			MimeType:     request.Class12Details.Marksheet.MimeType,
			FileUrl:      request.Class12Details.Marksheet.Url,
			IsVerified:   false,
		}

		if err := tx.Model(models.Document{}).Save(&classXIIMarksheetDocument).Error; err != nil {
			logrus.Error("ApplicationsRepository.SaveAcademicDetails: ", err)
			return err
		}

		classXIIDetails := models.School{
			Board:               request.Class12Details.BoardName,
			YearOfPassing:       request.Class12Details.YearOfPass,
			RollNumber:          request.Class12Details.RollNumber,
			Percentage:          request.Class12Details.Percentage,
			TotalMarks:          request.Class12Details.TotalMarks,
			MarksObtained:       request.Class12Details.Obtained,
			MarksheetDocumentId: classXIIMarksheetDocument.Id,
			SchoolName:          request.Class12Details.SchoolName,
		}

		if err := tx.Model(models.School{}).Save(&classXIIDetails).Error; err != nil {
			logrus.Error("ApplicationsRepository.SaveAcademicDetails: ", err)
			return err
		}

		classXIIDetailsId = &classXIIDetails.Id
	}

	var jeeMainsRank *int = nil
	var jeeMainsMarks *float32 = nil
	var jeeAdvancedRank *int = nil
	var jeeAdvancedMarks *float32 = nil
	var cuetRank *int = nil
	var cuetMarks *float32 = nil

	if request.JeeMainsDetails != nil {
		jeeMainsRank = &request.JeeMainsDetails.Rank
		jeeMainsMarks = &request.JeeMainsDetails.Score
	}

	if request.JeeAdvancedDetails != nil {
		jeeAdvancedRank = &request.JeeAdvancedDetails.Rank
		jeeAdvancedMarks = &request.JeeAdvancedDetails.Score
	}

	if request.CuetDetails != nil {
		cuetRank = &request.CuetDetails.Rank
		cuetMarks = &request.CuetDetails.Score
	}

	// Save Academic Details
	academicDetails := models.AcademicDetails{
		Class10SchoolId:  classXDetails.Id,
		Class12SchoolId:  classXIIDetailsId,
		JeeMainsRank:     jeeMainsRank,
		JeeMainsMarks:    jeeMainsMarks,
		JeeAdvancedRank:  jeeAdvancedRank,
		JeeAdvancedMarks: jeeAdvancedMarks,
		CuetRank:         cuetRank,
		CuetMarks:        cuetMarks,
		MeritScore:       float32(0),
		DiplomaId:        diplomaDetailsId,
	}

	if err := tx.Model(models.AcademicDetails{}).Save(&academicDetails).Error; err != nil {
		logrus.Error("ApplicationsRepository.SaveAcademicDetails: ", err)
		return err
	}

	// Save Application
	if err := tx.Model(models.Application{}).Where("id = ?", appId).Updates(models.Application{AcademicDetailsId: academicDetails.Id}).Error; err != nil {
		logrus.Error("ApplicationsRepository.SaveAcademicDetails: ", err)
		return err
	}

	if err := tx.Commit().Error; err != nil {
		logrus.Error("ApplicationsRepository.SaveAcademicDetails: ", err)
		return err
	}

	return nil
}
