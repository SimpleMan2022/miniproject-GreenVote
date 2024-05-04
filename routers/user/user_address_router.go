package user

import (
	"evoting/config"
	"evoting/handlers"
	"evoting/middlewares"
	"evoting/repositories"
	"evoting/usecases"
	"github.com/labstack/echo/v4"
)

func UserAddressRouter(r *echo.Group) {
	repositoryAddress := repositories.NewUserAddressRepository(config.DB)
	usecaseAddress := usecases.NewUserAddressUsecase(repositoryAddress)
	handlerAddress := handlers.NewUserAddressHandler(usecaseAddress)

	repositoryUser := repositories.NewUserRepository(config.DB)
	usecaseUser := usecases.NewUserUsecase(repositoryUser)
	handlerUser := handlers.NewUserHandler(usecaseUser)

	r.Use(middlewares.JWTMiddleware)
	r.POST("/address", handlerAddress.Create)
	r.GET("/address", handlerAddress.GetDetailUser)
	r.PUT("/:id", handlerUser.UpdateUser)
	r.PUT("/address", handlerAddress.Update)
	r.DELETE("/address", handlerAddress.Delete)

}
