package main

import (
	"./controller"
	"github.com/labstack/echo"
)

func routing(e *echo.Echo) {
	e.GET("/", controller.Test)
}
