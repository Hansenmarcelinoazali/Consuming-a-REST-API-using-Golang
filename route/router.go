package route

import (
	"external_api/api"
	"external_api/api/handler"
	"external_api/middleware"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.MiddlewareAuth)

	//home
	e.GET("/", api.Home)

	e.POST("/api/v1/auth/login", handler.HandlerUserLogin) //usecase 1

	e.GET("/api/v1/auth/health", handler.HandlerCheckHealth) //usecase 1

	e.GET("/api/v1/external/dummyjson/products", handler.GetDataUrlorRedis, middleware.MiddlewareIslogin) //usecase 2

	e.POST("/api/v1/products", handler.HandlerSaveToDB, middleware.MiddlewareIslogin) //usecase 4

	e.GET("/api/v1/products", handler.GetProductFromDb, middleware.MiddlewareIslogin) //usecase 4

	e.DELETE("/api/v1/auth/logout", handler.HandlerLogout) //usecase 1

	return e
}
