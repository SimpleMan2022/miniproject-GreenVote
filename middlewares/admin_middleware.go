package middlewares

import (
	"evoting/helpers"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func AdminOnlyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token is not provided")
		}

		if !strings.HasPrefix(authHeader, "bearer ") {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid token format. Use Bearer token")
		}
		tokenStr := strings.TrimPrefix(authHeader, "bearer ")

		claims, err := helpers.ParseJWT(tokenStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		isAdmin := claims.IsAdmin
		fmt.Println(isAdmin)
		if !isAdmin {
			return echo.NewHTTPError(http.StatusForbidden, "Only admins are allowed")
		}

		return next(c)
	}
}
