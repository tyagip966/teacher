package models

type Teacher struct {
	Name           string `json:"name"`
	Age            string `json:"age"`
	Department     string `json:"department"`
	IsHead         bool   `json:"is_head"`
	SchoolCode     int    `json:"school_code"`
	IdentityNumber int    `json:"identityNumber"`
}

type TeacherService interface {
	AddTeacher(input Teacher) (*Teacher, error)
	GetTeacher(id int) (*Teacher, error)
	GetTeachers(schoolCode int) ([]Teacher, error)
	DeleteTeacher(id int) (*Teacher, error)
	UpdateTeacher(id int, input Teacher) (*Teacher, error)
}
