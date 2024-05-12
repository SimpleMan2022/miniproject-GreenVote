package mocks

import (
	"evoting/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockPlaceAddressRepository struct {
	mock.Mock
}

func (m MockPlaceAddressRepository) FindById(id uuid.UUID) (*entities.PlaceAddress, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.PlaceAddress), nil
}

func (m MockPlaceAddressRepository) Create(address *entities.PlaceAddress) (*entities.PlaceAddress, error) {
	args := m.Called(address)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.PlaceAddress), nil
}

func (m MockPlaceAddressRepository) Update(address *entities.PlaceAddress) (*entities.PlaceAddress, error) {
	args := m.Called(address)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.PlaceAddress), nil
}

func (m MockPlaceAddressRepository) Delete(address *entities.PlaceAddress) error {
	args := m.Called(address)
	if args.Get(0) == nil {
		return nil
	}
	return args.Error(0)
}
