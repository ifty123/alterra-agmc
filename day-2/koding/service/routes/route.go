package routes

import (
	"day2-task1/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	//task 2 : Dynamic
	u := e.Group("/users")
	u.GET("", controllers.GetUserControllers)
	u.POST("", controllers.CreateUserController)
	u.PUT("/:id", controllers.UpdateUserController)
	u.GET("/:id", controllers.GetUserById)
	u.DELETE("/:id", controllers.DeleteUserById)

	//task 1 : static CRUD books
	g := e.Group("/books")
	g.GET("", controllers.GetBooksControllers)
	g.GET("/:id", controllers.GetBookByIdControllers)
	g.POST("", controllers.CreateBook)
	g.PUT("/:id", controllers.UpdateBook)
	g.DELETE("/:id", controllers.DeleteBookByIdControllers)

	return e
}
