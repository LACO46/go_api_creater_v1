package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func User(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"id": c.Param("id")})
}