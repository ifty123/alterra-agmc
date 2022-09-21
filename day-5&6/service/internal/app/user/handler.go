package user

import (
	"day2-task1/internal/dto"
	"day2-task1/internal/factory"
	mid "day2-task1/internal/middleware"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type handler struct {
	service UserUseCase
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewUserController(f),
	}
}

func (h *handler) Get(c echo.Context) error {

	//ke getUserController
	result := h.service.GetUserControllers(c.Request().Context())

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  result,
	})
}

func (h *handler) CreateUser(c echo.Context) error {
	//get data create user
	u := new(dto.User)

	//binding
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	//validate
	if errValidate := c.Validate(u); errValidate != nil {
		return c.String(http.StatusBadRequest, errValidate.Error())
	}

	err := h.service.CreateUserController(c.Request().Context(), u)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  u,
	})
}

func (h *handler) Login(c echo.Context) error {
	//get data create user
	u := new(dto.User)

	//binding
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	token, errToken := h.service.LoginController(c.Request().Context(), u)
	if errToken != nil {
		return c.String(http.StatusBadRequest, errToken.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   u.Email,
		"token":  token,
	})
}

func (h *handler) UpdateUsers(c echo.Context) error {
	//check userId from token
	userToken := mid.GetUserIdAndAllowedMethod(c)

	//get id
	idConvert, _ := strconv.Atoi(c.Param("id"))

	if userToken != idConvert {
		return c.String(http.StatusBadRequest, "invalid user")
	}

	//get data create user
	u := new(dto.User)

	//binding
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	//validate
	if errValidate := c.Validate(u); errValidate != nil {
		return c.String(http.StatusBadRequest, errValidate.Error())
	}

	err := h.service.UpdateUserController(c.Request().Context(), u, idConvert)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"id":     idConvert,
		"users":  u,
	})
}

func (h *handler) GetUserById(c echo.Context) error {
	//get id
	id := c.Param("id")

	getUser, err := h.service.GetUserById(c.Request().Context(), id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  getUser,
	})
}

func (h *handler) DeleteUserById(c echo.Context) error {

	//check userId from token
	userToken := mid.GetUserIdAndAllowedMethod(c)

	//get id
	idConvert, _ := strconv.Atoi(c.Param("id"))

	if userToken != idConvert {
		return c.String(http.StatusBadRequest, "invalid user")
	}

	if err := h.service.DeleteUserById(c.Request().Context(), idConvert); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}
