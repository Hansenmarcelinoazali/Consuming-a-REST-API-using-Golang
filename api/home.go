package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func Home(c echo.Context) error {

	return c.String(http.StatusOK, "Hello Welcome\n To \n My Page ")

}
