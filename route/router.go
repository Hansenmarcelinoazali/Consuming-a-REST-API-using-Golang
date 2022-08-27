package route

import (
	"external_api/api"
	"external_api/api/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", api.Home)
	e.GET("/api/v1/external/dummyjson/products", handler.GetDataUrlorRedis)

	return e
}
