package controller

import (
	"net/http"
	"github.com/labstack/echo"
	"../logic"
)

func Test(c echo.Context) error {
	test := logic.Test()
	return c.JSON(http.StatusOK, test)
}