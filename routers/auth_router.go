package routers

import (
	"evoting/config"
	"evoting/handlers"
	"evoting/repositories"
	"evoting/usecases"
	"github.com/labstack/echo/v4"
)

func AuthRouter(c *echo.Group) {
	repository := repositories.NewAuthRepository(config.DB)
	usecase := usecases.NewAuthUsecase(repository)
	handler := handlers.NewAuthHandler(usecase)

	c.POST("register", handler.Register)
}
