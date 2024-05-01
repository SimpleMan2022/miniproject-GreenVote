package routers

import (
	"evoting/config"
	"evoting/handlers"
	"evoting/repositories"
	"evoting/usecases"
	"github.com/labstack/echo/v4"
)

func AuthRouter(r *echo.Group) {
	repository := repositories.NewUserRepository(config.DB)
	usecase := usecases.NewUserUsecase(repository)
	handler := handlers.NewUserHandler(usecase)

	r.POST("register", handler.Create)
	r.POST("login", handler.Login)
}
