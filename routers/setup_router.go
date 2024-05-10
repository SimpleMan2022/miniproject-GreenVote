package routers

import (
	"evoting/routers/admin"
	"evoting/routers/user"
	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo) {
	v1User := e.Group("/api/v1")
	{
		auth := v1User.Group("/auth")
		AuthUserRouter(auth)

		userAddress := v1User.Group("/profile")
		user.UserAddressRouter(userAddress)

		places := v1User.Group("/places")
		user.PlaceRouter(places)

		votes := v1User.Group("/votes")
		user.VoteRouter(votes)

		comments := v1User.Group("/places")
		user.CommentRouter(comments)
	}

	v1Admin := e.Group("/api/v1/admin")
	{
		auth := v1Admin.Group("/auth")
		AuthAdminRouter(auth)

		users := v1Admin.Group("/users")
		admin.UserRouter(users)

		places := v1Admin.Group("/places")
		admin.PlaceRouter(places)

		weather := v1Admin.Group("/places")
		admin.WeatherRouter(weather)

		placeAddress := v1Admin.Group("/places")
		admin.PlaceAddressRouter(placeAddress)

	}
}
