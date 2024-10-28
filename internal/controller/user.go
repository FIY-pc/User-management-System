package controller

import (
	"User-management-System/internal/config"
	"User-management-System/internal/model"
	"User-management-System/internal/util"
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
		return e.JSON(http.StatusBadRequest, map[string]string{"msg": "username is empty"})
	}
	if Password == "" {
		return e.JSON(http.StatusBadRequest, map[string]string{"msg": "password is empty"})
	}
	if Email == "" {
		return e.JSON(http.StatusBadRequest, map[string]string{"msg": "email is empty"})
	}
	// 检查是否存在同名user
	if existUser, _ := model.GetUserByName(Name); existUser.Name == Name {
		return e.JSON(http.StatusBadRequest, map[string]string{"msg": "username already exist"})
	}
	// 注册
	newUser := model.User{}
	newUser.Name = Name
	// 密码加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(Password), config.Config.Bcrypt.Cost)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"msg": "password encrypt error", "error": err.Error()})
	}
	newUser.Password = string(hashPassword)

	if err := model.CreateUser(&newUser); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"msg": "user create failed", "error": err.Error()})
	} else {
		return e.JSON(http.StatusOK, map[string]string{"msg": "User register success", "username": newUser.Name})
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
			return e.JSON(http.StatusBadRequest, map[string]string{"msg": "user not existed", "error": err.Error()})
		}
		// 验证adminPass
		err = bcrypt.CompareHashAndPassword([]byte(resultAdmin.AdminPass), []byte(user.Password))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{"msg": "Password Invalid", "error": err.Error()})
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
			return e.JSON(http.StatusInternalServerError, map[string]string{"msg": "token generate failed", "error": err.Error()})
		}
		// 返回token
		return e.JSON(http.StatusOK, map[string]string{"msg": "Admin Login success", "AdminName": resultAdmin.AdminName, "token": token})
	}

	// user密码验证
	err = bcrypt.CompareHashAndPassword([]byte(resultUser.Password), []byte(user.Password))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"msg": "Password Invalid", "error": err.Error()})
	}

	// 生成token
	userClaims := utils.JwtClaims{
		UserId: resultUser.ID,
		Role:   "user",
		Exp:    jwt.TimeFunc().Unix() + config.Config.Jwt.Exp,
	}
	token, err := utils.GenerateToken(userClaims)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"msg": "token generate error", "error": err.Error()})
	}
	// 返回token
	return e.JSON(http.StatusOK, map[string]string{"msg": "Login success", "username": resultUser.Name, "token": token})
}

func UserCreate(e echo.Context) error {
	user := &model.User{}
	user.Name = e.QueryParam("username")
	Password := e.QueryParam("password")
	user.Email = e.QueryParam("email")

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(Password), config.Config.Bcrypt.Cost)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"msg": "password encrypt error", "error": err.Error()})
	}
	user.Password = string(hashPassword)
	// 检查是否存在同名admin
	if _, err := model.GetAdminByName(user.Name); err == nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"msg": "username exist"}) // 防止泄漏真实admin账号名
	}

	if err := model.CreateUser(user); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"msg": "User create failed", "error": err.Error()})
	}
	return e.JSON(http.StatusOK, map[string]string{"msg": "Create user success", "username": user.Name})
}

func UserGetByName(e echo.Context) error {
	name := e.QueryParam("username")
	resultUser, err := model.GetUserByName(name)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"msg": "Get user failed", "error": err.Error()})
	}
	return e.JSON(http.StatusOK, resultUser)
}

func UserUpdate(e echo.Context) error {
	user, err := model.GetUserByName(e.QueryParam("username"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"msg": "user not existed", "error": err.Error()})
	}

	if e.QueryParam("password") != "" {
		Password := e.QueryParam("password")
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(Password), config.Config.Bcrypt.Cost)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{"msg": "password encrypt error", "error": err.Error()})
		}
		user.Password = string(hashPassword)
	}
	if e.QueryParam("email") != "" {
		user.Email = e.QueryParam("email")
	}

	err = model.UpdateUser(user)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"msg": "user update failed", "error": err.Error()})
	}

	resultUser, err := model.GetUserByName(user.Name)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, resultUser)
}

func UserUpdateSelf(e echo.Context) error {
	return e.JSON(http.StatusOK, map[string]string{"msg": "update self"})
}

func UserDeleteByName(e echo.Context) error {
	resultUser, err := model.GetUserByName(e.QueryParam("username"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"msg": err.Error()})
	}
	if resultUser == nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"msg": "user not exist"})
	}
	if err := model.DeleteUserByName(resultUser.Name); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}
	return e.JSON(http.StatusOK, map[string]string{"msg": "Delete user success", "username": resultUser.Name})
}
