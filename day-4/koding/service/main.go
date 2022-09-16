package main

import (
	"day2-task1/config"
	"day2-task1/controllers"
	mid "day2-task1/middleware"
	"day2-task1/routes"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	config.InitDB()
	uh := controllers.NewUserController(nil)
	e := routes.New(uh)
	e.Validator = &CustomValidator{validator: validator.New()}
	mid.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":8000"))
}
