package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func UserRegister(e echo.Context) error {
	head := e.Request().Header.Get("head")

	return e.JSON(http.StatusOK, head)
}

func UserLogin(e echo.Context) error {
	return nil
}
