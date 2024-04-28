package main

import (
	"evoting/config"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	config.LoadConfig()
	config.LoadDb()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.ENV.PORT)))
}
