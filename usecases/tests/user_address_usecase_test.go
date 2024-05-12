package tests

import (
	"errors"
	"evoting/drivers/mocks"
	"evoting/dto"
	"evoting/entities"
	"evoting/usecases"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestUserAddressUsecase(t *testing.T) {
	mockRepo := new(mocks.MockUserAddressRepository)
	usecase := usecases.NewUserAddressUsecase(mockRepo)

	t.Run("Test GetDetailUser success", func(t *testing.T) {
		userID := uuid.New()
		expectedUser := &entities.User{
			Id:       userID,
			Fullname: "John Doe",
		}
		mockRepo.On("GetDetailUser", userID).Return(expectedUser, nil)

		user, err := usecase.GetDetailUser(userID)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, expectedUser, user)
	})

	t.Run("Test GetDetailUser error", func(t *testing.T) {
		userID := uuid.New()
		mockRepo.On("GetDetailUser", userID).Return(nil, errors.New("user not found"))

		user, err := usecase.GetDetailUser(userID)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.EqualError(t, err, "user not found")
	})

	t.Run("Test Create success", func(t *testing.T) {
		req := &dto.CreateUserAddressRequest{
			UserId:      uuid.New(),
			Province:    "Province",
			City:        "City",
			SubDistrict: "SubDistrict",
			StreetName:  "StreetName",
			ZipCode:     "ZipCode",
		}
		mockRepo.On("FindByUserId", req.UserId).Return(nil, nil)
		mockRepo.On("Create", mock.Anything).Return(&entities.UserAddress{}, nil)

		address, err := usecase.Create(req)
		assert.NoError(t, err)
		assert.NotNil(t, address)
	})

	t.Run("Test Create error - User has existing address", func(t *testing.T) {
		req := &dto.CreateUserAddressRequest{
			UserId:      uuid.New(),
			Province:    "Province",
			City:        "City",
			SubDistrict: "SubDistrict",
			StreetName:  "StreetName",
			ZipCode:     "ZipCode",
		}
		existingUser := &entities.UserAddress{
			Id:          uuid.New(),
			UserId:      req.UserId,
			Province:    "ExistingProvince",
			City:        "ExistingCity",
			SubDistrict: "ExistingSubDistrict",
			StreetName:  "ExistingStreetName",
			ZipCode:     "ExistingZipCode",
		}
		mockRepo.On("FindByUserId", req.UserId).Return(existingUser, nil)

		address, err := usecase.Create(req)
		assert.Error(t, err)
		assert.Nil(t, address)
		assert.EqualError(t, err, "Failed to create address, user only have one address")
	})
	t.Run("Test FindByUserId success", func(t *testing.T) {
		userID := uuid.New()
		expectedAddress := &entities.UserAddress{
			Id:          uuid.New(),
			UserId:      userID,
			Province:    "Province",
			City:        "City",
			SubDistrict: "SubDistrict",
			StreetName:  "StreetName",
			ZipCode:     "ZipCode",
		}
		mockRepo.On("FindByUserId", userID).Return(expectedAddress, nil)

		address, err := usecase.FindByUserId(userID)
		assert.NoError(t, err)
		assert.NotNil(t, address)
		assert.Equal(t, expectedAddress, address)
	})

	t.Run("Test FindByUserId error", func(t *testing.T) {
		userID := uuid.New()
		mockRepo.On("FindByUserId", userID).Return(nil, errors.New("address not found"))

		address, err := usecase.FindByUserId(userID)
		assert.Error(t, err)
		assert.Nil(t, address)
		assert.EqualError(t, err, "address not found")
	})

	// Test Update
	t.Run("Test Update success", func(t *testing.T) {
		req := &dto.CreateUserAddressRequest{
			UserId:      uuid.New(),
			Province:    "NewProvince",
			City:        "NewCity",
			SubDistrict: "NewSubDistrict",
			StreetName:  "NewStreetName",
			ZipCode:     "NewZipCode",
		}
		existingAddress := &entities.UserAddress{
			Id:          uuid.New(),
			UserId:      req.UserId,
			Province:    "Province",
			City:        "City",
			SubDistrict: "SubDistrict",
			StreetName:  "StreetName",
			ZipCode:     "ZipCode",
		}
		mockRepo.On("FindByUserId", req.UserId).Return(existingAddress, nil)
		mockRepo.On("Update", mock.Anything).Return(&entities.UserAddress{}, nil)

		updatedAddress, err := usecase.Update(req)
		assert.NoError(t, err)
		assert.NotNil(t, updatedAddress)
	})

	// Test Delete
	t.Run("Test Delete success", func(t *testing.T) {
		userID := uuid.New()
		existingAddress := &entities.UserAddress{
			Id:          uuid.New(),
			UserId:      userID,
			Province:    "Province",
			City:        "City",
			SubDistrict: "SubDistrict",
			StreetName:  "StreetName",
			ZipCode:     "ZipCode",
		}
		mockRepo.On("FindByUserId", userID).Return(existingAddress, nil)
		mockRepo.On("Delete", existingAddress).Return(nil)

		err := usecase.Delete(userID)
		assert.NoError(t, err)
	})

	t.Run("Test Delete error", func(t *testing.T) {
		userID := uuid.New()
		mockRepo.On("FindByUserId", userID).Return(nil, errors.New("address not found"))

		err := usecase.Delete(userID)
		assert.Error(t, err)
		assert.EqualError(t, err, "address not found")
	})
}
