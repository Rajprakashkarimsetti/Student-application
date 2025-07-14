package store

import (
	"database/sql"
	"github.com/rajprakash/student/models"
)

type StudentStores struct {
	db *sql.DB
}

func NewStudentStore(db *sql.DB) *StudentStores {
	return &StudentStores{db}
}

func (s *StudentStores) Post(student *models.Student) error {
	_, err := s.db.Query("INSERT INTO student(name,department,email) VALUES(?,?,?)", student.Name, student.Department, student.Email)
	if err != nil {
		return err
	}

	return nil
}

func (s *StudentStores) GetAll() ([]models.Student, error) {
	var students []models.Student
	rows, err := s.db.Query("SELECT * FROM student")
	if err != nil {
		return students, err
	}

	defer rows.Close()

	for rows.Next() {
		var student models.Student
		if err := rows.Scan(&student.ID, &student.Name, &student.Department, &student.Email); err != nil {
			return students, err
		}
		students = append(students, student)
	}
	return students, nil
}

func (s *StudentStores) GetById(id int) (*models.Student, error) {
	var student models.Student

	row := s.db.QueryRow("SELECT * FROM student WHERE id = ?", id)

	if err := row.Scan(&student.ID, &student.Name, &student.Department, &student.Email); err != nil {
		return &student, err
	}

	return &student, nil

}

func (s *StudentStores) Put(student *models.Student) (*models.Student, error) {
	_, err := s.db.Exec("UPDATE student SET name=?,department=?,email=? WHERE id = ?", student.Name, student.Department, student.Email, student.ID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *StudentStores) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM student WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
