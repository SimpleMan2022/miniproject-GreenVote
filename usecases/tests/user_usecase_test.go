package tests

import (
	"errors"
	"evoting/drivers/mysql/mocks"
	"evoting/dto"
	"evoting/entities"
	"evoting/errorHandlers"
	"evoting/helpers"
	"evoting/usecases"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mime/multipart"
	"testing"
)

func TestUserUsecase_Create(t *testing.T) {
	req := &dto.CreateRequest{
		Email:    "admin@example.com",
		Password: "admin123",
		Fullname: "Admin",
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		uc := usecases.NewUserUsecase(mockRepo)

		mockRepo.On("FindByEmail", req.Email).Return(nil, nil)
		mockRepo.On("Create", mock.Anything).Return(&entities.User{Id: uuid.New(), Email: req.Email, Fullname: req.Fullname}, nil)

		newUser, err := uc.Create(req)
		assert.NoError(t, err)
		assert.NotNil(t, newUser)
		assert.Equal(t, newUser.Email, req.Email)
	})

	t.Run("Email already used", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		uc := usecases.NewUserUsecase(mockRepo)

		existingUser := &entities.User{Id: uuid.New(), Email: req.Email, Fullname: req.Fullname}
		mockRepo.On("FindByEmail", req.Email).Return(existingUser, nil)

		newUser, err := uc.Create(req)
		assert.Error(t, err)
		assert.Nil(t, newUser)
		assert.Equal(t, err.Error(), "Register Failed: Email already used")
	})
	t.Run("Failed Create", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		uc := usecases.NewUserUsecase(mockRepo)

		mockRepo.On("FindByEmail", req.Email).Return(nil, nil)
		mockRepo.On("Create", mock.Anything).Return(nil, errors.New("Error creating user"))

		newUser, err := uc.Create(req)
		assert.Error(t, err)
		assert.Nil(t, newUser)
		assert.Equal(t, "Error creating user", err.Error())
	})
}

func TestUserUsecase_Login(t *testing.T) {
	req := &dto.LoginRequest{
		Email:    "admin@example.com",
		Password: "admin123",
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		uc := usecases.NewUserUsecase(mockRepo)
		password, _ := helpers.HashPassword(req.Password)

		expectedUser := &entities.User{
			Id:       uuid.New(),
			Email:    req.Email,
			Fullname: "example",
			Password: password,
			Image:    nil,
		}
		mockRepo.On("FindByEmail", req.Email).Return(expectedUser, nil)

		mockRepo.On("SaveRefreshToken", mock.Anything).Return(nil)
		expectedUser.RefreshToken = "token"
		resp, err := uc.Login(req)
		assert.NotNil(t, resp)
		assert.Equal(t, expectedUser.Id, resp.Id)
		assert.NoError(t, err)
		assert.Equal(t, expectedUser.Fullname, resp.Fullname)
		assert.NotEmpty(t, resp.AccessToken)
		assert.NotEmpty(t, resp.RefreshToken)
	})

	t.Run("User not found", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		uc := usecases.NewUserUsecase(mockRepo)

		mockRepo.On("FindByEmail", req.Email).Return(nil, errors.New("User not found"))

		resp, err := uc.Login(req)
		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, err.Error(), "Wrong email or password")
	})

	t.Run("Wrong password", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		uc := usecases.NewUserUsecase(mockRepo)

		user := &entities.User{Id: uuid.New(), Email: req.Email, Fullname: "Admin", Password: "$2a$10$5LRrZHz.Xp.zyVSFt1lFXOznIq9h2NQ7BWr0YpkPcKUKcP4tFjx1a"}
		mockRepo.On("FindByEmail", req.Email).Return(user, nil)

		req.Password = "wrongpassword"

		resp, err := uc.Login(req)
		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, err.Error(), "Wrong email or password")
	})
}

func TestUserUsecase_FindById(t *testing.T) {
	userID := uuid.New()
	expectedUser := &entities.User{
		Id:       userID,
		Email:    "admin@example.com",
		Fullname: "Admin",
		Password: "hashedpassword",
		Image:    nil,
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		uc := usecases.NewUserUsecase(mockRepo)

		mockRepo.On("FindById", userID).Return(expectedUser, nil)

		user, err := uc.FindById(userID)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, expectedUser.Id, user.Id)
		assert.Equal(t, expectedUser.Email, user.Email)
		assert.Equal(t, expectedUser.Fullname, user.Fullname)
		assert.Equal(t, expectedUser.Password, user.Password)
		assert.Equal(t, expectedUser.Image, user.Image)
	})

	t.Run("Failed", func(t *testing.T) {
		expectedError := &errorHandlers.BadRequestError{Message: "User not found"}

		mockRepo := new(mocks.MockUserRepository)
		uc := usecases.NewUserUsecase(mockRepo)

		mockRepo.On("FindById", userID).Return(nil, expectedError)

		user, err := uc.FindById(userID)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, expectedError, err)
	})
}

