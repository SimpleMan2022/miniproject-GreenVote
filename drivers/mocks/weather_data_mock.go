package mocks

import (
	"evoting/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockWeatherDataRepository struct {
	mock.Mock
}

func (m *MockWeatherDataRepository) FindPlace(placeId uuid.UUID) (*entities.Place, error) {
	args := m.Called(placeId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Place), nil
}

func (m *MockWeatherDataRepository) FindByPlaceId(placeId uuid.UUID) (*entities.WeatherData, error) {
	args := m.Called(placeId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.WeatherData), nil
}

func (m *MockWeatherDataRepository) Create(data *entities.WeatherData) (*entities.WeatherData, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.WeatherData), nil
}

func (m *MockWeatherDataRepository) Update(data *entities.WeatherData) (*entities.WeatherData, error) {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.WeatherData), nil
}

func (m *MockWeatherDataRepository) Delete(data *entities.WeatherData) error {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil
	}
	return args.Error(0)
}
