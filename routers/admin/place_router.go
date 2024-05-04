package routers

import (
	"evoting/middlewares"
	"github.com/labstack/echo/v4"
	"net/http"
)

func PlaceRouter(r *echo.Group) {
	r.Use(middlewares.JWTMiddleware)
	r.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
