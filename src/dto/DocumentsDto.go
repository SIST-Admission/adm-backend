package dto

type UpdateDocumentStatusRequest struct {
	DocumentId int    `json:"documentId"`
	Status     string `json:"status"`
	IsVerified bool   `json:"isVerified"`
}
