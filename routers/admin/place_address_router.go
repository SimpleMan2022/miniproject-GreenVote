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
	r.Use(middlewares.AdminOnlyMiddleware)
	repository := repositories.NewPlaceAddressRepository(config.DB)
	usecase := usecases.NewPlaceAddressUsecase(repository)
	handler := handlers.NewPlaceAddressHandler(usecase)
	r.POST("/:placeId/address", handler.CreateAddress)
	r.PUT("/:placeId/address/:addressId", handler.UpdateAddress)
	r.DELETE("/:placeId/address/:addressId", handler.DeleteAddress)
}
