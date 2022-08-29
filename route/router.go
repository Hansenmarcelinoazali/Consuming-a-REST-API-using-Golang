package route

import (
	"external_api/api"
	"external_api/api/handler"

	"github.com/labstack/echo"
	// "github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	// e.Use(middleware.CORS())

	e.GET("/", api.Home)

	//use case 1
	e.POST("/api/v1/auth/login", handler.HandlerUserLogin)

	e.GET("/api/v1/auth/health", handler.HandlerCheckHealth)

	e.DELETE("/api/v1/auth/logout", handler.HandlerLogout)

	//usecase 2
	e.GET("/api/v1/external/dummyjson/products", handler.GetDataUrlorRedis)

	return e
}
