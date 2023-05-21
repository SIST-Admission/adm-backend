package dto

type SaveBasicDetailsRequest struct {
	Name               string `json:"name"`
	DoB                string `json:"dob"`
	Gender             string `json:"gender"`
	Category           string `json:"category"`
	IsCoI              bool   `json:"isCoi"`
	IsPwD              bool   `json:"isPwd"`
	FatherName         string `json:"fatherName"`
	MotherName         string `json:"motherName"`
	Nationality        string `json:"nationality"`
	IdentityType       string `json:"identityType"`
	IdentityNumber     string `json:"identityNumber"`
	IdentityDocumentId int    `json:"identityDocumentId"`
}

type SaveBasicDetailsResponse struct {
	Id                 int    `json:"id"`
	Name               string `json:"name"`
	DoB                string `json:"dob"`
	Gender             string `json:"gender"`
	Category           string `json:"category"`
	IsCoI              bool   `json:"isCoi"`
	IsPwD              bool   `json:"isPwd"`
	FatherName         string `json:"fatherName"`
	MotherName         string `json:"motherName"`
	Nationality        string `json:"nationality"`
	IdentityType       string `json:"identityType"`
	IdentityNumber     string `json:"identityNumber"`
	IdentityDocumentId int    `json:"identityDocumentId"`
}
