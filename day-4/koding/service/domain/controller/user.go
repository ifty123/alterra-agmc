package controller

import "github.com/labstack/echo"

type UserController interface {
	GetUserControllers(c echo.Context) error
	CreateUserController(c echo.Context) error
	UpdateUserController(c echo.Context) error
	GetUserById(c echo.Context) error
	DeleteUserById(c echo.Context) error
	LoginController(c echo.Context) error
}
