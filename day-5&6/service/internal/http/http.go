package http

import (
	"day2-task1/internal/app/user"
	"day2-task1/internal/factory"
	"day2-task1/internal/pkg/util"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	e.GET("/status", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "OK"})
	})
	v1 := e.Group("/api/v1")
	user.NewHandler(f).Route(v1.Group("/user"))
	user.NewHandler(f).RouteNoAuth(v1.Group("/login"))
}
