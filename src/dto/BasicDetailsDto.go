package dto

type SaveBasicDetailsRequest struct {
	Name                     string `json:"name"`
	Email                    string `json:"email"`
	Phone                    string `json:"phone"`
	DoB                      string `json:"dob"`
	Gender                   string `json:"gender"`
	Category                 string `json:"category"`
	IsCoI                    bool   `json:"isCoi"`
	IsPwD                    bool   `json:"isPwd"`
	FatherName               string `json:"fatherName"`
	MotherName               string `json:"motherName"`
	Nationality              string `json:"nationality"`
	IdentityType             string `json:"identityType"`
	IdentityNumber           string `json:"identityNumber"`
	IdentityDocumentKey      string `json:"identityDocumentKey"`
	IdentityDocumentMimeType string `json:"identityDocumentMimeType"`
	IdentityDocumentUrl      string `json:"identityDocumentUrl"`
	PassportPhotoKey         string `json:"passportPhotoKey"`
	PassportPhotoMimeType    string `json:"passportPhotoMimeType"`
	PassportPhotoUrl         string `json:"passportPhotoUrl"`
	SignatureKey             string `json:"signatureKey"`
	SignatureMimeType        string `json:"signatureMimeType"`
	SignatureUrl             string `json:"signatureUrl"`
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
