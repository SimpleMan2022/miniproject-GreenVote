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

func TestPlaceAddressUsecase(t *testing.T) {
	placeId := uuid.New()
	addressId := uuid.New()
	request := &dto.PlaceAddressRequest{
		Province:    "Sumatera Utara",
		City:        "Medan",
		SubDistrict: "Medan Timur",
		StreetName:  "Jalan Gajah Mada",
		ZipCode:     "20218",
	}
	address := &entities.PlaceAddress{
		Id:          addressId,
		PlaceId:     placeId,
		Province:    request.Province,
		City:        request.City,
		SubDistrict: request.SubDistrict,
		StreetName:  request.StreetName,
		ZipCode:     request.ZipCode,
	}

	t.Run("Test Create success", func(t *testing.T) {
		mockRepo := new(mocks.MockPlaceAddressRepository)
		usecase := usecases.NewPlaceAddressUsecase(mockRepo)
		mockRepo.On("Create", mock.Anything).Return(address, nil)
		newAddress, err := usecase.Create(placeId, request)
		assert.NoError(t, err)
		assert.NotNil(t, newAddress)
	})

	t.Run("Test Create error", func(t *testing.T) {
		mockRepo := new(mocks.MockPlaceAddressRepository)
		usecase := usecases.NewPlaceAddressUsecase(mockRepo)
		mockRepo.On("Create", mock.Anything).Return(nil, errors.New("create error"))
		newAddress, err := usecase.Create(placeId, request)
		assert.Error(t, err)
		assert.Nil(t, newAddress)
	})

	t.Run("Test Update success", func(t *testing.T) {
		mockRepo := new(mocks.MockPlaceAddressRepository)
		usecase := usecases.NewPlaceAddressUsecase(mockRepo)
		mockRepo.On("FindById", addressId).Return(address, nil)
		mockRepo.On("Update", address).Return(address, nil)
		updatedAddress, err := usecase.Update(addressId, placeId, request)
		assert.NoError(t, err)
		assert.NotNil(t, updatedAddress)
		assert.Equal(t, addressId, updatedAddress.Id)
	})

	t.Run("Test Update error", func(t *testing.T) {
		mockRepo := new(mocks.MockPlaceAddressRepository)
		usecase := usecases.NewPlaceAddressUsecase(mockRepo)
		mockRepo.On("FindById", addressId).Return(nil, errors.New("find error"))
		updatedAddress, err := usecase.Update(addressId, placeId, request)
		assert.Error(t, err)
		assert.Nil(t, updatedAddress)
	})

	t.Run("Test Delete success", func(t *testing.T) {
		mockRepo := new(mocks.MockPlaceAddressRepository)
		usecase := usecases.NewPlaceAddressUsecase(mockRepo)
		mockRepo.On("FindById", addressId).Return(address, nil)
		mockRepo.On("Delete", address).Return(nil)
		err := usecase.Delete(addressId, placeId)
		assert.NoError(t, err)
	})

	t.Run("Test Delete error", func(t *testing.T) {
		mockRepo := new(mocks.MockPlaceAddressRepository)
		usecase := usecases.NewPlaceAddressUsecase(mockRepo)
		mockRepo.On("FindById", addressId).Return(nil, errors.New("find error"))
		err := usecase.Delete(addressId, placeId)
		assert.Error(t, err)
		assert.EqualError(t, err, "find error")
	})

}
