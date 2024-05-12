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

func TestPlaceUsecase(t *testing.T) {
	placeId := uuid.New()
	placeRequest := &dto.PlaceRequest{
		Name:        "Test Place",
		Description: "Test Description",
		Longitude:   1.0,
		Latitude:    1.0,
	}
	placeAddress := &entities.PlaceAddress{
		Id:          uuid.New(),
		PlaceId:     placeId,
		Province:    "Test Province",
		City:        "Test City",
		SubDistrict: "Test Subdistrict",
		StreetName:  "Test Street",
	}

	t.Run("Test FindAll success", func(t *testing.T) {
		page := 1
		limit := 10
		sortBy := "email"
		sortType := "asc"
		searchQuery := ""

		expectedPlace := []entities.Place{
			{
				Id:          uuid.UUID{},
				Name:        "TEST1",
				Description: "TEST1",
				Longitude:   23.23,
				Latitude:    23.23,
				MapImage:    nil,
			},
			{
				Id:          uuid.UUID{},
				Name:        "TEST2",
				Description: "TEST2",
				Longitude:   23.23,
				Latitude:    23.23,
				MapImage:    nil,
			},
		}
		expectedTotal := int64(len(expectedPlace))
		mockRepo := new(mocks.MockPlaceRepository)
		usecase := usecases.NewPlaceUsecase(mockRepo)
		mockRepo.On("FindAll", page, limit, sortBy, sortType, searchQuery).Return(&expectedPlace, &expectedTotal, nil)
		places, total, err := usecase.FindAll(page, limit, sortBy, sortType, searchQuery)
		assert.NoError(t, err)
		assert.NotNil(t, places)
		assert.Equal(t, expectedTotal, *total)
	})

	t.Run("Test FindAll error", func(t *testing.T) {
		page := 1
		limit := 10
		sortBy := "email"
		sortType := "asc"
		searchQuery := ""

		mockRepo := new(mocks.MockPlaceRepository)
		usecase := usecases.NewPlaceUsecase(mockRepo)
		mockRepo.On("FindAll", page, limit, sortBy, sortType, searchQuery).Return(nil, nil, errors.New("find error"))
		places, total, err := usecase.FindAll(page, limit, sortBy, sortType, searchQuery)
		assert.Error(t, err)
		assert.Nil(t, places)
		assert.Nil(t, total)
	})

	t.Run("Test FindById success", func(t *testing.T) {
		mockRepo := new(mocks.MockPlaceRepository)
		usecase := usecases.NewPlaceUsecase(mockRepo)
		mockRepo.On("FindById", placeId).Return(&entities.Place{}, nil)
		place, err := usecase.FindById(placeId)
		assert.NoError(t, err)
		assert.NotNil(t, place)
	})

	t.Run("Test FindById error", func(t *testing.T) {
		mockRepo := new(mocks.MockPlaceRepository)
		usecase := usecases.NewPlaceUsecase(mockRepo)
		mockRepo.On("FindById", placeId).Return(nil, errors.New("find error"))
		place, err := usecase.FindById(placeId)
		assert.Error(t, err)
		assert.Nil(t, place)
	})

	t.Run("Test Create success", func(t *testing.T) {
		mockRepo := new(mocks.MockPlaceRepository)
		usecase := usecases.NewPlaceUsecase(mockRepo)
		mockRepo.On("FindByName", placeRequest.Name).Return(nil, nil)
		mockRepo.On("Create", mock.Anything).Return(&entities.Place{}, nil)
		newPlace, err := usecase.Create(placeRequest)
		assert.NoError(t, err)
		assert.NotNil(t, newPlace)
	})

	t.Run("Test Create error - Place already exists", func(t *testing.T) {
		request := &dto.PlaceRequest{
			Name:        "ExistingPlace",
			Description: "Description",
			Longitude:   1.0,
			Latitude:    1.0,
		}

		mockRepo := new(mocks.MockPlaceRepository)
		usecase := usecases.NewPlaceUsecase(mockRepo)
		mockRepo.On("FindByName", request.Name).Return(&entities.Place{Name: request.Name}, nil)

		newPlace, err := usecase.Create(request)
		assert.Error(t, err)
		assert.Nil(t, newPlace)
		assert.EqualError(t, err, "Place is already created, please enter another place")
	})

	t.Run("Test Create error - Repository error", func(t *testing.T) {
		request := &dto.PlaceRequest{
			Name:        "NewPlace",
			Description: "Description",
			Longitude:   1.0,
			Latitude:    1.0,
		}

		mockRepo := new(mocks.MockPlaceRepository)
		usecase := usecases.NewPlaceUsecase(mockRepo)
		mockRepo.On("FindByName", request.Name).Return(nil, errors.New("repository error"))
		mockRepo.On("Create", mock.Anything).Return(nil, errors.New("repository error"))
		newPlace, err := usecase.Create(request)
		assert.Error(t, err)
		assert.Nil(t, newPlace)
		assert.EqualError(t, err, "repository error")
	})

	t.Run("Test CreateAddress success", func(t *testing.T) {
		mockRepo := new(mocks.MockPlaceRepository)
		usecase := usecases.NewPlaceUsecase(mockRepo)
		mockRepo.On("CreateAddress", mock.Anything).Return(&entities.PlaceAddress{}, nil)
		newAddress, err := usecase.CreateAddress(placeAddress)
		assert.NoError(t, err)
		assert.NotNil(t, newAddress)
	})

	t.Run("Test CreateAddress error - Repository error", func(t *testing.T) {
		place := &entities.PlaceAddress{
			PlaceId:     uuid.New(),
			Province:    "Province",
			City:        "City",
			SubDistrict: "SubDistrict",
		}

		mockRepo := new(mocks.MockPlaceRepository)
		usecase := usecases.NewPlaceUsecase(mockRepo)
		mockRepo.On("CreateAddress", mock.Anything).Return(nil, errors.New("repository error"))

		newAddress, err := usecase.CreateAddress(place)
		assert.Error(t, err)
		assert.Nil(t, newAddress)
		assert.EqualError(t, err, "repository error")
	})

	t.Run("Test Update success", func(t *testing.T) {
		mockRepo := new(mocks.MockPlaceRepository)
		usecase := usecases.NewPlaceUsecase(mockRepo)
		mockRepo.On("FindById", placeId).Return(&entities.Place{}, nil)
		mockRepo.On("FindByName", placeRequest.Name).Return(nil, nil)
		mockRepo.On("Update", mock.Anything).Return(&entities.Place{}, nil)
		updatedPlace, err := usecase.Update(placeId, placeRequest)
		assert.NoError(t, err)
		assert.NotNil(t, updatedPlace)
	})

	t.Run("Test Update error - Place not found", func(t *testing.T) {
		id := uuid.New()
		request := &dto.PlaceRequest{
			Name:        "NewPlace",
			Description: "Description",
			Longitude:   1.0,
			Latitude:    1.0,
		}

		mockRepo := new(mocks.MockPlaceRepository)
		usecase := usecases.NewPlaceUsecase(mockRepo)
		mockRepo.On("FindById", id).Return(nil, errors.New("not found"))

		updatedPlace, err := usecase.Update(id, request)
		assert.Error(t, err)
		assert.Nil(t, updatedPlace)
		assert.EqualError(t, err, "not found")
	})

	t.Run("Test Update error - Place already exists", func(t *testing.T) {
		id := uuid.New()
		request := &dto.PlaceRequest{
			Name:        "ExistingPlace",
			Description: "Description",
			Longitude:   1.0,
			Latitude:    1.0,
		}

		existingPlace := &entities.Place{
			Id:          id,
			Name:        "ExistingPlace",
			Description: "ExistingDescription",
			Longitude:   2.0,
			Latitude:    2.0,
			MapImage:    nil,
		}

		mockRepo := new(mocks.MockPlaceRepository)
		usecase := usecases.NewPlaceUsecase(mockRepo)
		mockRepo.On("FindById", id).Return(existingPlace, nil)
		mockRepo.On("FindByName", request.Name).Return(existingPlace, nil)
		mockRepo.On("Update", mock.Anything).Return(nil, errors.New("Place is already created, please enter another place"))

		updatedPlace, err := usecase.Update(id, request)
		assert.Error(t, err)
		assert.Nil(t, updatedPlace)
		assert.EqualError(t, err, "Place is already created, please enter another place")
	})

	t.Run("Test Delete success", func(t *testing.T) {
		mockRepo := new(mocks.MockPlaceRepository)
		usecase := usecases.NewPlaceUsecase(mockRepo)
		mockRepo.On("FindById", placeId).Return(&entities.Place{}, nil)
		mockRepo.On("Delete", mock.Anything).Return(nil)
		err := usecase.Delete(placeId)
		assert.NoError(t, err)
	})

	t.Run("Test UpdateAddress success", func(t *testing.T) {
		mockRepo := new(mocks.MockPlaceRepository)
		usecase := usecases.NewPlaceUsecase(mockRepo)
		mockRepo.On("FindAddress", placeId).Return(&entities.PlaceAddress{}, nil)
		mockRepo.On("UpdateAddress", mock.Anything).Return(&entities.PlaceAddress{}, nil)
		updatedAddress, err := usecase.UpdateAddress(placeAddress)
		assert.NoError(t, err)
		assert.NotNil(t, updatedAddress)
	})
}
