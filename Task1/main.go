package main

import (
	"Task1/handler"
	"Task1/repository"
	"log"

	"github.com/labstack/echo/v4"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	dbConnection, err := gorm.Open(postgres.Open("host=localhost user=postgres password=76435026 dbname=Students port=5432"), &gorm.Config{})
	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}
	log.Println("Hey! You successfully connected to your database.")
	studentHandler := handler.StudentHandler{
		StudentRepository: repository.StudentRepository{
			DB: dbConnection,
		},
	}
	e.GET("/students", studentHandler.GetAll)
	e.POST("/students", studentHandler.Create)
	e.PUT("/students/:id", studentHandler.Update)
	e.PATCH("/students/:id", studentHandler.Patch)
	e.DELETE("/students/:id", studentHandler.Delete)

	e.Logger.Fatal(e.Start(":8080"))
}
