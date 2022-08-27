package handler

import (
	"external_api/api/service"
	"net/http"

	"github.com/labstack/echo"
	// "github.com/labstack/echo/v4"
)

func GetDataUrlorRedis(c echo.Context) error {
	// url := "https://dummyjson.com/products?limit=100&skip=0" ////example
	urlparam := "https://dummyjson.com/products?limit=100&skip=0"

	//param
	limit := c.QueryParam("limit")
	skip := c.QueryParam("skip")

	// var err error
	tokenKey := c.Request().Header.Get("X-App-Key")
	tokenSecret := c.Request().Header.Get("X-App-Secret")

	keyS := "TRAINING"
	secret := "RAYA.OJT"

	if tokenKey != keyS {
		return c.JSON(http.StatusUnauthorized, " key salah ")
	}
	if tokenSecret != secret {
		return c.JSON(http.StatusUnauthorized, " secret salah")
	}

	result, errorr := service.GetRedisData(urlparam, limit, skip)
	if errorr != nil {
		// return errorr
		results, err := service.ServiceGetUrl(urlparam)
		if err != nil {
			return err
		}
		return c.JSONPretty(http.StatusOK, results, " ")
	}
	return c.JSONPretty(http.StatusOK, result, "  ")
}
