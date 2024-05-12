package mocks

import (
	"evoting/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockUserAddressRepository struct {
	mock.Mock
}

func (m *MockUserAddressRepository) GetDetailUser(id uuid.UUID) (*entities.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), nil
}

func (m *MockUserAddressRepository) Create(address *entities.UserAddress) (*entities.UserAddress, error) {
	args := m.Called(address)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.UserAddress), nil
}

func (m *MockUserAddressRepository) FindByUserId(id uuid.UUID) (*entities.UserAddress, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.UserAddress), nil
}

func (m *MockUserAddressRepository) Update(address *entities.UserAddress) (*entities.UserAddress, error) {
	args := m.Called(address)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.UserAddress), nil
}

func (m *MockUserAddressRepository) Delete(address *entities.UserAddress) error {
	args := m.Called(address)
	if args.Get(0) == nil {
		return nil
	}
	return args.Error(0)
}
