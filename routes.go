package main

import (
	"github.com/labstack/echo"
	"./controller"
)

func routing(e *echo.Echo) {
	e.GET("/", controller.Test)
	e.GET("/users/:id/", controller.User)
	e.GET("/show/", controller.Show)
	e.POST("/save/", controller.Save)
	e.POST("/users/", controller.UserJson)
}