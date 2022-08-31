package middleware

import (
	"external_api/middleware/servicemiddleware"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func MiddlewareAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenKey := c.Request().Header.Get("X-App-Key")
		tokenSecret := c.Request().Header.Get("X-App-Secret")
		// fmt.Println(tokenKey)

		keyS := "TRAINING"
		secret := "RAYA.OJT"

		if tokenKey != keyS || tokenSecret != secret {
			return c.JSON(http.StatusUnauthorized, "key atau secret salah")
		}
		return next(c)
	}
}

func MiddlewareIslogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		refreshToken := c.Request().Header.Get("Refresh-Token")
		fmt.Println("refresh token middleware", refreshToken)

		status, err := servicemiddleware.ServiceMiddlewareIsLogin(refreshToken)
		if err != nil {
			return err
		}
		if status == "sudah login" {
			return next(c)
		}
		return next(c)
	}
}
