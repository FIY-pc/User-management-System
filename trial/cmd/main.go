package main

import (
	"User-management-System/trial/internal/config"
	"User-management-System/trial/internal/model"
	"User-management-System/trial/internal/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	config.InitConfig()
	model.InitPostgres()
	router.InitRouter(e)
	err := e.Start(":" + config.Config.Server.Port)
	if err != nil {
		e.Logger.Fatal(err)
	}
}
