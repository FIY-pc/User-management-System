package router

import (
	"User-management-System/internal/config"
	"User-management-System/internal/controller"
	"User-management-System/internal/util"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

// InitRouter 初始化所有路由
func InitRouter(e *echo.Echo) {
	InitHomeRouter(e)
	InitUserRouter(e)
	InitAdminRouter(e)
}

// InitHomeRouter 初始化首页路由
func InitHomeRouter(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Please use api to access")
	})
}

// InitUserRouter 初始化用户路由
func InitUserRouter(e *echo.Echo) {
	// Login 包含user和admin的login
	e.GET(fmt.Sprintf("%s/Login", config.Config.Server.ApiPrefix), controller.Login)
	// Register 只能注册user
	e.GET(fmt.Sprintf("%s/Register", config.Config.Server.ApiPrefix), controller.Register)
	// user令牌具有更改自身信息的权限
	UserGroup := e.Group(fmt.Sprintf("%s/user", config.Config.Server.ApiPrefix))
	// 验证token
	UserGroup.Use(utils.JWTAuthMiddleware())
	UserGroup.Use(utils.UserRoleMiddleware)
	// 更改自身信息
	UserGroup.GET("/UserUpdateSelf", controller.UserUpdateSelf)
}

// InitAdminRouter 初始化管理员路由
func InitAdminRouter(e *echo.Echo) {
	// admin令牌具有所有权限
	AdminGroup := e.Group(fmt.Sprintf("%s/admin", config.Config.Server.ApiPrefix))
	// 验证token
	AdminGroup.Use(utils.JWTAuthMiddleware())
	AdminGroup.Use(utils.AdminRoleMiddleware)
	// 管理员管理路由CURD
	AdminGroup.GET("/CreateAdmin", controller.CreateAdmin)
	AdminGroup.GET("/UpdateAdmin", controller.UpdateAdmin)
	AdminGroup.GET("/DeleteAdmin", controller.DeleteAdmin)
	AdminGroup.GET("/GetAdminByName", controller.GetAdminByName)
	// 用户管理路由
	AdminGroup.GET("/UserCreate", controller.UserCreate)
	AdminGroup.GET("/UserGet", controller.UserGet)
	AdminGroup.GET("/UserUpdate", controller.UserUpdate)
	AdminGroup.GET("/UserDelete", controller.UserDelete)
}
