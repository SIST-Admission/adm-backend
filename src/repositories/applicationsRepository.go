package repositories

import (
	"net/http"
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
		Email:               payload.Email,
		Phone:               payload.Phone,
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
		Address:             payload.Address,
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
		Email:               payload.Email,
		Phone:               payload.Phone,
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
		Address:             payload.Address,
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
		Preload("Submissions").
		Preload("Submissions.MeritList").
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

func (repo *ApplicationsRepository) GetAllApplications(p *dto.GetAllApplicationsRequest) ([]*models.Application, error) {
	logrus.Info("ApplicationsRepository.GetAllApplications: ", p.Status)
	var applications []*models.Application
	db := db.GetInstance()

	if p.Status != "" {
		if err := db.Model(models.Application{}).
			Preload("Submissions").
			Preload("BasicDetails").
			Preload("BasicDetails.PhotoDocument").
			Where("status = ?", p.Status).
			Find(&applications).Error; err != nil {
			logrus.Error("ApplicationsRepository.GetAllApplications: ", err)
			return nil, err
		}
	} else {
		if err := db.Model(models.Application{}).
			Preload("Submissions").
			Preload("BasicDetails").
			Preload("BasicDetails.PhotoDocument").
			Where("status != ?", "DRAFT").
			Find(&applications).Error; err != nil {
			logrus.Error("ApplicationsRepository.GetAllApplications: ", err)
			return nil, err
		}
	}

	return applications, nil
}

func (repo *ApplicationsRepository) UpdateDocumentStatus(req *dto.UpdateDocumentStatusRequest) *dto.Error {
	logrus.Info("ApplicationsRepository.UpdateDocumentStatus: ")
	db := db.GetInstance()
	if err := db.Model(models.Document{}).Where("id = ?", req.DocumentId).Updates(models.Document{IsVerified: req.IsVerified, Status: req.Status}).Error; err != nil {
		logrus.Error("ApplicationsRepository.UpdateDocumentStatus: ", err)
		return &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error updating document status",
		}
	}
	return nil
}

func (repo *ApplicationsRepository) UpdateApplicationStatus(req *dto.UpdateApplicationRequest) *dto.Error {
	logrus.Info("ApplicationsRepository.UpdateApplicationStatus: ")
	db := db.GetInstance()
	if err := db.Model(models.Application{}).Where("id = ?", req.Id).Updates(models.Application{Status: req.Status}).Error; err != nil {
		logrus.Error("ApplicationsRepository.UpdateApplicationStatus: ", err)
		return &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error updating application status",
		}
	}
	return nil
}

