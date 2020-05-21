package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

type UserJsonType struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func User(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"id": c.Param("id")})
}

func UserJson(c echo.Context) error {
	u := new(UserJsonType)
	if err := c.Bind(u); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
	// return c.XML(http.StatusCreated, u)
}