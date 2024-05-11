package usecases

import (
	"evoting/dto"
	"evoting/errorHandlers"
	"evoting/helpers"
	"evoting/repositories"
)

type AdminUsecase interface {
	Login(request *dto.LoginAdminRequest) (*dto.LoginAdminResponse, error)
	Logout(token string) error
}

type adminUsecase struct {
	repository repositories.AdminRepository
}

func NewAdminUsecase(repository repositories.AdminRepository) *adminUsecase {
	return &adminUsecase{repository}
}

func (uc *adminUsecase) Login(request *dto.LoginAdminRequest) (*dto.LoginAdminResponse, error) {
	admin, err := uc.repository.FindByEmail(request.Email)
	if err != nil {
		return nil, &errorHandlers.BadRequestError{Message: "Wrong email or password"}
	}
	if err := helpers.VerifyPassword(admin.Password, request.Password); err != nil {
		return nil, &errorHandlers.BadRequestError{Message: "Wrong email or password"}
	}

	accessToken, err := helpers.GenerateAccessToken(admin)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	refreshToken, err := helpers.GenerateRefreshToken(admin)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}

	admin.RefreshToken = refreshToken
	if err := uc.repository.SaveRefreshToken(admin); err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	response := &dto.LoginAdminResponse{
		Id:           admin.Id,
		Username:     admin.Username,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}

func (uc *adminUsecase) Logout(token string) error {
	user, err := uc.repository.GetUserByRefreshToken(token)
	if err != nil {
		return &errorHandlers.UnAuthorizedError{Message: "Token is not valid"}
	}

	user.RefreshToken = ""
	if err := uc.repository.SaveRefreshToken(user); err != nil {
		return &errorHandlers.InternalServerError{Message: err.Error()}
	}

	return nil
}
