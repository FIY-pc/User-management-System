package controller

import (
	"User-management-System/internal/model"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// CreateAdmin 创建管理员信息
func CreateAdmin(e echo.Context) error {
	newAdmin := model.Admin{}
	newAdmin.AdminName = e.QueryParam("username")
	// 原始密码
	rawAdminPass := e.QueryParam("password")
	// 对密码进行哈希加密
	hashedAdminPass, err := bcrypt.GenerateFromPassword([]byte(rawAdminPass), bcrypt.DefaultCost)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{"massage": "CreateAdmin fail"})
	}
	newAdmin.AdminPass = string(hashedAdminPass)
	// 插入数据库
	if err := model.CreateAdmin(&newAdmin); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{"massage": "CreateAdmin fail"})
	}
	// 返回结果
	return e.JSON(http.StatusOK, map[string]interface{}{"message": "CreateAdmin success", "AdminName": newAdmin.AdminName})
}

// UpdateAdmin 更新管理员信息
func UpdateAdmin(e echo.Context) error {
	newAdmin := model.Admin{}
	newAdmin.AdminName = e.QueryParam("username")
	// 原始密码
	rawPassword := e.QueryParam("password")
	// 对密码进行哈希加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{"massage": "UpdateAdmin fail"})
	}
	newAdmin.AdminPass = string(hashedPassword)
	// 更新数据库
	if err := model.UpdateAdmin(&newAdmin); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{"massage": "UpdateAdmin fail"})
	}
	// 返回结果
	resultAdmin, err := model.GetAdminByName(newAdmin.AdminName)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{"massage": "UpdateAdmin fail"})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{"message": "UpdateAdmin success", "AdminName": resultAdmin.AdminName, "AdminPass": resultAdmin.AdminPass})
}

// GetAdminByName 根据用户名获取管理员信息
func GetAdminByName(e echo.Context) error {
	resultAdmin, err := model.GetAdminByName(e.QueryParam("username"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{"massage": "GetAdmin fail"})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{"message": "GetAdmin success", "AdminName": resultAdmin.AdminName, "AdminPass": resultAdmin.AdminPass, "CreateAt": resultAdmin.CreatedAt})
}

// DeleteAdmin 删除管理员信息
func DeleteAdmin(e echo.Context) error {
	if err := model.DeleteAdminByName(e.QueryParam("username")); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{"massage": "DeleteAdmin fail"})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{"message": "DeleteAdmin success", "AdminName": e.QueryParam("AdminName")})
}
