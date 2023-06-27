package dto

type CreateMeritListRequest struct {
	DepartmentCode  string `json:"departmentCode"`
	BatchCode       string `json:"batchCode"`
	PublishedDate   string `json:"publishedDate"`
	LastPaymentDate string `json:"lastPaymentDate"`
	IsPublished     bool   `json:"isPublished"`
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
