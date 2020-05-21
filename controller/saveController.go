package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func Save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")

	return c.JSON(http.StatusOK, map[string]string{"name":name, "email":email})
}