package models

type Batch struct {
	BatchCode      string     `gorm:"column:batch_code; primary_key" json:"batchCode"`
	BatchName      string     `gorm:"column:batch_name" json:"batchName"`
	DepartmentCode string     `gorm:"column:department_code" json:"departmentCode"`
	StartYear      int        `gorm:"column:start_year" json:"startYear"`
	EndYear        int        `gorm:"column:end_year" json:"endYear"`
	Department     Department `gorm:"foreignKey:department_code;references:department_code" json:"department"`
}
