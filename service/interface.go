package service

import "github.com/rajprakash/student/models"

type Student interface {
	Post(student *models.Student) error
	GetAll() ([]models.Student, error)
	GetByID(id int) (*models.Student, error)
	Put(student *models.Student) (*models.Student, error)
	Delete(id int) error
}
