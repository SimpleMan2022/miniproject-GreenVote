package admin

import (
	"evoting/config"
	"evoting/handlers"
	"evoting/middlewares"
	"evoting/repositories"
	"evoting/usecases"
	"github.com/labstack/echo/v4"
)

func WeatherRouter(r *echo.Group) {
	r.Use(middlewares.JWTMiddleware)
	repository := repositories.NewWeatherDataRepository(config.DB)
	usecase := usecases.NewWeatherDataUsecase(repository)
	handler := handlers.NewWeatherDataHandler(usecase)
	r.POST("/weather", handler.Create)
	r.PUT("/:id/weather", handler.Update)
	r.DELETE("/:id/weather", handler.Delete)
}
