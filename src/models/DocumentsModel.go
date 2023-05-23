package models

type Document struct {
	Id           int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	DocumentName string `gorm:"column:document_name" json:"document_name"`
	MimeType     string `gorm:"column:mime_type" json:"mime_type"`
	Key          string `gorm:"column:key" json:"key"`
	FileUrl      string `gorm:"column:file_url" json:"file_url"`
	UserID       int    `gorm:"column:user_id" json:"user_id"`
	IsVerified   bool   `gorm:"column:is_verified" json:"is_verified"`
}

func (Document) TableName() string {
	return "documents"
}
