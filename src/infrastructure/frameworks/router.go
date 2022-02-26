package frameworks

import (
	"cleanArchitectureGolang/src/interfaceadaptors/controllers"
	"cleanArchitectureGolang/src/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	// Echo instance
	e := echo.New()
	e.Use(middleware.CORS())
	interactor := usecases.NewHelloWorldInteractor()
	helloController := controllers.NewHelloController(interactor)

	e.GET("/hello-world", func(c echo.Context) error {
		helloworld := helloController.GetHelloWorld()
		c.Bind(&helloworld)
		return c.JSON(http.StatusOK, helloworld)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
