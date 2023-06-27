package models

type MeritList struct {
	Id              int         `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	DepartmentCode  string      `gorm:"column:department_code" json:"departmentCode"`
	BatchCode       string      `gorm:"column:batch_code" json:"batchCode"`
	PublishedDate   string      `gorm:"column:published_date" json:"publishedDate"`
	LastPaymentDate string      `gorm:"column:last_payment_date" json:"lastPaymentDate"`
	IsPublished     bool        `gorm:"column:is_published" json:"isPublished"`
	Department      *Department `gorm:"foreignKey:department_code;references:department_code" json:"department"`
	Batch           *Batch      `gorm:"foreignKey:batch_code;references:batch_code" json:"batch"`
}

func (MeritList) TableName() string {
	return "merit_lists"
}
