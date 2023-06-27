package models

type Submission struct {
	Id                int          `gorm:"column:id;primary_key" json:"id"`
	UserId            int          `gorm:"column:user_id" json:"userId"`
	ApplicationId     int          `gorm:"column:application_id" json:"applicationId"`
	DepartmentCode    string       `gorm:"column:department_code" json:"departmentCode"`
	BatchCode         string       `gorm:"column:batch_code" json:"batchCode"`
	Status            string       `gorm:"column:status" json:"status"`
	IsVerified        bool         `gorm:"column:is_verified" json:"isVerified"`
	IsAdmitted        bool         `gorm:"column:is_admitted" json:"isAdmitted"`
	MeritListId       *int         `gorm:"column:merit_list_id" json:"meritListId"`
	MeritList         *MeritList   `gorm:"foreignKey:merit_list_id;references:id" json:"meritList"`
	Application       *Application `gorm:"foreignKey:application_id;references:id" json:"application"`
	PaymentId         *int         `gorm:"column:payment_id" json:"paymentId"`
	Payment           *Payment     `gorm:"foreignKey:payment_id;references:id" json:"payment"`
	BatchDetails      *Batch       `gorm:"foreignKey:batch_code;references:batch_code" json:"batchDetails"`
	DepartmentDetails *Department  `gorm:"foreignKey:department_code;references:department_code" json:"departmentDetails"`
}

func (Submission) TableName() string {
	return "submissions"
}
