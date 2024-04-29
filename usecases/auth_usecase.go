package usecases

import (
	"evoting/dto"
	"evoting/entities"
	"evoting/errorHandlers"
	"evoting/helpers"
	"evoting/repositories"
	"fmt"
	"github.com/google/uuid"
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

func (uc authUsecase) Register(request *dto.UserRequest) (*entities.User, error) {
	existingUser, _ := uc.repository.FindByEmail(request.Email)
	if existingUser != nil {
		return nil, &errorHandlers.BadRequestError{Message: "Register Failed: Email already used"}
	}

	hash, err := helpers.HashPassword(request.Password)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}

	user := &entities.User{
		Id:       uuid.New(),
		Email:    request.Email,
		Fullname: request.Fullname,
		Password: hash,
	}
	fmt.Println(user)
	newUser, err := uc.repository.CreateUser(user)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	return newUser, nil
}
