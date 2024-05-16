package main

import (
	"evoting/config"
	"evoting/middlewares"
	"evoting/routers"
	"evoting/schedulers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.LoadConfig()
	config.LoadDb()
	e := echo.New()
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	middlewares.LogMiddleware(e)

	go schedulers.StartScheduler()
	e.Static("/images", "public/images")
	routers.SetupRouter(e)
	e.Logger.Fatal(e.Start(":1323"))
}
