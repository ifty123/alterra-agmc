package controllers

import (
	"day2-task1/entities"
	mid "day2-task1/middleware"
	"net/http"
	"strconv"

	"day2-task1/domain/repositories"

	"github.com/labstack/echo"
)

type UserUseCase interface {
	GetUserControllers(c echo.Context) error
	CreateUserController(c echo.Context) error
	UpdateUserController(c echo.Context) error
	GetUserById(c echo.Context) error
	DeleteUserById(c echo.Context) error
	LoginController(c echo.Context) error
}

type UserController struct {
	UserRepo repositories.UserDatabase
}

func NewUserController(userRepo repositories.UserDatabase) *UserController {
	return &UserController{
		UserRepo: userRepo,
	}
}

func (uc *UserController) GetUserControllers(c echo.Context) error {
	users, e := uc.UserRepo.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func (uc *UserController) CreateUserController(c echo.Context) error {
	//get data create user
	u := new(entities.User)

	//binding
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	//validate
	if errValidate := c.Validate(u); errValidate != nil {
		return c.String(http.StatusBadRequest, errValidate.Error())
	}

	//change into DTO
	userDTO := entities.ToUserDTO(u)

	//change to models
	userModels := entities.AssembUserDTO(userDTO)

	//save user
	if err := uc.UserRepo.SaveUser(userModels); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  userModels,
	})
}

func (uc *UserController) UpdateUserController(c echo.Context) error {

	//check userId from token
	userToken := mid.GetUserIdAndAllowedMethod(c)

	//get id
	idConvert, _ := strconv.Atoi(c.Param("id"))

	if userToken != idConvert {
		return c.String(http.StatusBadRequest, "invalid user")
	}

	//get data create user
	u := new(entities.User)

	//binding
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	//validate
	if errValidate := c.Validate(u); errValidate != nil {
		return c.String(http.StatusBadRequest, errValidate.Error())
	}

	//change into DTO
	userDTO := entities.ToUserDTO(u)

	//change to models
	userModels := entities.AssembUserDTO(userDTO)

	//save user
	if err := uc.UserRepo.UpdateUsers(userModels, idConvert); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"id":     idConvert,
		"users":  userModels,
	})
}

func (uc *UserController) GetUserById(c echo.Context) error {

	//get id
	id := c.Param("id")

	//get user
	getUser, err := uc.UserRepo.GetUserById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  getUser,
	})
}

func (uc *UserController) DeleteUserById(c echo.Context) error {

	//check userId from token
	userToken := mid.GetUserIdAndAllowedMethod(c)

	//get id
	idConvert, _ := strconv.Atoi(c.Param("id"))

	if userToken != idConvert {
		return c.String(http.StatusBadRequest, "invalid user")
	}

	//get user
	err := uc.UserRepo.DeleteUser(idConvert)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}

func (uc *UserController) LoginController(c echo.Context) error {
	//get data create user
	u := new(entities.User)

	//binding
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	//change into DTO
	userDTO := entities.ToUserDTO(u)

	//change to models
	userModels := entities.AssembUserDTO(userDTO)

	login, err := uc.UserRepo.LoginUser(userModels)
	if err != nil || login == nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	//create token
	token, errToken := mid.CreateToken(strconv.Itoa(userDTO.Id))
	if errToken != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   login,
		"token":  token,
	})
}
