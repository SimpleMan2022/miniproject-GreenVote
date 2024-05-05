package admin

import (
	"evoting/config"
	"evoting/handlers"
	"evoting/middlewares"
	"evoting/repositories"
	"evoting/usecases"
	"github.com/labstack/echo/v4"
)

func PlaceAddressRouter(r *echo.Group) {
	r.Use(middlewares.JWTMiddleware)
	repository := repositories.NewPlaceAddressRepository(config.DB)
	usecase := usecases.NewPlaceAddressUsecase(repository)
	handler := handlers.NewPlaceAddressHandler(usecase)
	r.POST("", handler.CreateAddress)
	r.PUT("/:id", handler.UpdateAddress)
	r.DELETE("/:id", handler.DeleteAddress)
}
