package handler

import (
	"encoding/json"
	"external_api/api/service"
	"external_api/model"
	"fmt"
	"net/http"
	"strings"

	// "github.com/Lionparcel/common-backend/sql/db"
	"github.com/labstack/echo"
	// "github.com/labstack/echo/v4"
)

//usecase 1
func HandlerUserLogin(c echo.Context) error {

	// email := c.QueryParam("email")
	// password := c.QueryParam("password")
	// username := c.QueryParam("username")

	u := new(model.Login)
	if err := c.Bind(u); err != nil {
		return err
	}

	login := model.Login{
		EmailorUsername: u.EmailorUsername,
		Password:        u.Password,
	}

	isLogin, IsError := service.ServiceIsLogin(login.EmailorUsername, login.EmailorUsername, login.Password)
	if IsError != nil {
		return IsError
	} else if isLogin == "Sudah Login" {
		return c.JSON(http.StatusAccepted, isLogin)
	}

	identity, err := service.ServiceLoginUser(login.EmailorUsername, login.EmailorUsername, login.Password)
	if err != nil {
		return err
	}
	var responseToken model.ResponseToken

	tokenAccess, err := service.GenerateAccessToken(login.EmailorUsername)
	if err != nil {
		return nil
	}

	TokenRefresh, err := service.GenerateRefreshToken(login.EmailorUsername)
	if err != nil {
		return err
	}

	responseToken.RefreshToken = TokenRefresh
	responseToken.TokenAccess = tokenAccess

	err = service.TurnToTrue(login.EmailorUsername, login.EmailorUsername, login.Password, TokenRefresh)
	if err != nil {
		return err
	}
	//validasi is login
	// isLogin, IsError := service.ServiceIsLogin(email, username, password)
	// if IsError != nil {
	// 	return IsError
	// } else if isLogin == "Sudah Login" {
	// 	return c.JSON(http.StatusAccepted, isLogin)
	// }

	// identity, err := service.ServiceLoginUser(email, username, password)
	// if err != nil {
	// 	return err
	// }
	// var responseToken model.ResponseToken

	// tokenAccess, err := service.GenerateAccessToken(email)
	// if err != nil {
	// 	return nil
	// }

	// TokenRefresh, err := service.GenerateRefreshToken(email)
	// if err != nil {
	// 	return err
	// }

	// responseToken.RefreshToken = TokenRefresh
	// responseToken.TokenAccess = tokenAccess

	// err = service.TurnToTrue(email, username, password, TokenRefresh)
	// if err != nil {
	// 	return err
	// }
	marshaljson, err := json.Marshal(identity)
	if err != nil {
		return err
	}
	errs := service.ServiceSetRedis(tokenAccess, string(marshaljson))
	if errs != nil {
		return errs
	}

	return c.JSON(http.StatusOK, responseToken)

}

func HandlerCheckHealth(c echo.Context) error {
	accessToken := c.Request().Header.Get("Authorization")

	tokenParamsAuth := strings.Split(accessToken, " ")
	responseHealth, err := service.ServiceCheckHealth(tokenParamsAuth[1])
	fmt.Println("ini handler", responseHealth)

	if err != nil {
		return nil
	}

	return c.JSON(http.StatusOK, responseHealth)
}

func HandlerLogout(c echo.Context) error {
	refreshToken := c.Request().Header.Get("Refresh-Token")
	tokenAccess := c.Request().Header.Get("Authorization")
	tokenAccessaramAuthLogout := strings.Split(tokenAccess, " ")

	err := service.ServiceLogout(tokenAccessaramAuthLogout[1], refreshToken)

	return err
}

//usecase 2
func GetDataUrlorRedis(c echo.Context) error {
	// url := "https://dummyjson.com/products?limit=100&skip=0" ////example
	urlparam := "https://dummyjson.com/products?limit=100&skip=0"

	//param
	limit := c.QueryParam("limit")
	skip := c.QueryParam("skip")

	// var err error

	result, errorr := service.GetRedisData(urlparam, limit, skip)
	if errorr != nil {
		return errorr

	}
	return c.JSON(http.StatusUnauthorized, result)
	// return c.JSONPretty(http.StatusOK, result, "  ")

}
