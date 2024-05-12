package tests

import (
	"errors"
	"evoting/drivers/mocks"
	"evoting/dto"
	"evoting/entities"
	"evoting/errorHandlers"
	"evoting/usecases"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestWeatherDataUsecase(t *testing.T) {
	t.Run("Test GetPlace Success", func(t *testing.T) {
		placeId := uuid.New()
		expectedPlace := &entities.Place{
			Id:        placeId,
			Name:      "Test Place",
			Latitude:  2.1,
			Longitude: 3.2,
		}
		mockRepo := new(mocks.MockWeatherDataRepository)
		usecase := usecases.NewWeatherDataUsecase(mockRepo)
		mockRepo.On("FindPlace", placeId).Return(expectedPlace, nil)

		place, err := usecase.GetPlace(placeId)

		assert.NoError(t, err)
		assert.NotNil(t, place)
		assert.Equal(t, expectedPlace, place)
	})

	//t.Run("Test Create Success", func(t *testing.T) {
	//	mockRepo := new(mocks.MockWeatherDataRepository)
	//	usecase := usecases.NewWeatherDataUsecase(mockRepo)
	//	place := &entities.Place{
	//		Id:          uuid.New(),
	//		Name:        "Place 1",
	//		Description: "Desc 1",
	//		Longitude:   0.91601259,
	//		Latitude:    0.91601259,
	//		MapImage:    nil,
	//	}
	//
	//	newWeather := &entities.WeatherData{
	//		Id:          uuid.New(),
	//		PlaceId:     place.Id,
	//		Temperature: 28,
	//		WindSpeed:   29.1,
	//		Humadity:    29.1,
	//		Condition:   "Clear",
	//		RecordedAt:  time.Now(),
	//	}
	//	mockRepo.On("FindPlace", place.Id).Return(place, nil)
	//	mockRepo.On("FindByPlaceId", place.Id).Return(nil, nil)
	//	mockRepo.On("Create", mock.Anything).Return(newWeather, nil)
	//
	//	request := &dto.WeatherDataRequest{
	//		PlaceId: place.Id,
	//	}
	//
	//	weather, err := usecase.Create(request)
	//	fmt.Println(weather)
	//	assert.NoError(t, err)
	//	assert.NotNil(t, weather)
	//})

	t.Run("Test Create Fail - Existing Weather Data", func(t *testing.T) {
		// Test setup
		placeId := uuid.New()
		mockRepo := new(mocks.MockWeatherDataRepository)
		usecase := usecases.NewWeatherDataUsecase(mockRepo)
		mockRepo.On("FindPlace", placeId).Return(&entities.Place{Id: placeId, Latitude: 0.0, Longitude: 0.0}, nil)
		mockRepo.On("FindByPlaceId", placeId).Return(&entities.WeatherData{}, nil)

		request := &dto.WeatherDataRequest{
			PlaceId: placeId,
		}

		// Test execution
		weather, err := usecase.Create(request)

		// Test assertion
		assert.Error(t, err)
		assert.Nil(t, weather)
		assert.IsType(t, &errorHandlers.BadRequestError{}, err)
	})

	t.Run("Test Create Fail - FindPlace Error", func(t *testing.T) {
		// Test setup
		placeId := uuid.New()
		mockRepo := new(mocks.MockWeatherDataRepository)
		usecase := usecases.NewWeatherDataUsecase(mockRepo)
		mockRepo.On("FindPlace", placeId).Return(nil, errors.New("place not found"))

		request := &dto.WeatherDataRequest{
			PlaceId: placeId,
		}

		// Test execution
		weather, err := usecase.Create(request)

		// Test assertion
		assert.Error(t, err)
		assert.Nil(t, weather)
		assert.IsType(t, &errorHandlers.InternalServerError{}, err)
	})

	//t.Run("Test Update Success", func(t *testing.T) {
	//	// Test setup
	//	placeId := uuid.New()
	//	mockRepo := new(mocks.MockWeatherDataRepository)
	//	usecase := usecases.NewWeatherDataUsecase(mockRepo)
	//	mockRepo.On("FindByPlaceId", placeId).Return(&entities.WeatherData{}, nil)
	//	mockRepo.On("FindPlace", placeId).Return(&entities.Place{Id: placeId, Latitude: 0.0, Longitude: 0.0}, nil)
	//	mockRepo.On("Update", mock.Anything).Return(&entities.WeatherData{}, nil)
	//
	//	// Test execution
	//	weather, err := usecase.Update(placeId)
	//
	//	// Test assertion
	//	assert.NoError(t, err)
	//	assert.NotNil(t, weather)
	//})

	t.Run("Test Update Fail - FindByPlaceId Error", func(t *testing.T) {
		// Test setup
		placeId := uuid.New()
		mockRepo := new(mocks.MockWeatherDataRepository)
		usecase := usecases.NewWeatherDataUsecase(mockRepo)
		mockRepo.On("FindByPlaceId", placeId).Return(nil, errors.New("weather data not found"))

		// Test execution
		weather, err := usecase.Update(placeId)

		// Test assertion
		assert.Error(t, err)
		assert.Nil(t, weather)
		assert.IsType(t, &errorHandlers.InternalServerError{}, err)
	})
	t.Run("Test Delete Success", func(t *testing.T) {
		// Test setup
		placeId := uuid.New()
		mockRepo := new(mocks.MockWeatherDataRepository)
		usecase := usecases.NewWeatherDataUsecase(mockRepo)
		mockRepo.On("FindByPlaceId", placeId).Return(&entities.WeatherData{}, nil)
		mockRepo.On("Delete", mock.Anything).Return(nil)

		// Test execution
		err := usecase.Delete(placeId)

		// Test assertion
		assert.NoError(t, err)
	})

	t.Run("Test Delete Fail - FindByPlaceId Error", func(t *testing.T) {
		// Test setup
		placeId := uuid.New()
		mockRepo := new(mocks.MockWeatherDataRepository)
		usecase := usecases.NewWeatherDataUsecase(mockRepo)
		mockRepo.On("FindByPlaceId", placeId).Return(nil, errors.New("weather data not found"))

		// Test execution
		err := usecase.Delete(placeId)

		// Test assertion
		assert.Error(t, err)
		assert.IsType(t, &errorHandlers.BadRequestError{}, err)
	})

	t.Run("Test Delete Fail - Repository Delete Error", func(t *testing.T) {
		// Test setup
		placeId := uuid.New()
		mockRepo := new(mocks.MockWeatherDataRepository)
		usecase := usecases.NewWeatherDataUsecase(mockRepo)
		mockRepo.On("FindByPlaceId", placeId).Return(&entities.WeatherData{}, nil)
		mockRepo.On("Delete", mock.Anything).Return(errors.New("failed to delete weather data"))

		// Test execution
		err := usecase.Delete(placeId)

		// Test assertion
		assert.Error(t, err)
		assert.IsType(t, &errorHandlers.InternalServerError{}, err)
	})

}
