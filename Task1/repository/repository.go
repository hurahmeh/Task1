package repository

import (
	"Task1/entity"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type StudentRepository struct {
	DB *gorm.DB
}

// GetAll students
func (s StudentRepository) GetAll() []*entity.Student {
	var students []*entity.Student
	if err := s.DB.Find(&students).Error; err != nil {
		return nil
	}
	return students
}

// Create student
func (s StudentRepository) Create(u *entity.Student) error {

	result := s.DB.Create(&u)
	fmt.Println(result.Row())
	s.DB.Save(&u)
	return nil
}

// Delete student
func (s StudentRepository) Delete(id int) error {

	deleted := s.DB.Delete(&entity.Student{}, id)
	if deleted.RowsAffected == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Student not found")

	}

	return nil
}

// Update student
func (s StudentRepository) Update(id int, u *entity.Student) (error, *entity.Student) {

	student := new(entity.Student)
	if err := s.DB.First(&student, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Student not found"), nil
	}
	if u.FirstName == "" || u.LastName == "" {
		return echo.NewHTTPError(http.StatusNotFound, "first and last name required"), nil
	}

	student.FirstName = u.LastName
	student.LastName = u.LastName
	s.DB.Save(student)
	return nil, student
}

// Patch Update student
func (s StudentRepository) Patch(id int, u *entity.Student) (error, *entity.Student) {

	student := new(entity.Student)
	if err := s.DB.First(&student, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "student not found"), nil
	}

	if u.FirstName != "" {
		student.FirstName = u.FirstName
	}
	if u.LastName != "" {
		student.LastName = u.LastName
	}
	s.DB.Save(student)
	return nil, student
}
