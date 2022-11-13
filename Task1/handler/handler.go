package handler

import (
	"Task1/entity"
	"Task1/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type StudentHandler struct {
	StudentRepository repository.StudentRepository
}

func (r StudentHandler) GetAll(c echo.Context) error {

	return c.JSON(http.StatusOK, r.StudentRepository.GetAll())
}

func (r StudentHandler) Create(c echo.Context) error {

	s := &entity.Student{}
	if err := c.Bind(s); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, r.StudentRepository.Create(s))
}

func (r StudentHandler) Update(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	s := new(entity.Student)
	if err := c.Bind(s); err != nil {
		return err
	}
	res, s := r.StudentRepository.Update(id, s)
	if s == nil {
		return c.JSON(http.StatusNotFound, res)
	}
	return c.JSON(http.StatusOK, s)
}

func (r StudentHandler) Patch(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	s := new(entity.Student)
	if err := c.Bind(s); err != nil {
		return err
	}
	res, s := r.StudentRepository.Patch(id, s)
	if s == nil {
		return c.JSON(http.StatusNotFound, res)
	}
	return c.JSON(http.StatusOK, s)
}

func (r StudentHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if r.StudentRepository.Delete(id) == nil {
		return c.JSON(http.StatusAccepted, nil)
	}
	return c.JSON(http.StatusNotFound, "student not found")
}
