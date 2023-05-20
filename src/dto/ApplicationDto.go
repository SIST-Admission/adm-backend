package dto

type StartApplicationRequst struct {
	ApplicationType string `json:"applicationType"`
}

type StartApplicationResponse struct {
	Id                   int    `json:"id"`
	BasicDetailsId       int    `json:"basicDetailsId"`
	AcademicDetailsId    int    `json:"academicDetailsId"`
	PaymetId             int    `json:"paymetId"`
	ApplicationType      string `json:"applicationType"`
	ApplicationStartDate string `json:"applicationStartDate"`
	Status               string `json:"status"`
}
