package routers

import (
	"evoting/config"
	"evoting/handlers"
	"evoting/middlewares"
	"evoting/repositories"
	"evoting/usecases"
	"github.com/labstack/echo/v4"
)

func UserRouter(r *echo.Group) {
	repository := repositories.NewUserRepository(config.DB)
	usecase := usecases.NewUserUsecase(repository)
	handler := handlers.NewUserHandler(usecase)
	r.Use(middlewares.JWTMiddleware)
	r.GET("", handler.FindAllUsers)
	r.GET("/:id", handler.FindUserById)
	r.POST("/create", handler.Create)
	r.PUT("/:id", handler.UpdateUser)
	r.DELETE("/:id", handler.DeleteUser)

}
