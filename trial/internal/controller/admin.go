package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func AdminLogin(e echo.Context) error {
	return e.JSON(http.StatusOK, "admin login")
}

func AddAdmin(e echo.Context) error {
	return e.JSON(http.StatusOK, "admin add")
}

func DeleteAdmin(e echo.Context) error {
	return e.JSON(http.StatusOK, "admin delete")
}

func GetAdminByName(e echo.Context) error {
	return e.JSON(http.StatusOK, "admin get by name")
}
