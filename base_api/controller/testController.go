package controller

import (
	"net/http"

	"../logic"
	"github.com/labstack/echo"
)

func Test(c echo.Context) error {
	test := logic.Test()
	return c.JSON(http.StatusOK, test)
}
