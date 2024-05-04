package main

import (
	"evoting/config"
	"evoting/routers"
	"evoting/routers/admin"
	"evoting/routers/user"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadConfig()
	config.LoadDb()
	e := echo.New()
	e.Static("/images", "public/images")
	auth := e.Group("/")
	routers.AuthRouter(auth)

	users := e.Group("/users")
	admin.UserRouter(users)
	places := e.Group("/places")
	routers.PlaceRouter(places)
	userAddress := e.Group("/profile")
	user.UserAddressRouter(userAddress)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.ENV.PORT)))
}
