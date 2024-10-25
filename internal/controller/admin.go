package controller

import (
	"User-management-System/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateAdmin(e echo.Context) error {
	newAdmin := model.Admin{}
	newAdmin.AdminName = e.QueryParam("username")
	newAdmin.AdminPass = e.QueryParam("password")
	if err := model.CreateAdmin(&newAdmin); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{"massage": "CreateAdmin fail"})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{"message": "CreateAdmin success", "AdminName": newAdmin.AdminName})
}

func UpdateAdmin(e echo.Context) error {
	newAdmin := model.Admin{}
	newAdmin.AdminName = e.QueryParam("username")
	newAdmin.AdminPass = e.QueryParam("password")
	if err := model.UpdateAdmin(&newAdmin); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{"massage": "UpdateAdmin fail"})
	}
	resultAdmin, err := model.GetAdminByName(newAdmin.AdminName)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{"massage": "UpdateAdmin fail"})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{"message": "UpdateAdmin success", "AdminName": resultAdmin.AdminName, "AdminPass": resultAdmin.AdminPass})
}

func GetAdminByName(e echo.Context) error {
	resultAdmin, err := model.GetAdminByName(e.QueryParam("username"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{"massage": "GetAdmin fail"})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{"message": "GetAdmin success", "AdminName": resultAdmin.AdminName, "AdminPass": resultAdmin.AdminPass, "CreateAt": resultAdmin.CreatedAt})
}

func DeleteAdmin(e echo.Context) error {
	if err := model.DeleteAdminByName(e.QueryParam("username")); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{"massage": "DeleteAdmin fail"})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{"message": "DeleteAdmin success", "AdminName": e.QueryParam("AdminName")})
}
