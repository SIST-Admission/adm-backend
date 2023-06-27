package service

import (
	"github.com/SIST-Admission/adm-backend/src/dto"
	"github.com/SIST-Admission/adm-backend/src/models"
	"github.com/SIST-Admission/adm-backend/src/repositories"
	"github.com/sirupsen/logrus"
)

type MeritListService struct{}

var meritListRepo repositories.MeritListRepository = repositories.MeritListRepository{}

func (meritListService *MeritListService) CreateMeritList(request *dto.CreateMeritListRequest) (*models.MeritList, *dto.Error) {
	logrus.Info("MeritListService.CreateMeritList")
	meritList, err := meritListRepo.CreateMeritList(request)
	if err != nil {
		logrus.Error("Failed to create merit list: ", err)
		return nil, &dto.Error{
			Code:    500,
			Message: "Failed to create merit list",
		}
	}

	if err = meritListRepo.AddStudents(&dto.AddStudentsToMeritListRequest{
		SubmissionIds: request.SubmissionIds,
		MeritListId:   meritList.Id,
	}); err != nil {
		logrus.Error("Failed to add students to merit list: ", err)
		return nil, &dto.Error{
			Code:    500,
			Message: "Failed to add students to merit list",
		}
	}

	return meritList, nil
}

func (meritListService *MeritListService) AddStudents(request *dto.AddStudentsToMeritListRequest) (*map[string]interface{}, *dto.Error) {
	logrus.Info("MeritListService.AddStudents")
	if err := meritListRepo.AddStudents(request); err != nil {
		logrus.Error("Failed to add students to merit list: ", err)
		return nil, &dto.Error{
			Code:    500,
			Message: "Failed to add students to merit list",
		}
	}

	return &map[string]interface{}{
		"code":    200,
		"success": true,
		"message": "Students added to merit list successfully",
	}, nil
}

func (meritListService *MeritListService) GetAllMeritLists(request *dto.GetAllMeritListsRequest) (*[]models.MeritList, *dto.Error) {
	logrus.Info("MeritListService.GetAllMeritLists")
	meritLists, err := meritListRepo.GetAllMeritLists(request)
	if err != nil {
		logrus.Error("Failed to get merit lists: ", err)
		return nil, &dto.Error{
			Code:    500,
			Message: "Failed to get merit lists",
		}
	}
	return meritLists, nil
}

func (meritListService *MeritListService) GetUnListedCandidates(request *dto.GetUnListedCandidatesRequest) (*[]models.Submission, *dto.Error) {
	logrus.Info("MeritListService.GetAllMeritLists")
	meritLists, err := meritListRepo.GetUnListedCandidatesRequest(request)
	if err != nil {
		logrus.Error("Failed to get merit lists: ", err)
		return nil, &dto.Error{
			Code:    500,
			Message: "Failed to get merit lists",
		}
	}
	return meritLists, nil
}

func (meritListService *MeritListService) GetListedCandidates(request *dto.GetListedCandidatesRequest) (*dto.GetListedCandidatesResponse, *dto.Error) {
	logrus.Info("MeritListService.GetListedCandidates")
	meritLists, err := meritListRepo.GetListedCandidates(request)
	if err != nil {
		logrus.Error("Failed to get merit lists: ", err)
		return nil, &dto.Error{
			Code:    500,
			Message: "Failed to get merit lists",
		}
	}
	return meritLists, nil
}
