package models

type Submission struct {
	Id             int    `gorm:"column:id;primary_key"`
	UserId         int    `gorm:"column:user_id"`
	ApplicationId  int    `gorm:"column:application_id"`
	DepartmentCode string `gorm:"column:department_code"`
	BatchCode      string `gorm:"column:batch_code"`
	Status         string `gorm:"column:status"`
	IsVerified     bool   `gorm:"column:is_verified"`
	IsAdmitted     bool   `gorm:"column:is_admitted"`
}

func (Submission) TableName() string {
	return "submissions"
}
