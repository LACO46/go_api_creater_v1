package controller

import (
	"net/http"
	
	"github.com/labstack/echo"
)

func Show(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"team": c.QueryParam("team"),"member": c.QueryParam("member")})
}