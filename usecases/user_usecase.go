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

type UserUsecase interface {
	Create(request *dto.CreateRequest) (*entities.User, error)
	Login(request *dto.LoginRequest) (*dto.LoginResponse, error)
	FindById(id uuid.UUID) (*entities.User, error)
	FindAll() (*[]entities.User, error)
	Update(id uuid.UUID, request *dto.UpdateRequest) (*entities.User, error)
	Delete(id uuid.UUID) error
}

type userUsecase struct {
	repository repositories.UserRepository
}

func NewUserUsecase(repository repositories.UserRepository) *userUsecase {
	return &userUsecase{repository}
}

func (uc *userUsecase) Create(request *dto.CreateRequest) (*entities.User, error) {
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
	newUser, err := uc.repository.Create(user)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	return newUser, nil
}

func (uc *userUsecase) Login(request *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := uc.repository.FindByEmail(request.Email)
	if err != nil {
		return nil, &errorHandlers.BadRequestError{Message: "Wrong email or password"}
	}
	if err := helpers.VerifyPassword(user.Password, request.Password); err != nil {
		return nil, &errorHandlers.BadRequestError{Message: "Wrong email or password"}
	}

	accessToken, err := helpers.GenerateAccessToken(user)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	refreshToken, err := helpers.GenerateRefreshToken(user)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}

	user.RefreshToken = refreshToken
	if err := uc.repository.SaveRefreshToken(user); err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	response := &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}

func (uc *userUsecase) FindById(id uuid.UUID) (*entities.User, error) {
	user, err := uc.repository.FindById(id)
	if err != nil {
		return nil, &errorHandlers.BadRequestError{Message: err.Error()}
	}
	return user, nil
}

func (uc *userUsecase) FindAll() (*[]entities.User, error) {
	users, err := uc.repository.FindAll()
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	return users, nil
}

func (uc *userUsecase) Update(id uuid.UUID, request *dto.UpdateRequest) (*entities.User, error) {
	user, err := uc.repository.FindById(id)
	if err != nil {
		return nil, &errorHandlers.BadRequestError{Message: err.Error()}
	}

	user.Fullname = request.Fullname
	user.Email = request.Email

	password, err := helpers.HashPassword(request.Password)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	user.Password = password

	updateUser, err := uc.repository.Update(user)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	return updateUser, nil
}

func (uc *userUsecase) Delete(id uuid.UUID) error {
	user, err := uc.repository.FindById(id)
	if err != nil {
		return &errorHandlers.BadRequestError{Message: err.Error()}
	}
	if err := uc.repository.Delete(user); err != nil {
		return &errorHandlers.BadRequestError{Message: err.Error()}
	}
	return nil
}
