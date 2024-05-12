package routers

import (
	"evoting/config"
	"evoting/handlers"
	"evoting/repositories"
	"evoting/usecases"
	"github.com/labstack/echo/v4"
)

func AuthUserRouter(r *echo.Group) {
	repository := repositories.NewUserRepository(config.DB)
	usecase := usecases.NewUserUsecase(repository)
	handler := handlers.NewUserHandler(usecase)

	r.POST("/register", handler.Create)
	r.POST("/login", handler.Login)
	r.DELETE("/logout", handler.Logout)
	r.GET("/token", handler.GetNewAccessToken)

}

func AuthAdminRouter(r *echo.Group) {
	repository := repositories.NewAdminRepository(config.DB)
	usecase := usecases.NewAdminUsecase(repository)
	handler := handlers.NewAdminHandle(usecase)

	r.POST("/login", handler.Login)
	r.DELETE("/logout", handler.Logout)
	r.GET("/token", handler.GetNewAccessToken)
}
