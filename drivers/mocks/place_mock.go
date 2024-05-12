package mocks

import (
	"evoting/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockPlaceRepository struct {
	mock.Mock
}

func (m *MockPlaceRepository) Create(place *entities.Place) (*entities.Place, error) {
	args := m.Called(place)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Place), nil
}

func (m *MockPlaceRepository) CreateAddress(place *entities.PlaceAddress) (*entities.PlaceAddress, error) {
	args := m.Called(place)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.PlaceAddress), nil
}

func (m *MockPlaceRepository) FindAll(page, limit int, sortBy, sortType, searchQuery string) (*[]entities.Place, *int64, error) {
	args := m.Called(page, limit, sortBy, sortType, searchQuery)
	if args.Get(0) == nil {
		return nil, nil, args.Error(2)
	}
	return args.Get(0).(*[]entities.Place), args.Get(1).(*int64), nil
}

func (m *MockPlaceRepository) FindByName(place string) (*entities.Place, error) {
	args := m.Called(place)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Place), nil
}

func (m *MockPlaceRepository) FindById(id uuid.UUID) (*entities.Place, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Place), nil
}

func (m *MockPlaceRepository) FindAddress(id uuid.UUID) (*entities.PlaceAddress, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.PlaceAddress), nil
}

func (m *MockPlaceRepository) Update(place *entities.Place) (*entities.Place, error) {
	args := m.Called(place)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Place), nil
}

func (m *MockPlaceRepository) UpdateAddress(place *entities.PlaceAddress) (*entities.PlaceAddress, error) {
	args := m.Called(place)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.PlaceAddress), nil
}

func (m *MockPlaceRepository) Delete(place *entities.Place) error {
	args := m.Called(place)
	if args.Get(0) == nil {
		return nil
	}
	return args.Error(0)
}
