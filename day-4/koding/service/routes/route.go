package routes

import (
	"day2-task1/controllers"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(h controllers.UserUseCase) *echo.Echo {
	e := echo.New()

	//task 2 : Dynamic
	u := e.Group("/users")
	u.Use(middleware.JWT([]byte(os.Getenv("TOKEN_SECRET"))))
	u.GET("", h.GetUserControllers)

	u.PUT("/:id", h.UpdateUserController)
	u.GET("/:id", h.GetUserById)
	u.DELETE("/:id", h.DeleteUserById)

	//login not using auth
	e.POST("/public/users", h.CreateUserController)
	e.POST("/login", h.LoginController)

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
