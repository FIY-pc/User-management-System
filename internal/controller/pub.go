package controller

import (
	"github.com/FIY-pc/User-management-System/internal/config"
	"github.com/FIY-pc/User-management-System/internal/controller/Params"
	"github.com/FIY-pc/User-management-System/internal/model"
	utils "github.com/FIY-pc/User-management-System/internal/util"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(e echo.Context) error {
	Name := e.QueryParam("username")
	Password := e.QueryParam("password")
	Email := e.QueryParam("email")
	// 非空检查
	if Name == "" {
		return e.JSON(http.StatusBadRequest, Params.ErrorRec{Code: http.StatusBadRequest, Msg: "username is empty"})
	}
	if Password == "" {
		return e.JSON(http.StatusBadRequest, Params.ErrorRec{Code: http.StatusBadRequest, Msg: "password is empty"})
	}
	if Email == "" {
		return e.JSON(http.StatusBadRequest, Params.ErrorRec{Code: http.StatusBadRequest, Msg: "email is empty"})
	}
	// 检查是否存在同名user
	if existUser, _ := model.GetUserByName(Name); existUser.Name == Name {
		return e.JSON(http.StatusBadRequest, Params.ErrorRec{Code: http.StatusBadRequest, Msg: "username is exist"})
	}
	// 注册
	newUser := model.User{}
	newUser.Name = Name
	// 密码加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(Password), config.Config.Bcrypt.Cost)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, Params.ErrorRec{Code: http.StatusInternalServerError, Msg: "failed to hash password", Error: err.Error()})
	}
	newUser.Password = string(hashPassword)

	if err := model.CreateUser(&newUser); err != nil {
		return e.JSON(http.StatusInternalServerError, Params.ErrorRec{Code: http.StatusInternalServerError, Msg: "failed to create user", Error: err.Error()})
	} else {
		return e.JSON(http.StatusOK, Params.SuccessRec{Code: http.StatusOK, Msg: "register success"})
	}
}

// Login 账密登陆
func Login(e echo.Context) error {
	user := &model.User{}
	user.Name = e.QueryParam("username")
	user.Password = e.QueryParam("password")

	// user验证，先查询user表中是否存在该用户
	resultUser, err := model.GetUserByName(user.Name)
	// 若用户不存在于user表，进行进一步admin检查
	if err != nil {
		// 开启admin验证
		// 查询是否存在该admin
		resultAdmin, err := model.GetAdminByName(user.Name)
		if err != nil {
			// 确定admin表和user表都不存在该用户
			return e.JSON(http.StatusBadRequest, Params.ErrorRec{Code: http.StatusBadRequest, Msg: "failed to get admin", Error: err.Error()})
		}
		// 验证adminPass
		err = bcrypt.CompareHashAndPassword([]byte(resultAdmin.AdminPass), []byte(user.Password))
		if err != nil {
			return e.JSON(http.StatusBadRequest, Params.ErrorRec{Code: http.StatusBadRequest, Msg: "wrong password", Error: err.Error()})
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
			return e.JSON(http.StatusInternalServerError, Params.ErrorRec{Code: http.StatusInternalServerError, Msg: "failed to generate token", Error: err.Error()})
		}
		// 返回token
		return e.JSON(http.StatusOK, Params.LoginRec{Code: http.StatusOK, Msg: "Login success", Username: user.Name, Token: token, IsAdmin: true})
	}

	// user密码验证
	err = bcrypt.CompareHashAndPassword([]byte(resultUser.Password), []byte(user.Password))
	if err != nil {
		return e.JSON(http.StatusBadRequest, Params.ErrorRec{Code: http.StatusBadRequest, Msg: "wrong password", Error: err.Error()})
	}

	// 生成token
	userClaims := utils.JwtClaims{
		UserId: resultUser.ID,
		Role:   "user",
		Exp:    jwt.TimeFunc().Unix() + config.Config.Jwt.Exp,
	}
	token, err := utils.GenerateToken(userClaims)
	if err != nil {
		return e.JSON(http.StatusBadRequest, Params.ErrorRec{Code: http.StatusInternalServerError, Msg: "failed to generate token", Error: err.Error()})
	}
	// 返回token
	return e.JSON(http.StatusOK, Params.LoginRec{Code: http.StatusOK, Msg: "Login success", Username: user.Name, Token: token, IsAdmin: false})
}
