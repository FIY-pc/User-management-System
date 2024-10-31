package controller

import (
	"github.com/FIY-pc/User-management-System/internal/config"
	"github.com/FIY-pc/User-management-System/internal/controller/Params"
	"github.com/FIY-pc/User-management-System/internal/model"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func UserCreate(e echo.Context) error {
	user := &model.User{}
	user.Name = e.QueryParam("username")
	Password := e.QueryParam("password")
	user.Email = e.QueryParam("email")

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(Password), config.Config.Bcrypt.Cost)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, Params.ErrorRec{Code: http.StatusInternalServerError, Msg: "password encrypt error", Error: err.Error()})
	}
	user.Password = string(hashPassword)
	// 检查是否存在同名admin
	if _, err := model.GetAdminByName(user.Name); err == nil {
		return e.JSON(http.StatusBadRequest, Params.ErrorRec{
			Code:  http.StatusBadRequest,
			Msg:   "user already exists",
			Error: "",
		}) // 防止泄漏真实admin账号名
	}

	if err := model.CreateUser(user); err != nil {
		return e.JSON(http.StatusInternalServerError, Params.ErrorRec{Code: http.StatusInternalServerError, Msg: "create user error", Error: err.Error()})
	}
	return e.JSON(http.StatusOK, Params.SuccessRec{Code: http.StatusOK, Data: user})
}

func UserGetByName(e echo.Context) error {
	name := e.QueryParam("username")
	resultUser, err := model.GetUserByName(name)
	if err != nil {
		return e.JSON(http.StatusBadRequest, Params.ErrorRec{Code: http.StatusBadRequest, Msg: "get user error", Error: err.Error()})
	}
	return e.JSON(http.StatusOK, Params.SuccessRec{Code: http.StatusOK, Data: resultUser})
}

func UserUpdate(e echo.Context) error {
	user, err := model.GetUserByName(e.QueryParam("username"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, Params.ErrorRec{Code: http.StatusBadRequest, Msg: "get user error", Error: err.Error()})
	}

	if e.QueryParam("password") != "" {
		Password := e.QueryParam("password")
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(Password), config.Config.Bcrypt.Cost)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, Params.ErrorRec{Code: http.StatusInternalServerError, Msg: "hash password error", Error: err.Error()})
		}
		user.Password = string(hashPassword)
	}
	if e.QueryParam("email") != "" {
		user.Email = e.QueryParam("email")
	}

	err = model.UpdateUser(user)
	if err != nil {
		return e.JSON(http.StatusBadRequest, Params.ErrorRec{Code: http.StatusInternalServerError, Msg: "update user error", Error: err.Error()})
	}

	resultUser, err := model.GetUserByName(user.Name)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, Params.SuccessRec{Code: http.StatusOK, Data: resultUser})
}

func UserDeleteByName(e echo.Context) error {
	resultUser, err := model.GetUserByName(e.QueryParam("username"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, Params.ErrorRec{Code: http.StatusBadRequest, Msg: "get user error", Error: err.Error()})
	}
	if resultUser == nil {
		return e.JSON(http.StatusBadRequest, Params.ErrorRec{Code: http.StatusBadRequest, Msg: "user not found", Error: ""})
	}
	if err := model.DeleteUserByName(resultUser.Name); err != nil {
		return e.JSON(http.StatusInternalServerError, Params.ErrorRec{Code: http.StatusInternalServerError, Msg: "delete user error", Error: err.Error()})
	}
	return e.JSON(http.StatusOK, Params.SuccessRec{Code: http.StatusOK, Data: resultUser})
}

func UserUpdateSelf(e echo.Context) error {
	return e.JSON(http.StatusOK, Params.SuccessRec{Code: http.StatusOK, Data: e.Param("username")})
}
