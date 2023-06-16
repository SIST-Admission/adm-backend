package models

type Department struct {
	DepartmentCode string `gorm:"column:department_code; primary_key" json:"departmentCode"`
	DepartmentName string `gorm:"column:department_name" json:"departmentName"`
}

func (Department) TableName() string {
	return "departments"
}
