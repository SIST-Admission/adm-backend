package models

type Application struct {
	Id                   int              `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	BasicDetailsId       int              `gorm:"column:basic_details_id;default:null;" json:"basic_details_id"`
	AcademicDetailsId    int              `gorm:"column:academic_details_id;default:null" json:"academic_details_id"`
	PaymentId            int              `gorm:"column:payment_id;default:null" json:"payment_id"`
	ApplicationType      string           `gorm:"column:application_type" json:"application_type"`
	ApplicationStartDate string           `gorm:"column:application_start_date" json:"application_start_date"`
	Status               string           `gorm:"column:status" json:"status"`
	BasicDetails         *BasicDetails    `gorm:"foreignKey:basic_details_id;AssociationForeignKey:id" json:"basic_details"`
	AcademicDetails      *AcademicDetails `gorm:"foreignKey:academic_details_id;AssociationForeignKey:id" json:"academic_details"`
}

func (Application) TableName() string {
	return "applications"
}
