package models

type Submission struct {
	Id             int    `gorm:"column:id;primary_key" json:"id"`
	UserId         int    `gorm:"column:user_id" json:"userId"`
	ApplicationId  int    `gorm:"column:application_id" json:"applicationId"`
	DepartmentCode string `gorm:"column:department_code" json:"departmentCode"`
	BatchCode      string `gorm:"column:batch_code" json:"batchCode"`
	Status         string `gorm:"column:status" json:"status"`
	IsVerified     bool   `gorm:"column:is_verified" json:"isVerified"`
	IsAdmitted     bool   `gorm:"column:is_admitted" json:"isAdmitted"`
}

func (Submission) TableName() string {
	return "submissions"
}
