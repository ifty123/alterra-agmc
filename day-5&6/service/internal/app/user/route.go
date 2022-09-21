package user

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func (h *handler) Route(g *echo.Group) {
	// g.Use(middleware.JWTMiddleware(dto.JWTClaims{}, util.JWT_SECRET))
	g.Use(middleware.JWT([]byte(os.Getenv("TOKEN_SECRET"))))
	g.GET("", h.Get)
	g.POST("/create", h.CreateUser)
	g.PUT("/update", h.UpdateUsers)
	g.GET("/user/:id", h.GetUserById)
}

func (h *handler) RouteNoAuth(g *echo.Group) {
	g.POST("", h.Login)
}
