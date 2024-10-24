package router

import (
	"User-management-System/internal/config"
	"User-management-System/internal/controller"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func InitUserRouter(e *echo.Echo) {
	e.POST(fmt.Sprintf("%s/Login", config.Config.Server.ApiPrefix), controller.UserLogin)
	e.POST(fmt.Sprintf("%s/Register", config.Config.Server.ApiPrefix), controller.UserRegister)
}

func InitAdminRouter(e *echo.Echo) {

}

func TestRouter(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})
}
