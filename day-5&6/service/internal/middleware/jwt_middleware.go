package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func CreateToken(id string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claims["iat"] = 5184000

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

func ExtractToken(e echo.Context) int {

	var generate *jwt.Token
	if strings.Contains(e.Request().URL.Path, "users") {
		generate = e.Get("users").(*jwt.Token)
	} else {
		generate = e.Get("books").(*jwt.Token)
	}

	if generate.Valid {
		claim := generate.Claims.(jwt.MapClaims)
		idGenerate := claim["id"].(int)
		return idGenerate
	}
	return 0
}

func GetUserIdAndAllowedMethod(e echo.Context) int {
	if e.Request().Method == "PUT" || e.Request().Method == "DELETE" {
		return ExtractToken(e)
	} else {
		return 0
	}
}
