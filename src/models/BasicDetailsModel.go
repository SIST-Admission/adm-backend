package models

type BasicDetails struct {
	Id                  int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Name                string    `gorm:"column:name" json:"name"`
	DoB                 string    `gorm:"column:dob" json:"dob"`
	Gender              string    `gorm:"column:gender" json:"gender"`
	Category            string    `gorm:"column:category" json:"category"`
	IsCoI               bool      `gorm:"column:is_coi" json:"isCoi"`
	IsPwD               bool      `gorm:"column:is_pwd" json:"isPwd"`
	FatherName          string    `gorm:"column:father_name" json:"fatherName"`
	MotherName          string    `gorm:"column:mother_name" json:"motherName"`
	Nationality         string    `gorm:"column:nationality" json:"nationality"`
	IdentityType        string    `gorm:"column:identity_type" json:"identityType"`
	IdentityNumber      string    `gorm:"column:identity_number" json:"identityNumber"`
	IdentityDocumentId  int       `gorm:"column:identity_document_id" json:"identityDocumentId"`
	PhotoDocumentId     int       `gorm:"column:photo_document_id" json:"photoDocumentId"`
	SignatureDocumentId int       `gorm:"column:signature_document_id" json:"signatureDocumentId"`
	IdentityDocument    *Document `gorm:"foreignKey:identity_document_id;AssociationForeignKey:id" json:"identityDocument"`
	PhotoDocument       *Document `gorm:"foreignKey:photo_document_id;AssociationForeignKey:id" json:"photoDocument"`
	SignatureDocument   *Document `gorm:"foreignKey:signature_document_id;AssociationForeignKey:id" json:"signatureDocument"`
}

func (BasicDetails) TableName() string {
	return "basic_details"
}
