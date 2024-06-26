package admin

import (
	"evoting/config"
	"evoting/handlers"
	"evoting/middlewares"
	"evoting/repositories"
	"evoting/usecases"
	"github.com/labstack/echo/v4"
)

func PlaceRouter(r *echo.Group) {
	r.Use(middlewares.JWTMiddleware)
	r.Use(middlewares.AdminOnlyMiddleware)
	repository := repositories.NewPlaceRepository(config.DB)
	usecase := usecases.NewPlaceUsecase(repository)
	handler := handlers.NewPlaceHandler(usecase)
	r.GET("", handler.FindAllPlaces)
	r.GET("/:id", handler.FindPlaceById)
	r.POST("", handler.CreatePlace)
	r.PUT("/:id", handler.UpdatePlace)
	r.DELETE("/:id", handler.Delete)
}
