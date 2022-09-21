package main

import (
	"day2-task1/internal/factory"
	"day2-task1/internal/http"
	mid "day2-task1/internal/middleware"

	"github.com/labstack/echo"
)

func main() {
	f := factory.NewFactory()
	e := echo.New()
	mid.LogMiddleware(e)
	http.NewHttp(e, f)
	e.Logger.Fatal(e.Start(":8000"))
}
