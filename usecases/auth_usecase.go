package usecases

import (
	"evoting/dto"
	"evoting/entities"
	"evoting/repositories"
)

type AuthUsecase interface {
	Register(request *dto.UserRequest) (*entities.User, error)
}

type authUsecase struct {
	repository repositories.AuthRepository
}

func NewAuthUsecase(repository repositories.AuthRepository) *authUsecase {
	return &authUsecase{repository}
}
