package main

import (
	"github.com/labstack/echo"
	"./controller"
)

func routing(e *echo.Echo) {
	e.GET("/", controller.Hello)
}