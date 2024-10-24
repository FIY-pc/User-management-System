package main

import (
	"User-management-System/internal/config"
	"User-management-System/internal/model"
	"User-management-System/internal/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	config.InitConfig()
	model.InitPostgres()
	router.TestRouter(e)
	router.InitAdminRouter(e)
	router.InitUserRouter(e)

	err := e.Start(":" + config.Config.Server.Port)
	if err != nil {
		e.Logger.Fatal(err)
	}
}
