package middleware

import (
	"github.com/labstack/echo"
)

func Middleware(c echo.Context, next echo.HandlerFunc) (bool, error) {

	tokenKey := c.Request().Header.Get("X-App-Key")
	tokenSecret := c.Request().Header.Get("X-App-Secret")

	keyS := "TRAINING"
	secret := "RAYA.OJT"

	if tokenKey != keyS || tokenSecret != secret {
		return false, nil
	}
	return true, nil
}
