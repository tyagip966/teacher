package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"teacher/constants"
	"teacher/models"
)

//TeacherService ...
type TeacherService struct {
	Database *mgo.Session
}

//AddTeacher ...
func (s TeacherService) AddTeacher(input models.Teacher) (*models.Teacher, error) {
	err := s.Database.DB(constants.DataBaseName).C(constants.CollectionName).Insert(&input)
	if err != nil {
		return nil,err
	}
	return &input, nil
}

//GetTeacher ...
func (s TeacherService) GetTeacher(id int) (*models.Teacher, error) {
	var response models.Teacher
	err := s.Database.DB(constants.DataBaseName).C(constants.CollectionName).Find(bson.M{"_id":id}).One(&response)
	if err != nil {
		return nil,err
	}
	return &response, nil
}

//GetTeachers ...
func (s TeacherService) GetTeachers(schoolCode int) ([]models.Teacher, error) {
	var response []models.Teacher
	err := s.Database.DB(constants.DataBaseName).C(constants.CollectionName).Find(bson.M{"school_code":schoolCode}).All(&response)
	if err != nil {
		return nil,err
	}
	return response, nil
}

//DeleteTeacher ...
func (s TeacherService) DeleteTeacher(id int) (*models.Teacher, error) {
	var response models.Teacher
	err := s.Database.DB(constants.DataBaseName).C(constants.CollectionName).FindId(id).One(&response)
	if err != nil {
		return nil,err
	}
	err = s.Database.DB(constants.DataBaseName).C(constants.CollectionName).RemoveId(id)
	if err != nil {
		return nil,err
	}
	return &response, nil
}

//UpdateTeacher ...
func (s TeacherService) UpdateTeacher(id int, input models.Teacher) (*models.Teacher, error) {
	err := s.Database.DB(constants.DataBaseName).C(constants.CollectionName).UpdateId(id,&input)
	if err != nil {
		return nil,err
	}
	var response models.Teacher
	err = s.Database.DB(constants.DataBaseName).C(constants.CollectionName).FindId(id).One(&response)
	if err != nil {
		return nil,err
	}
	return &response, nil
}
