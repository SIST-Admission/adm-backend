package dto

type Class10Details struct {
	RollNumber string            `json:"rollNumber"`
	SchoolName string            `json:"schoolName"`
	BoardName  string            `json:"boardName"`
	YearOfPass string            `json:"yearOfPass"`
	Percentage float32           `json:"percentage"`
	TotalMarks float32           `json:"totalMarks"`
	Obtained   float32           `json:"obtained"`
	Subjects   []Class10Subjects `json:"subjects"`
	Marksheet  MarksheetDocument `json:"marksheet"`
}

type Class10Subjects struct {
	SubjectName string  `json:"subjectName"`
	TotalMarks  float32 `json:"totalMarks"`
	Obtained    float32 `json:"obtained"`
}

type Class12Details struct {
	RollNumber string            `json:"rollNumber"`
	SchoolName string            `json:"schoolName"`
	BoardName  string            `json:"boardName"`
	Stream     string            `json:"stream"`
	YearOfPass string            `json:"yearOfPass"`
	Percentage float32           `json:"percentage"`
	TotalMarks float32           `json:"totalMarks"`
	Obtained   float32           `json:"obtained"`
	Subjects   []Class12Subjects `json:"subjects"`
	Marksheet  MarksheetDocument `json:"marksheet"`
}

type Class12Subjects struct {
	SubjectName string  `json:"subjectName"`
	TotalMarks  float32 `json:"totalMarks"`
	Obtained    float32 `json:"obtained"`
}

type CuetDetails struct {
	YearOfPass string  `json:"yearOfPass"`
	Score      float32 `json:"score"`
	Rank       int     `json:"rank"`
}

type JeeMainsDetails struct {
	YearOfPass string  `json:"yearOfPass"`
	Score      float32 `json:"score"`
	Rank       int     `json:"rank"`
}

type JeeAdvancedDetails struct {
	YearOfPass string  `json:"yearOfPass"`
	Score      float32 `json:"score"`
	Rank       int     `json:"rank"`
}

type SaveAcademicDetailsRequest struct {
	Class10Details     Class10Details      `json:"class10Details"`
	Class12Details     *Class12Details     `json:"class12Details"`
	CuetDetails        *CuetDetails        `json:"cuetDetails"`
	JeeMainsDetails    *JeeMainsDetails    `json:"jeeMainsDetails"`
	JeeAdvancedDetails *JeeAdvancedDetails `json:"jeeAdvancedDetails"`
}

type MarksheetDocument struct {
	Key      string `json:"key"`
	MimeType string `json:"mimeType"`
	Url      string `json:"url"`
}
