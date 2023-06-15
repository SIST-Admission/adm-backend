package models

type AcademicDetails struct {
	Id               int      `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Class10SchoolId  int      `gorm:"column:class_10_school_id" json:"class10SchoolId"`
	Class12SchoolId  *int     `gorm:"column:class_12_school_id" json:"class12SchoolId"`
	DiplomaId        *int     `gorm:"column:diploma_id" json:"diplomaId"`
	JeeMainsRank     *int     `gorm:"column:jee_mains_rank" json:"jeeMainsRank"`
	JeeMainsMarks    *float32 `gorm:"column:jee_mains_marks" json:"jeeMainsMarks"`
	JeeAdvancedRank  *int     `gorm:"column:jee_advanced_rank" json:"jeeAdvancedRank"`
	JeeAdvancedMarks *float32 `gorm:"column:jee_advanced_marks" json:"jeeAdvancedMarks"`
	CuetRank         *int     `gorm:"column:cuet_rank" json:"cuetRank"`
	CuetMarks        *float32 `gorm:"column:cuet_marks" json:"cuetMarks"`
	MeritScore       float32  `gorm:"column:merit_score" json:"meritScore"`
	ClassXDetails    *School  `gorm:"foreignKey:class_10_school_id;AssociationForeignKey:id" json:"classXDetails"`
	ClassXIIDetails  *School  `gorm:"foreignKey:class_12_school_id;AssociationForeignKey:id" json:"classXIIDetails"`
	DiplomaDetails   *Diploma `gorm:"foreignKey:diploma_id;AssociationForeignKey:id" json:"diplomaDetails"`
	PaymentId        int      `gorm:"column:payment_id" json:"paymentId"`
}

func (AcademicDetails) TableName() string {
	return "academic_details"
}

type School struct {
	Id                  int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Board               string    `gorm:"column:board" json:"board"`
	YearOfPassing       string    `gorm:"column:year_of_passing" json:"yearOfPassing"`
	RollNumber          string    `gorm:"column:roll_number" json:"rollNumber"`
	Percentage          float32   `gorm:"column:percentage" json:"percentage"`
	MarksheetDocumentId int       `gorm:"column:document_id" json:"marksheetDocumentId"`
	MarksheetDocument   *Document `gorm:"foreignKey:document_id;AssociationForeignKey:id" json:"marksheetDocument"`
	SchoolName          string    `gorm:"column:school_name" json:"schoolName"`
	TotalMarks          float32   `gorm:"column:total_marks" json:"totalMarks"`
	MarksObtained       float32   `gorm:"column:marks_obtained" json:"marksObtained"`
}

func (School) TableName() string {
	return "school"
}

type Diploma struct {
	Id                  int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	CollegeName         string    `gorm:"column:college_name" json:"collegeName"`
	Department          string    `gorm:"column:department" json:"department"`
	YearOfPassing       string    `gorm:"column:year_of_passing" json:"yearOfPassing"`
	Cgpa                float32   `gorm:"column:cgpa" json:"cgpa"`
	MarksheetDocumentId int       `gorm:"column:document_id" json:"marksheetDocumentId"`
	MarksheetDocument   *Document `gorm:"foreignKey:document_id;AssociationForeignKey:id" json:"marksheetDocument"`
}

func (Diploma) TableName() string {
	return "diploma"
}
