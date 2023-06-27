package dto

import "github.com/SIST-Admission/adm-backend/src/models"

type CreateMeritListRequest struct {
	DepartmentCode  string `json:"departmentCode"`
	Year            string `json:"year"`
	PublishedDate   string `json:"publishedDate"`
	LastPaymentDate string `json:"lastPaymentDate"`
	IsPublished     bool   `json:"isPublished"`
	SubmissionIds   []int  `json:"submissionIds"`
}

type AddStudentsToMeritListRequest struct {
	SubmissionIds []int `json:"submissionIds"`
	MeritListId   int   `json:"meritListId"`
}

type GetAllMeritListsRequest struct {
	DepartmentCode string `json:"departmentCode"`
}

type GetUnListedCandidatesRequest struct {
	DepartmentCode string `json:"departmentCode"`
	Year           string `json:"year"`
}

type GetListedCandidatesRequest struct {
	MeritListId int `json:"meritListId"`
}

type GetListedCandidatesResponse struct {
	MeritListDetails *models.MeritList    `json:"meritListDetails"`
	Submissions      []*models.Submission `json:"submissions"`
}
