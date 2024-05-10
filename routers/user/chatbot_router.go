package user

import (
	"evoting/config"
	"evoting/handlers"
	"evoting/middlewares"
	"evoting/repositories"
	"evoting/usecases"
	"github.com/labstack/echo/v4"
)

func ChatbotRouter(r *echo.Group) {
	r.Use(middlewares.JWTMiddleware)
	placeRepository := repositories.NewPlaceRepository(config.DB)
	placeUsecae := usecases.NewPlaceUsecase(placeRepository)

	usecase := usecases.NewChatbotUseCase()
	handler := handlers.NewChatbotHandler(usecase, placeUsecae)
	r.GET("", handler.HandleWebSocket)
}