func TestUserUsecase_FindAll(t *testing.T) {
	page := 1
	limit := 10
	sortBy := "email"
	sortType := "asc"
	searchQuery := ""

	expectedUsers := []entities.User{
		{Id: uuid.New(), Email: "user1@example.com", Fullname: "User 1", Password: "hashedpassword1", Image: nil},
		{Id: uuid.New(), Email: "user2@example.com", Fullname: "User 2", Password: "hashedpassword2", Image: nil},
	}
	total := int64(len(expectedUsers))

	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		uc := usecases.NewUserUsecase(mockRepo)

		mockRepo.On("FindAll", page, limit, sortBy, sortType, searchQuery).Return(&expectedUsers, &total, nil)

		users, total, err := uc.FindAll(page, limit, sortBy, sortType, searchQuery)
		assert.NoError(t, err)
		assert.NotNil(t, users)
		assert.Equal(t, len(expectedUsers), len(*users))
		assert.Equal(t, total, total)
	})

	t.Run("Failed", func(t *testing.T) {
		expectedError := &errorHandlers.InternalServerError{Message: "Failed to find users"}

		mockRepo := new(mocks.MockUserRepository)
		uc := usecases.NewUserUsecase(mockRepo)

		mockRepo.On("FindAll", page, limit, sortBy, sortType, searchQuery).Return(nil, nil, expectedError)

		users, total, err := uc.FindAll(page, limit, sortBy, sortType, searchQuery)
		assert.Error(t, err)
		assert.Nil(t, users)
		assert.Nil(t, total)
		assert.Equal(t, expectedError, err)
	})

}

func TestUserUsecase_Update(t *testing.T) {
	userID := uuid.New()
	request := &dto.UpdateRequest{
		Email:    "newemail@example.com",
		Fullname: "New Name",
		Password: "newpassword",
		Image:    nil,
	}

	t.Run("Success", func(t *testing.T) {
		existingUser := &entities.User{
			Id:       userID,
			Email:    "oldemail@example.com",
			Fullname: "Old Name",
			Password: "oldhashedpassword",
			Image:    nil,
		}

		mockRepo := new(mocks.MockUserRepository)
		uc := usecases.NewUserUsecase(mockRepo)

		mockRepo.On("FindById", userID).Return(existingUser, nil)
		mockRepo.On("FindByEmail", request.Email).Return(nil, nil)
		mockRepo.On("Update", existingUser).Return(existingUser, nil)

		updatedUser, err := uc.Update(userID, request)
		assert.NoError(t, err)
		assert.NotNil(t, updatedUser)
		assert.Equal(t, request.Email, updatedUser.Email)
		assert.Equal(t, request.Fullname, updatedUser.Fullname)
	})

	t.Run("Email already used", func(t *testing.T) {
		existingUser := &entities.User{
			Id:       userID,
			Email:    "oldemail@example.com",
			Fullname: "Old Name",
			Password: "oldhashedpassword",
			Image:    nil,
		}

		mockRepo := new(mocks.MockUserRepository)
		uc := usecases.NewUserUsecase(mockRepo)

		mockRepo.On("FindById", userID).Return(existingUser, nil)
		mockRepo.On("FindByEmail", request.Email).Return(existingUser, nil)

		updatedUser, err := uc.Update(userID, request)
		assert.Error(t, err)
		assert.Nil(t, updatedUser)
		assert.Equal(t, "Update Failed: Email already used", err.Error())
	})
	t.Run("Invalid image format", func(t *testing.T) {
		var image multipart.FileHeader
		image.Filename = "invalidimage"
		requestWithInvalidImage := &dto.UpdateRequest{
			Email:    "newemail@example.com",
			Fullname: "New Name",
			Password: "newpassword",
			Image:    &image,
		}

		existingUser := &entities.User{
			Id:       userID,
			Email:    "oldemail@example.com",
			Fullname: "Old Name",
			Password: "oldhashedpassword",
			Image:    nil,
		}

		mockRepo := new(mocks.MockUserRepository)
		uc := usecases.NewUserUsecase(mockRepo)

		mockRepo.On("FindById", userID).Return(existingUser, nil)
		mockRepo.On("FindByEmail", "newemail@example.com").Return(nil, nil)

		updatedUser, err := uc.Update(userID, requestWithInvalidImage)
		assert.Error(t, err)
		assert.Nil(t, updatedUser)
		assert.Equal(t, "Invalid image format. Only JPG, JPEG and PNG are allowed.", err.Error())
	})

	t.Run("FindById error", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		uc := usecases.NewUserUsecase(mockRepo)

		mockRepo.On("FindById", userID).Return(nil, errors.New("FindById error"))

		updatedUser, err := uc.Update(userID, request)
		assert.Error(t, err)
		assert.Nil(t, updatedUser)
		assert.Equal(t, "FindById error", err.Error())
	})

	t.Run("Update error", func(t *testing.T) {
		existingUser := &entities.User{
			Id:       userID,
			Email:    "oldemail@example.com",
			Fullname: "Old Name",
			Password: "oldhashedpassword",
			Image:    nil,
		}

		mockRepo := new(mocks.MockUserRepository)
		uc := usecases.NewUserUsecase(mockRepo)

		mockRepo.On("FindById", userID).Return(existingUser, nil)
		mockRepo.On("FindByEmail", request.Email).Return(nil, nil)
		mockRepo.On("Update", existingUser).Return(nil, errors.New("Update error"))

		updatedUser, err := uc.Update(userID, request)
		assert.Error(t, err)
		assert.Nil(t, updatedUser)
		assert.Equal(t, "Update error", err.Error())
	})

}

