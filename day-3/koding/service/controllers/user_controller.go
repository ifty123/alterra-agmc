package controllers

import (
	"day2-task1/entities"
	"day2-task1/lib/database"
	mid "day2-task1/middleware"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetUserControllers(c echo.Context) error {
	users, e := database.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func CreateUserController(c echo.Context) error {
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
	if err := database.SaveUser(userModels); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  userModels,
	})
}

func UpdateUserController(c echo.Context) error {

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
	if err := database.UpdateUsers(userModels, idConvert); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"id":     idConvert,
		"users":  userModels,
	})
}

func GetUserById(c echo.Context) error {

	//get id
	id := c.Param("id")
	idConvert, _ := strconv.Atoi(id)

	//get user
	getUser, err := database.GetUserById(idConvert)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  getUser,
	})
}

func DeleteUserById(c echo.Context) error {

	//check userId from token
	userToken := mid.GetUserIdAndAllowedMethod(c)

	//get id
	idConvert, _ := strconv.Atoi(c.Param("id"))

	if userToken != idConvert {
		return c.String(http.StatusBadRequest, "invalid user")
	}

	//get user
	err := database.DeleteUser(idConvert)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}

func LoginController(c echo.Context) error {
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

	login, err := database.LoginUser(userModels)
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
