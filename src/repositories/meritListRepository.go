package repositories

import (
	"github.com/SIST-Admission/adm-backend/src/db"
	"github.com/SIST-Admission/adm-backend/src/dto"
	"github.com/SIST-Admission/adm-backend/src/models"
	"github.com/sirupsen/logrus"
)

type MeritListRepository struct{}

func (repo *MeritListRepository) CreateMeritList(p *dto.CreateMeritListRequest) (*models.MeritList, error) {
	logrus.Info("MeritListRepository.CreateMeritList")
	db := db.GetInstance()

	meritList := models.MeritList{
		DepartmentCode:  p.DepartmentCode,
		BatchCode:       p.BatchCode,
		PublishedDate:   p.PublishedDate,
		LastPaymentDate: p.LastPaymentDate,
		IsPublished:     p.IsPublished,
	}

	if err := db.Model(models.MeritList{}).Create(&meritList).Error; err != nil {
		logrus.Error("Failed to create merit list: ", err)
		return nil, err
	}

	return &meritList, nil
}

func (repo *MeritListRepository) AddStudents(p *dto.AddStudentsToMeritListRequest) error {
	logrus.Info("MeritListRepository.AddStudents")

	db := db.GetInstance()

	for _, sId := range p.SubmissionIds {
		if err := db.Model(models.Submission{}).Where("id = ?", sId).Update("merit_list_id", p.MeritListId).Error; err != nil {
			logrus.Error("Failed to add student to merit list: ", err)
			return err
		}
	}
	return nil
}

func (repo *MeritListRepository) GetAllMeritLists(p *dto.GetAllMeritListsRequest) (*[]models.MeritList, error) {
	logrus.Info("MeritListRepository.GetAllMeritLists")

	db := db.GetInstance()

	var meritLists []models.MeritList

	if p.DepartmentCode == "" {
		if err := db.Model(models.MeritList{}).Preload("Department").Preload("Batch").Find(&meritLists).Error; err != nil {
			logrus.Error("Failed to get merit lists: ", err)
			return nil, err
		}
		return &meritLists, nil
	} else {
		if err := db.Model(models.MeritList{}).Preload("Department").Preload("Batch").Where("department_code = ?", p.DepartmentCode).Find(&meritLists).Error; err != nil {
			logrus.Error("Failed to get merit lists: ", err)
			return nil, err
		}
	}

	return &meritLists, nil
}

func (repo *MeritListRepository) GetUnListedCandidatesRequest(p *dto.GetUnListedCandidatesRequest) (*[]models.Submission, *dto.Error) {
	logrus.Info("MeritListRepository.GetUnListedCandidatesRequest")
	db := db.GetInstance()

	// Get Batch Code
	var batch models.Batch
	if err := db.Model(models.Batch{}).Where("department_code = ? and start_year = ?", p.DepartmentCode, p.Year).First(&batch).Error; err != nil {
		logrus.Error("Failed to get batch: ", err)
		return nil, &dto.Error{Code: 500, Message: "Failed to get batch"}
	}

	// Get Submissions of the batch
	var submissions []models.Submission
	if err := db.Model(models.Submission{}).
		Preload("Application", "status = 'APPROVED'").Preload("Application.BasicDetails").Preload("Application.BasicDetails.PhotoDocument").
		Preload("Application.AcademicDetails").
		Preload("Application.AcademicDetails.ClassXIIDetails").Preload("Application.AcademicDetails.DiplomaDetails").
		Where("batch_code = ? and merit_list_id is null", batch.BatchCode).
		Find(&submissions).Error; err != nil {
		logrus.Error("Failed to get submissions: ", err)
		return nil, &dto.Error{Code: 500, Message: "Failed to get submissions"}
	}

	// Filter the submissions
	var filteredSubmissions []models.Submission
	for _, submission := range submissions {
		if submission.Application != nil {
			filteredSubmissions = append(filteredSubmissions, submission)
		}
	}
	return &filteredSubmissions, nil
}
