package usecases

import (
	"evoting/dto"
	"evoting/entities"
	"evoting/errorHandlers"
	"evoting/helpers"
	"evoting/repositories"
	"github.com/google/uuid"
	"time"
)

type WeatherDataUsecase interface {
	GetPlace(placeId uuid.UUID) (*entities.Place, error)
	Create(request *dto.WeatherDataRequest) (*entities.WeatherData, error)
	Update(placeId uuid.UUID) (*entities.WeatherData, error)
	Delete(placeId uuid.UUID) error
}

type weatherDataUsecase struct {
	repository repositories.WeatherDataRepository
}

func NewWeatherDataUsecase(repository repositories.WeatherDataRepository) *weatherDataUsecase {
	return &weatherDataUsecase{repository}
}

func (uc *weatherDataUsecase) GetPlace(placeId uuid.UUID) (*entities.Place, error) {
	place, err := uc.repository.FindPlace(placeId)
	if err != nil {
		return nil, err
	}
	return place, nil
}

func (uc *weatherDataUsecase) Create(request *dto.WeatherDataRequest) (*entities.WeatherData, error) {
	place, err := uc.repository.FindPlace(request.PlaceId)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	existingWeather, _ := uc.repository.FindByPlaceId(request.PlaceId)
	if existingWeather != nil {
		return nil, &errorHandlers.BadRequestError{Message: "Weather data already exists for the specified placeId"}
	}

	weatherData, err := helpers.GenerateWeatherData(place.Latitude, place.Longitude)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}

	newWeather := &entities.WeatherData{
		Id:          uuid.New(),
		PlaceId:     request.PlaceId,
		Temperature: weatherData.Temperature,
		WindSpeed:   weatherData.WindSpeed,
		Humadity:    weatherData.Humadity,
		Condition:   weatherData.Condition,
		RecordedAt:  time.Now(),
	}

	createdWeather, err := uc.repository.Create(newWeather)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}

	return createdWeather, nil

}

func (uc *weatherDataUsecase) Update(placeId uuid.UUID) (*entities.WeatherData, error) {
	exist, err := uc.repository.FindByPlaceId(placeId)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}

	place, err := uc.repository.FindPlace(placeId)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}

	weatherData, err := helpers.GenerateWeatherData(place.Latitude, place.Longitude)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}

	if exist == nil {
		return nil, &errorHandlers.BadRequestError{Message: "Weather data not found for the specified placeId"}
	}

	exist.Temperature = weatherData.Temperature
	exist.Humadity = weatherData.Humadity
	exist.WindSpeed = weatherData.WindSpeed
	exist.Condition = weatherData.Condition
	exist.RecordedAt = time.Now()

	updatedWeather, err := uc.repository.Update(exist)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}

	return updatedWeather, nil
}

func (uc *weatherDataUsecase) Delete(placeId uuid.UUID) error {
	weather, err := uc.repository.FindByPlaceId(placeId)
	if err != nil {
		return &errorHandlers.BadRequestError{Message: err.Error()}
	}
	if err := uc.repository.Delete(weather); err != nil {
		return &errorHandlers.InternalServerError{Message: err.Error()}
	}
	return nil
}
