package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Debug = true
	routing(e)
	e.Start(":80")
}