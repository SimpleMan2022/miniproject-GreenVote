package usecases

import (
	"evoting/dto"
	"evoting/entities"
	"evoting/errorHandlers"
	"evoting/helpers"
	"evoting/repositories"
	"github.com/google/uuid"
)

type UserUsecase interface {
	Create(request *dto.CreateRequest) (*entities.User, error)
	Login(request *dto.LoginRequest) (*dto.LoginResponse, error)
	FindById(id uuid.UUID) (*entities.User, error)
	FindAll(page, limit int, sortBy, sortType string) (*[]entities.User, *int64, error)
	FindSoftDelete(page, limit int, sortBy, sortType string) (*[]entities.User, *int64, error)
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
		Id:           user.Id,
		Fullname:     user.Fullname,
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

func (uc *userUsecase) FindAll(page, limit int, sortBy, sortType string) (*[]entities.User, *int64, error) {
	users, total, err := uc.repository.FindAll(page, limit, sortBy, sortType)
	if err != nil {
		return nil, nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	return users, total, nil
}

func (uc *userUsecase) FindSoftDelete(page, limit int, sortBy, sortType string) (*[]entities.User, *int64, error) {
	users, total, err := uc.repository.FindSoftDelete(page, limit, sortBy, sortType)
	if err != nil {
		return nil, nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	return users, total, nil
}

func (uc *userUsecase) Update(id uuid.UUID, request *dto.UpdateRequest) (*entities.User, error) {
	user, err := uc.repository.FindById(id)
	if err != nil {
		return nil, &errorHandlers.BadRequestError{Message: err.Error()}
	}
	if user.Email != request.Email {
		existingUser, _ := uc.repository.FindByEmail(request.Email)
		if existingUser != nil {
			return nil, &errorHandlers.BadRequestError{Message: "Update Failed: Email already used"}
		}
	}

	user.Fullname = request.Fullname
	user.Email = request.Email

	password, err := helpers.HashPassword(request.Password)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	user.Password = password

	if request.Image != nil {
		if imageType := helpers.IsValidImageType(request.Image); !imageType {
			return nil, &errorHandlers.BadRequestError{Message: "Invalid image format. Only JPG, JPEG and PNG are allowed."}
		}

		if size := helpers.IsValidImageSize(request.Image); !size {
			return nil, &errorHandlers.BadRequestError{Message: "Image size exceeds the limit of 2MB."}
		}
		if user.Image != nil {
			if err := helpers.DeleteImage("public/images/users", user.Image); err != nil {
				return nil, &errorHandlers.InternalServerError{Message: err.Error()}
			}
		}
		filename, err := helpers.UploadUserImage(request.Image)
		if err != nil {
			return nil, &errorHandlers.InternalServerError{Message: err.Error()}
		}
		user.Image = &filename
	}

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
	if user.Image != nil {
		if err := helpers.DeleteImage("public/images/users", user.Image); err != nil {
			return &errorHandlers.InternalServerError{Message: err.Error()}
		}
	}
	if err := uc.repository.Delete(user); err != nil {
		return &errorHandlers.BadRequestError{Message: err.Error()}
	}
	return nil
}
