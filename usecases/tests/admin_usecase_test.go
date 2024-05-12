package tests

import (
	"errors"
	"evoting/drivers/mocks"
	"evoting/dto"
	"evoting/entities"
	"evoting/helpers"
	"evoting/usecases"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestLogin(t *testing.T) {
	req := &dto.LoginAdminRequest{
		Email:    "admin@example.com",
		Password: "admin123",
	}
	t.Run("success", func(t *testing.T) {
		mockRepo := new(mocks.MockAdminRepository)
		usecase := usecases.NewAdminUsecase(mockRepo)
		password, _ := helpers.HashPassword(req.Password)
		expectedAdmin := &entities.Admin{
			Id:           uuid.New(),
			Email:        "admin@example.com",
			Username:     "admin",
			Password:     password,
			RefreshToken: "token",
		}

		mockRepo.On("FindByEmail", req.Email).Return(expectedAdmin, nil)
		mockRepo.On("SaveRefreshToken", mock.Anything).Return(nil)

		admin, err := usecase.Login(req)
		assert.NoError(t, err)
		assert.Equal(t, expectedAdmin.Id, admin.Id)
	})

	t.Run("failure_email_not_found", func(t *testing.T) {
		mockRepo := new(mocks.MockAdminRepository)
		usecase := usecases.NewAdminUsecase(mockRepo)
		mockRepo.On("FindByEmail", req.Email).Return(nil, errors.New("Wrong email or password"))
		admin, err := usecase.Login(req)
		assert.Error(t, err)
		assert.Nil(t, admin)
	})

	t.Run("failure_wrong_password", func(t *testing.T) {
		mockRepo := new(mocks.MockAdminRepository)
		usecase := usecases.NewAdminUsecase(mockRepo)

		expectedAdmin := &entities.Admin{
			Id:           uuid.New(),
			Email:        "admin@example.com",
			Username:     "admin",
			Password:     "admin123",
			RefreshToken: "token",
		}

		mockRepo.On("FindByEmail", req.Email).Return(expectedAdmin, nil)
		req.Password = "wrongpassword"

		admin, err := usecase.Login(req)
		assert.Error(t, err)
		assert.Nil(t, admin)
	})
}

func TestLogout(t *testing.T) {
	token := "validToken"

	t.Run("success", func(t *testing.T) {
		mockRepo := new(mocks.MockAdminRepository)
		usecase := usecases.NewAdminUsecase(mockRepo)

		mockRepo.On("GetUserByRefreshToken", token).Return(&entities.Admin{RefreshToken: token}, nil)

		mockRepo.On("SaveRefreshToken", mock.AnythingOfType("*entities.Admin")).Return(nil)

		err := usecase.Logout(token)
		assert.NoError(t, err)
	})

	t.Run("failure_invalid_token", func(t *testing.T) {
		mockRepo := new(mocks.MockAdminRepository)
		usecase := usecases.NewAdminUsecase(mockRepo)

		mockRepo.On("GetUserByRefreshToken", token).Return(nil, errors.New("Token is not valid"))

		err := usecase.Logout(token)
		assert.Error(t, err)
		assert.EqualError(t, err, "Token is not valid")
	})

	t.Run("failure_save_refresh_token", func(t *testing.T) {
		mockRepo := new(mocks.MockAdminRepository)
		usecase := usecases.NewAdminUsecase(mockRepo)

		mockRepo.On("GetUserByRefreshToken", token).Return(&entities.Admin{RefreshToken: token}, nil)

		mockRepo.On("SaveRefreshToken", mock.AnythingOfType("*entities.Admin")).Return(errors.New("Failed to save refresh token"))

		err := usecase.Logout(token)
		assert.Error(t, err)
		assert.EqualError(t, err, "Failed to save refresh token")
	})
}

func TestAdminUsecase_GetNewAccessToken(t *testing.T) {
	mockRepo := new(mocks.MockAdminRepository)
	usecase := usecases.NewAdminUsecase(mockRepo)

	t.Run("Test Token is not valid", func(t *testing.T) {
		tokenOld := "invalid_token"
		mockRepo.On("GetUserByRefreshToken", tokenOld).Return(nil, errors.New("Token is not valid"))

		newToken, err := usecase.GetNewAccessToken(tokenOld)

		assert.Error(t, err)
		assert.Nil(t, newToken)
		assert.EqualError(t, err, "Token is not valid")
	})

	t.Run("Test Token is empty", func(t *testing.T) {
		tokenOld := "empty_token"
		mockRepo.On("GetUserByRefreshToken", tokenOld).Return(&entities.Admin{RefreshToken: ""}, nil)

		newToken, err := usecase.GetNewAccessToken(tokenOld)

		assert.Error(t, err)
		assert.Nil(t, newToken)
		assert.EqualError(t, err, "Token is empty")
	})

	t.Run("Test GetNewAccessToken success", func(t *testing.T) {
		tokenOld := "valid_token"
		user := &entities.Admin{RefreshToken: "valid_token"}
		mockRepo.On("GetUserByRefreshToken", tokenOld).Return(user, nil)

		newToken, err := usecase.GetNewAccessToken(tokenOld)

		assert.NoError(t, err)
		assert.NotNil(t, newToken)

	})
}
