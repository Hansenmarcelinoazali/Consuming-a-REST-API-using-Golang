package handler

import (
	"errors"
	"external_api/api/service"
	"external_api/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func HandlerSaveToDB(c echo.Context) error {
	refreshToken := c.Request().Header.Get("Id-User")
	inputData := new(model.RequestProduct)
	if err := c.Bind(inputData); err != nil {
		return err
	}
	if inputData.Datas == nil {
		fmt.Println("XXXXX", &inputData.Data)

		data, err := service.ServiceSavetoDb(&inputData.Data, refreshToken)
		if err != nil {
			return errors.New(" ")
		}

		return c.JSON(http.StatusOK, data)

	}

	data, err := service.ServiceInputDatas(inputData, refreshToken)
	fmt.Println("INI DATAS HANDLER", inputData.Datas)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)

}

func GetProductFromDb(c echo.Context) error {

	filter := c.QueryParam("identity")
	limit := c.QueryParam("limit")
	skip := c.QueryParam("skip")

	result, err := service.ServiceGetDatafromDb(filter, limit, skip)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func GetUsers(c echo.Context) error {

	urlparam := "https://dummyjson.com/users"

	result, errorr := service.ServiceGetUsers(urlparam)
	if errorr != nil {
		return errorr

	}
	return c.JSON(http.StatusOK, result)

}
