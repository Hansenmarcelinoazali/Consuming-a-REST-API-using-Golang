package handler

import (
	"external_api/api/service"
	"net/http"

	"github.com/labstack/echo"
)

func HandlerUsecase6(c echo.Context) error {

	data, err := service.ServiceUsecase6()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}
