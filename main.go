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

	v1 := e.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		routers.AuthRouter(auth)

		users := v1.Group("/users")
		admin.UserRouter(users)

		places := v1.Group("/places")
		admin.PlaceRouter(places)

		userAddress := v1.Group("/profile")
		user.UserAddressRouter(userAddress)
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.ENV.PORT)))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.ENV.PORT)))
}
