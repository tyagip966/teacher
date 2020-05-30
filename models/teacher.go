package models

import "github.com/tyagip966/common-repo/models"

type TeacherService interface {
	AddTeacher(input models.Teacher) (*models.Teacher, error)
	GetTeacher(id int) (*models.Teacher, error)
	GetTeachers(schoolCode int) ([]models.Teacher, error)
	DeleteTeacher(id int) (*models.Teacher, error)
	UpdateTeacher(id int, input models.Teacher) (*models.Teacher, error)
}
