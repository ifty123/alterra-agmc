package routes

import (
	"day2-task1/controllers"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	//task 2 : Dynamic
	u := e.Group("/users")
	u.Use(middleware.JWT([]byte(os.Getenv("TOKEN_SECRET"))))
	u.GET("", controllers.GetUserControllers)

	u.PUT("/:id", controllers.UpdateUserController)
	u.GET("/:id", controllers.GetUserById)
	u.DELETE("/:id", controllers.DeleteUserById)

	//login not using auth
	e.POST("/public/users", controllers.CreateUserController)
	e.POST("/login", controllers.LoginController)

	//no auth
	e.GET("/public/books", controllers.GetBooksControllers)
	e.GET("/public/books/:id", controllers.GetBookByIdControllers)

	//with auth
	g := e.Group("/books")
	g.Use(middleware.JWT([]byte(os.Getenv("TOKEN_SECRET"))))
	g.POST("", controllers.CreateBook)
	g.PUT("/:id", controllers.UpdateBook)
	g.DELETE("/:id", controllers.DeleteBookByIdControllers)

	return e
}