func (repo *ApplicationsRepository) GetApplicationStats() (*map[string]interface{}, *dto.Error) {
	db := db.GetInstance()
	var pendingApplications int64
	var approvedApplications int64
	var rejectedApplications int64
	var fesherApplications int64
	var lateralApplications int64

	if err := db.Model(models.Application{}).Where("status = ?", "SUBMITTED").Count(&pendingApplications).Error; err != nil {
		logrus.Error("ApplicationsRepository.UpdateApplicationStatus: ", err)
		return nil, &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error Finding application status",
		}
	}

	if err := db.Model(models.Application{}).Where("status = ?", "APPROVED").Count(&approvedApplications).Error; err != nil {
		logrus.Error("ApplicationsRepository.UpdateApplicationStatus: ", err)
		return nil, &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error Finding application status",
		}
	}

	if err := db.Model(models.Application{}).Where("status = ?", "REJECTED").Count(&rejectedApplications).Error; err != nil {
		logrus.Error("ApplicationsRepository.UpdateApplicationStatus: ", err)
		return nil, &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error Finding application status",
		}
	}

	if err := db.Model(models.Application{}).Where("application_type = ?", "FRESHER").Count(&fesherApplications).Error; err != nil {
		logrus.Error("ApplicationsRepository.UpdateApplicationStatus: ", err)
		return nil, &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error Finding application status",
		}
	}

	if err := db.Model(models.Application{}).Where("application_type = ?", "LATERAL").Count(&lateralApplications).Error; err != nil {
		logrus.Error("ApplicationsRepository.UpdateApplicationStatus: ", err)
		return nil, &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error Finding application status",
		}
	}

	var admittedStudents int64
	var cseAdmittedStudents int64
	var cvlAdmittedStudents int64
	var totalSubmissions int64

	if err := db.Model(models.Submission{}).Count(&totalSubmissions).Error; err != nil {
		logrus.Error("ApplicationsRepository.UpdateApplicationStatus: ", err)
		return nil, &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error Finding application status",
		}
	}

	if err := db.Model(models.Submission{}).Where("is_admitted = ?", true).Count(&admittedStudents).Error; err != nil {
		logrus.Error("ApplicationsRepository.UpdateApplicationStatus: ", err)
		return nil, &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error Finding application status",
		}
	}

	if err := db.Model(models.Submission{}).Where("is_admitted = ? AND department_code = ?", true, "CSE").Count(&cseAdmittedStudents).Error; err != nil {
		logrus.Error("ApplicationsRepository.UpdateApplicationStatus: ", err)
		return nil, &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error Finding application status",
		}
	}

	if err := db.Model(models.Submission{}).Where("is_admitted = ? AND department_code = ?", true, "CVL").Count(&cvlAdmittedStudents).Error; err != nil {
		logrus.Error("ApplicationsRepository.UpdateApplicationStatus: ", err)
		return nil, &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error Finding application status",
		}
	}

	var csePendingAdmissions int64
	var cvlPendingAdmissions int64
	if err := db.Model(models.Submission{}).Where("is_admitted = ? AND department_code = ? and merit_list_id is not null", false, "CSE").Count(&csePendingAdmissions).Error; err != nil {
		logrus.Error("ApplicationsRepository.UpdateApplicationStatus: ", err)
		return nil, &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error Finding cse pending admissions",
		}
	}

	if err := db.Model(models.Submission{}).Where("is_admitted = ? AND department_code = ? and merit_list_id is not null", false, "CVL").Count(&cvlPendingAdmissions).Error; err != nil {
		logrus.Error("ApplicationsRepository.UpdateApplicationStatus: ", err)
		return nil, &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error Finding cvl pending admissions",
		}
	}

	var totalMeritList int64
	var cseMeritList int64
	var cvlMeritList int64
	var totalCseListedCandidates int64
	var totalCvlListedCandidates int64

	if err := db.Model(models.MeritList{}).Count(&totalMeritList).Error; err != nil {
		logrus.Error("ApplicationsRepository.UpdateApplicationStatus: ", err)
		return nil, &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error counting total merit list",
		}
	}

	if err := db.Model(models.MeritList{}).Where("department_code = ?", "CSE").Count(&cseMeritList).Error; err != nil {
		logrus.Error("ApplicationsRepository.UpdateApplicationStatus: ", err)
		return nil, &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error counting cse merit list",
		}
	}

	if err := db.Model(models.MeritList{}).Where("department_code = ?", "CVL").Count(&cvlMeritList).Error; err != nil {
		logrus.Error("ApplicationsRepository.UpdateApplicationStatus: ", err)
		return nil, &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: "error counting cvl merit list",
		}
	}

	if err := db.Model(models.Submission{}).Where("department_code = ? AND merit_list_id is not null", "CSE").Count(&totalCseListedCandidates).Error; err != nil {
		logrus.Error("ApplicationsRepository.UpdateApplicationStatus: ", err)
		return nil, &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: "error counting total cse listed candidates",
		}
	}

	if err := db.Model(models.Submission{}).Where("department_code = ? AND merit_list_id is not null", "CVL").Count(&totalCvlListedCandidates).Error; err != nil {
		logrus.Error("ApplicationsRepository.UpdateApplicationStatus: ", err)
		return nil, &dto.Error{
			Code:    http.StatusInternalServerError,
			Message: "error counting total cvl listed candidates",
		}
	}

	var totalCseAdmittedCandidates int64

	var resp = map[string]interface{}{
		"pending_applications":  pendingApplications,
		"approved_applications": approvedApplications,
		"rejected_applications": rejectedApplications,
		"fresh_applications":    fesherApplications,
		"lateral_applications":  lateralApplications,
		"total_submissions":     totalSubmissions,
		"total_admitted":        admittedStudents,
		"cse_admitted":          cseAdmittedStudents,
		"cvl_admitted":          cvlAdmittedStudents,
		"cse_pending_admission": csePendingAdmissions,
		"cvl_pending_admission": cvlPendingAdmissions,
		"total_merit_list":      totalMeritList,
		"cse_merit_list":        cseMeritList,
		"cvl_merit_list":        cvlMeritList,
		"cse_listed":            totalCseListedCandidates,
		"cvl_listed":            totalCvlListedCandidates,
		"cse_admitted_list":     totalCseAdmittedCandidates,
	}

	return &resp, nil
}
