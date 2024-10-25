package controller

import (
	"User-management-System/internal/config"
	"User-management-System/internal/model"
	"User-management-System/internal/util"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Register(e echo.Context) error {
	user := &model.User{}
	user.Name = e.QueryParam("username")
	user.Password = e.QueryParam("password")
	// 非空检查
	if user.Name == "" {
		return e.JSON(http.StatusOK, map[string]string{"message": "username is empty"})
	}
	if user.Password == "" {
		return e.JSON(http.StatusOK, map[string]string{"message": "password is empty"})
	}
	// 检查是否存在同名user
	if existuser, _ := model.GetUserByName(user.Name); existuser.Name == user.Name {
		return e.JSON(http.StatusOK, map[string]string{"message": "username already exist"})
	}
	// 注册
	if err := model.CreateUser(user); err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	} else {
		return e.JSON(http.StatusOK, user)
	}
}

func Login(e echo.Context) error {
	user := &model.User{}
	user.Name = e.QueryParam("username")
	user.Password = e.QueryParam("password")

	// user验证
	resultuser, err := model.GetUserByName(user.Name)
	// 若用户不存在于user表，进行进一步admin检查
	if err != nil {
		// 开启admin验证
		// 查询是否存在该admin
		resultAdmin, err := model.GetAdminByName(user.Name)
		if err != nil {
			// 确定admin表和user表都不存在该用户
			return e.String(http.StatusInternalServerError, err.Error())
		}
		// 验证adminPass
		if user.Password != resultAdmin.AdminPass {
			return e.String(http.StatusInternalServerError, "adminPass Invalid")
		}
		// 验证成功
		// 生成token
		adminClaims := utils.JwtClaims{
			UserId: resultAdmin.ID,
			Role:   "admin",
			Exp:    jwt.TimeFunc().Unix() + config.Config.Jwt.Exp,
		}
		token, err := utils.GenerateToken(adminClaims)
		if err != nil {
			return e.String(http.StatusInternalServerError, err.Error())
		}
		// 返回token
		return e.JSON(http.StatusOK, map[string]string{"message": "Admin Login success", "AdminName": resultAdmin.AdminName, "token": token})
	}

	// user密码验证
	if resultuser.Password != user.Password {
		return e.JSON(http.StatusOK, map[string]string{"message": "password error"})
	}

	// 生成token
	userClaims := utils.JwtClaims{
		UserId: resultuser.ID,
		Role:   "user",
		Exp:    jwt.TimeFunc().Unix() + config.Config.Jwt.Exp,
	}
	token, err := utils.GenerateToken(userClaims)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	// 返回token
	return e.JSON(http.StatusOK, map[string]string{"message": "Login success", "username": resultuser.Name, "token": token})
}

func UserCreate(e echo.Context) error {
	user := &model.User{}
	user.Name = e.QueryParam("username")
	user.Password = e.QueryParam("password")

	// 检查是否存在同名admin
	if _, err := model.GetAdminByName(user.Name); err == nil {
		return e.String(http.StatusInternalServerError, "username exist") // 防止泄漏真实admin账号名
	}

	if err := model.CreateUser(user); err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]string{"message": "Create user success", "username": user.Name})
}

func UserGet(e echo.Context) error {
	name := e.QueryParam("username")
	resultuser, err := model.GetUserByName(name)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, resultuser)
}

func UserUpdate(e echo.Context) error {
	user := model.User{}
	user.Name = e.QueryParam("username")
	user.Password = e.QueryParam("password")

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

func UserUpdateSelf(e echo.Context) error {
	return e.JSON(http.StatusOK, map[string]string{"message": "update self"})
}

func UserDelete(e echo.Context) error {
	resultuser, err := model.GetUserByName(e.QueryParam("username"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	if resultuser == nil {
		return e.JSON(http.StatusOK, map[string]string{"message": "user not exist"})
	}
	if err := model.DeleteUserByName(resultuser.Name); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return e.JSON(http.StatusOK, resultuser)
}
