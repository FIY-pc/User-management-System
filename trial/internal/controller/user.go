package controller

import (
	"User-management-System/trial/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UserRegister(e echo.Context) error {
	user := &model.User{}
	user.Name = e.FormValue("username")
	user.Password = e.FormValue("password")
	if user.Name == "" || user.Password == "" {
		return e.JSON(http.StatusOK, map[string]string{"message": "username or password is empty"})
	}
	if existuser, _ := model.GetUserByName(user.Name); existuser != nil {
		return e.JSON(http.StatusOK, map[string]string{"message": "username already exist"})
	}

	if err := model.CreateUser(user); err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	} else {
		return e.JSON(http.StatusOK, user)
	}
}

func UserLogin(e echo.Context) error {
	user := &model.User{}
	user.Name = e.FormValue("username")
	user.Password = e.FormValue("password")

	resultuser, err := model.GetUserByName(user.Name)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	if resultuser == nil {
		return e.String(http.StatusOK, "user not exist")
	}
	if resultuser.Password != user.Password {
		return e.JSON(http.StatusOK, map[string]string{"message": "password error"})
	}
	return e.JSON(http.StatusOK, map[string]string{"message": "success", "username": resultuser.Name})
}

func UserGet(e echo.Context) error {
	name := e.FormValue("username")
	resultuser, err := model.GetUserByName(name)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, resultuser)
}

func UserUpdate(e echo.Context) error {
	user := model.User{}
	user.Name = e.FormValue("name")
	user.Password = e.FormValue("password")

	err := model.UpdateUser(&user)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	resultuser, err := model.GetUserByName(user.Name)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, resultuser)
}
