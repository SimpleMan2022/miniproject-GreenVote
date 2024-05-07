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
	v1User := e.Group("/api/v1")
	{
		userAddress := v1User.Group("/profile")
		user.UserAddressRouter(userAddress)

		places := v1User.Group("/places")
		user.PlaceRouter(places)
	}

	v1Admin := e.Group("/api/v1/admin")
	{
		auth := v1Admin.Group("/auth")
		routers.AuthRouter(auth)

		users := v1Admin.Group("/users")
		admin.UserRouter(users)

		places := v1Admin.Group("/places")
		admin.PlaceRouter(places)

		weather := v1Admin.Group("/places")
		admin.WeatherRouter(weather)

		placeAddress := v1Admin.Group("/places/address")
		admin.PlaceAddressRouter(placeAddress)

	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.ENV.PORT)))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.ENV.PORT)))
}
