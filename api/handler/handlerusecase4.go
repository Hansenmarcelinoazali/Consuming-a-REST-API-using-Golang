package handler

import (
	"errors"
	"external_api/api/service"
	"external_api/model"
	"net/http"

	"github.com/labstack/echo"
)

func HandlerSaveToDB(c echo.Context) error {
	refreshToken := c.Request().Header.Get("Refresh-Token")
	inputData := new(model.RequestProduct)
	if err := c.Bind(inputData); err != nil {
		return err
	}
	if inputData.Datas == nil {
		// marshalinput, err := json.Marshal(inputData.Data)
		// if err != nil {
		// 	return err
		// }

		// unmarInputData := model.Product{}
		// errUnmars := json.Unmarshal(marshalinput, &unmarInputData)
		// if errUnmars != nil {
		// 	return errUnmars
		// }

		data, err := service.ServiceSavetoDb(&inputData.Data, refreshToken)
		if err != nil {
			return errors.New(" ")
		}

		return c.JSON(http.StatusOK, data)

	}

	data, err := service.ServiceInputDatas(inputData, refreshToken)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)

}
