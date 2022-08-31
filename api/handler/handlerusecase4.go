package handler

import (
	"encoding/json"
	"errors"
	"external_api/api/service"
	"external_api/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func HandlerSaveToDB(c echo.Context) error {
	refreshToken := c.Request().Header.Get("Refresh-Token")
	inputData := new(model.RequestProduct)
	if err := c.Bind(inputData); err != nil {
		return err
	}
	if inputData.Type == "1" {
		marshalinput, err := json.Marshal(inputData.Data)
		if err != nil {
			return err
		}

		unmarInputData := model.Product{}
		errUnmars := json.Unmarshal(marshalinput, &unmarInputData)
		if errUnmars != nil {
			return errUnmars
		}

		data, err := service.ServiceSavetoDb(&unmarInputData, refreshToken)
		if err != nil {
			return errors.New(" ")
		}

		return c.JSON(http.StatusOK, data)

	}
	if inputData.Type == "2" {
		marshalinputS, err := json.Marshal(inputData.Data)
		if err != nil {
			return err
		}
		unmarshalInputDatas := []model.Product{}
		errUnmarsslice:= json.Unmarshal(marshalinputS, &unmarshalInputDatas)
		fmt.Println(unmarshalInputDatas)
		if errUnmarsslice != nil{
			fmt.Println("unmars2",errUnmarsslice)
			return errUnmarsslice
		}
		data,err := service.ServiceInputDatas(&unmarshalInputDatas, refreshToken)
		if err != nil{
			return  err
		}


		return c.JSON(http.StatusOK, data)
	}

	// data, err := service.ServiceSavetoDb(inputData, refreshToken)
	// if err != nil {
	// 	return c.JSON(http.StatusNotFound, "data tidak ditemukan")
	// }

	return nil
}
