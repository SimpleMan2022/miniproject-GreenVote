package main

import (
	"evoting/config"
	"evoting/middlewares"
	"evoting/routers"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadConfig()
	config.LoadDb()
	e := echo.New()
	middlewares.LogMiddleware(e)
	e.Static("/images", "public/images")
	routers.SetupRouter(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.ENV.PORT)))
}