func TestUserUsecase_Delete(t *testing.T) {
	userID := uuid.New()

	t.Run("Success", func(t *testing.T) {
		existingUser := &entities.User{
			Id:       userID,
			Email:    "user@example.com",
			Fullname: "User",
			Password: "hashedpassword",
			Image:    nil,
		}

		mockRepo := new(mocks.MockUserRepository)
		uc := usecases.NewUserUsecase(mockRepo)

		mockRepo.On("FindById", userID).Return(existingUser, nil)
		mockRepo.On("Delete", existingUser).Return(nil)

		err := uc.Delete(userID)
		assert.NoError(t, err)
	})

	t.Run("User not found", func(t *testing.T) {
		expectedError := &errorHandlers.BadRequestError{Message: "User not found"}
		mockRepo := new(mocks.MockUserRepository)
		uc := usecases.NewUserUsecase(mockRepo)

		mockRepo.On("FindById", userID).Return(nil, expectedError)

		err := uc.Delete(userID)
		assert.Error(t, err)
		assert.Equal(t, "User not found", err.Error())
	})
	t.Run("Delete error", func(t *testing.T) {
		existingUser := &entities.User{
			Id:       userID,
			Email:    "oldemail@example.com",
			Fullname: "Old Name",
			Password: "oldhashedpassword",
			Image:    nil,
		}

		mockRepo := new(mocks.MockUserRepository)
		uc := usecases.NewUserUsecase(mockRepo)

		mockRepo.On("FindById", userID).Return(existingUser, nil)
		mockRepo.On("Delete", existingUser).Return(errors.New("Delete error"))

		err := uc.Delete(userID)
		assert.Error(t, err)
		assert.Equal(t, "Delete error", err.Error())
	})

}

func TestUserUsecase_Logout(t *testing.T) {
	refreshToken := "valid_refresh_token"

	t.Run("Success", func(t *testing.T) {
		existingUser := &entities.User{
			Id:           uuid.New(),
			Email:        "user@example.com",
			Fullname:     "User",
			Password:     "hashedpassword",
			Image:        nil,
			RefreshToken: refreshToken,
		}

		mockRepo := new(mocks.MockUserRepository)
		uc := usecases.NewUserUsecase(mockRepo)

		mockRepo.On("GetUserByRefreshToken", refreshToken).Return(existingUser, nil)
		mockRepo.On("SaveRefreshToken", existingUser).Return(nil)

		err := uc.Logout(refreshToken)
		assert.NoError(t, err)
		assert.Empty(t, existingUser.RefreshToken)
	})

	t.Run("Token is not valid", func(t *testing.T) {
		expectedError := &errorHandlers.InternalServerError{Message: "token not valid"}
		mockRepo := new(mocks.MockUserRepository)
		uc := usecases.NewUserUsecase(mockRepo)

		mockRepo.On("GetUserByRefreshToken", refreshToken).Return(nil, expectedError)

		err := uc.Logout(refreshToken)
		assert.Error(t, err)
		assert.Equal(t, "Token is not valid", err.Error())
	})

}
