package user

import (
	"evoting/config"
	"evoting/handlers"
	"evoting/middlewares"
	"evoting/repositories"
	"evoting/usecases"
	"github.com/labstack/echo/v4"
)

func CommentRouter(r *echo.Group) {
	r.Use(middlewares.JWTMiddleware)
	repository := repositories.NewCommentRepository(config.DB)
	usecase := usecases.NewCommentUsecase(repository)
	handler := handlers.NewCommentHandler(usecase)
	r.POST("/:id/comments", handler.CreateComment)
}
