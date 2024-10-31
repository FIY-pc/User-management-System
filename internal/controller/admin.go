package controller

import (
	"github.com/FIY-pc/User-management-System/internal/controller/Params"
	"github.com/FIY-pc/User-management-System/internal/model"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// CreateAdmin 创建管理员信息
func CreateAdmin(e echo.Context) error {
	newAdmin := model.Admin{}
	newAdmin.AdminName = e.QueryParam("username")
	// 检查管理员是否已存在
	_, err := model.GetAdminByName(newAdmin.AdminName)
	if err == nil {
		return e.JSON(http.StatusBadRequest, Params.ErrorRec{Code: http.StatusBadRequest, Msg: "CreateAdmin fail", Error: "Admin already exists"})
	}
	// 原始密码
	rawAdminPass := e.QueryParam("password")
	// 对密码进行哈希加密
	hashedAdminPass, err := bcrypt.GenerateFromPassword([]byte(rawAdminPass), bcrypt.DefaultCost)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, Params.ErrorRec{Code: http.StatusBadRequest, Msg: "CreateAdmin fail", Error: err.Error()})
	}
	newAdmin.AdminPass = string(hashedAdminPass)
	// 插入数据库
	if err := model.CreateAdmin(&newAdmin); err != nil {
		return e.JSON(http.StatusInternalServerError, Params.ErrorRec{Code: http.StatusBadRequest, Msg: "CreateAdmin fail", Error: err.Error()})
	}
	// 返回结果
	return e.JSON(http.StatusOK, Params.SuccessRec{Code: http.StatusOK, Msg: "CreateAdmin success", Data: newAdmin})
}

// UpdateAdmin 更新管理员信息
func UpdateAdmin(e echo.Context) error {
	newAdmin, err := model.GetAdminByName(e.QueryParam("username"))
	// 原始新密码
	rawPassword := e.QueryParam("password")
	// 对密码进行哈希加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, Params.ErrorRec{Code: http.StatusBadRequest, Msg: "UpdateAdmin fail", Error: err.Error()})
	}
	newAdmin.AdminPass = string(hashedPassword)
	// 更新数据库
	if err := model.UpdateAdmin(newAdmin); err != nil {
		return e.JSON(http.StatusInternalServerError, Params.ErrorRec{Code: http.StatusBadRequest, Msg: "UpdateAdmin fail", Error: err.Error()})
	}
	// 返回结果
	resultAdmin, err := model.GetAdminByName(newAdmin.AdminName)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, Params.ErrorRec{Code: http.StatusBadRequest, Msg: "UpdateAdmin fail", Error: err.Error()})
	}
	return e.JSON(http.StatusOK, Params.SuccessRec{Code: http.StatusOK, Msg: "UpdateAdmin success", Data: resultAdmin})
}

// GetAdminByName 根据用户名获取管理员信息
func GetAdminByName(e echo.Context) error {
	resultAdmin, err := model.GetAdminByName(e.QueryParam("username"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, Params.ErrorRec{Code: http.StatusBadRequest, Msg: "GetAdminByName fail", Error: err.Error()})
	}
	return e.JSON(http.StatusOK, Params.SuccessRec{Code: http.StatusOK, Data: resultAdmin})
}

// DeleteAdminByName 删除管理员信息
func DeleteAdminByName(e echo.Context) error {
	adminName := e.QueryParam("username")
	if err := model.DeleteAdminByName(adminName); err != nil {
		return e.JSON(http.StatusBadRequest, Params.ErrorRec{Code: http.StatusBadRequest, Msg: "DeleteAdmin fail", Error: err.Error()})
	}
	return e.JSON(http.StatusOK, Params.SuccessRec{Code: http.StatusOK, Msg: "DeleteAdmin success"})
}
