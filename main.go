package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	registerAPI(e)

	e.Logger.Fatal(e.Start(":8000"))

}
