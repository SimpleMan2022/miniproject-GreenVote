package user

import (
	"evoting/config"
	"evoting/handlers"
	"evoting/middlewares"
	"evoting/repositories"
	"evoting/usecases"
	"github.com/labstack/echo/v4"
)

func VoteRouter(r *echo.Group) {
	r.Use(middlewares.JWTMiddleware)
	repository := repositories.NewvoteRepository(config.DB)
	usecase := usecases.NewVoteUsecase(repository)
	handler := handlers.NewVoteHandler(usecase)
	r.POST("", handler.CreateVote)
	r.GET("", handler.GetPlaceWithTotalVotes)
}
