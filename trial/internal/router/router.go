package router

import (
	"User-management-System/trial/internal/config"
	controller2 "User-management-System/trial/internal/controller"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func InitRouter(e *echo.Echo) {
	InitUserRouter(e)
	InitAdminRouter(e)
	InitStaticRouter(e)
	InitHomeRouter(e)
	InitRedirectRouter(e)
}

func InitUserRouter(e *echo.Echo) {
	e.POST(fmt.Sprintf("%s/Login", config.Config.Server.ApiPrefix), controller2.UserLogin)
	e.POST(fmt.Sprintf("%s/Register", config.Config.Server.ApiPrefix), controller2.UserRegister)
	e.GET(fmt.Sprintf("%s/GetUser", config.Config.Server.ApiPrefix), controller2.UserGet)
	e.GET(fmt.Sprintf("%s/UpdateUser", config.Config.Server.ApiPrefix), controller2.UserUpdate)
}

func InitAdminRouter(e *echo.Echo) {
	e.POST(fmt.Sprintf("%s/AdminLogin", config.Config.Server.ApiPrefix), controller2.AdminLogin)
	e.POST(fmt.Sprintf("%s/AddAdmin", config.Config.Server.ApiPrefix), controller2.AddAdmin)
	e.GET(fmt.Sprintf("%s/DeleteAdmin", config.Config.Server.ApiPrefix), controller2.DeleteAdmin)
	e.GET(fmt.Sprintf("%s/GetAdminByName", config.Config.Server.ApiPrefix), controller2.GetAdminByName)
}

func InitStaticRouter(e *echo.Echo) {
	e.Static("/static", "./trial/web/static")
}

// InitRedirectRouter 测试路由
func InitRedirectRouter(e *echo.Echo) {
	e.GET("/Register", func(c echo.Context) error {
		return c.Redirect(http.StatusPermanentRedirect, "/static/register.html")
	})
}

// InitHomeRouter 测试路由
func InitHomeRouter(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusPermanentRedirect, "/static/home.html")
	})
}
