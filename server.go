package main

import (
	"cleanArchitectureGolang/src/infrastructure/frameworks"

	"github.com/labstack/echo/v4"
)

func main() {
	frameworks.Init()
	e := echo.New()
	e.Logger.Fatal(e.Start(":1323"))
}
