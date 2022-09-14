package controllers

import (
	"day2-task1/entities"
	"day2-task1/lib/database"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetBooksControllers(c echo.Context) error {

	//get data from books
	books, err := database.GetBooks()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  books,
	})
}

func GetBookByIdControllers(c echo.Context) error {

	id := c.Param("id")

	idConvert, _ := strconv.Atoi(id)

	//get data from books
	books, err := database.GetBookById(idConvert)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  books,
	})
}

func CreateBook(c echo.Context) error {
	//get data create book
	u := new(entities.Books)

	//binding
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	//change into DTO
	createBook := entities.ToBooksDTO(u)

	//from DTO to models
	bookModels := entities.AssembBooksDTO(createBook)

	//get data from books
	books, err := database.CreateBook(bookModels)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  books,
	})

}

func UpdateBook(c echo.Context) error {

	id := c.Param("id")
	idConvert, _ := strconv.Atoi(id)
	//get data create user
	u := new(entities.Books)

	//binding
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	//change into DTO
	createBook := entities.ToBooksDTO(u)

	//from DTO to models
	bookModels := entities.AssembBooksDTO(createBook)

	//get data from books
	books, err := database.UpdateBookById(idConvert, bookModels)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  books,
	})

}

func DeleteBookByIdControllers(c echo.Context) error {

	id := c.Param("id")

	idConvert, _ := strconv.Atoi(id)

	//get data from books
	_, err := database.GetBookById(idConvert)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success delete",
	})
}
