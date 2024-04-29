package main

import (
	"evoting/config"
	"evoting/routers"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadConfig()
	config.LoadDb()
	e := echo.New()
	auth := e.Group("/")
	routers.AuthRouter(auth)

	places := e.Group("/places")
	routers.PlaceRouter(places)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.ENV.PORT)))
}
