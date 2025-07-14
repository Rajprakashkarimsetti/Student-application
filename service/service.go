package service

import (
	"errors"
	"github.com/rajprakash/student/models"
	"github.com/rajprakash/student/store"
)

type StudentService struct {
	Student store.Student
}

func NewStudentService(student store.Student) *StudentService {
	return &StudentService{student}
}

func (s *StudentService) Post(student *models.Student) error {
	if student.Name == "" || student.Department == "" || student.Email == "" {
		return errors.New("student name and department are required")
	}

	return s.Student.Post(student)
}

func (s *StudentService) GetAll() ([]models.Student, error) {
	return s.Student.GetAll()
}

func (s *StudentService) GetByID(id int) (*models.Student, error) {
	if id == 0 {
		return nil, errors.New("student id is required")
	}

	return s.Student.GetById(id)
}

func (s *StudentService) Put(student *models.Student) (*models.Student, error) {
	if student.ID == 0 {
		return nil, errors.New("student id is required")
	}

	_, err := s.Student.Put(student)
	res, err := s.Student.GetById(student.ID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *StudentService) Delete(id int) error {
	if id == 0 {
		return errors.New("student id is required")
	}

	return s.Student.Delete(id)
}
