package usecases

import (
	"evoting/dto"
	"evoting/entities"
	"evoting/errorHandlers"
	"evoting/repositories"
	"github.com/google/uuid"
	"time"
)

type WeatherDataUsecase interface {
	CreateOrUpdate(placeId uuid.UUID, request *dto.WeatherDataRequest) (*entities.WeatherData, error)
}

type weatherDataUsecase struct {
	repository repositories.WeatherDataRepository
}

func NewWeatherDataUsecase(repository repositories.WeatherDataRepository) *weatherDataUsecase {
	return &weatherDataUsecase{repository}
}

func (uc *weatherDataUsecase) CreateOrUpdate(placeId uuid.UUID, request *dto.WeatherDataRequest) (*entities.WeatherData, error) {
	exist, err := uc.repository.FindByPlaceId(placeId)
	if err != nil {
		return nil, &errorHandlers.BadRequestError{Message: err.Error()}
	}
	if exist.Id == uuid.Nil {
		weather := &entities.WeatherData{
			Id:          uuid.New(),
			PlaceId:     placeId,
			Temperature: request.Temperature,
			WindSpeed:   request.WindSpeed,
			Humadity:    request.Humadity,
			Summary:     request.Summary,
			RecordedAt:  time.Now(),
		}
		newWeather, err := uc.repository.Create(weather)
		if err != nil {
			return nil, &errorHandlers.InternalServerError{err.Error()}
		}
		return newWeather, nil
	} else {
		exist.Temperature = request.Temperature
		exist.Humadity = request.Humadity
		exist.WindSpeed = request.WindSpeed
		exist.Summary = request.Summary

		updateWeather, err := uc.repository.Update(exist)
		if err != nil {
			return nil, &errorHandlers.InternalServerError{Message: err.Error()}
		}
		return updateWeather, nilg
	}
}
