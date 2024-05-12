package mocks

import (
	"evoting/entities"
	"github.com/stretchr/testify/mock"
)

type MockAdminRepository struct {
	mock.Mock
}

func (m *MockAdminRepository) FindByEmail(email string) (*entities.Admin, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Admin), nil
}

func (m *MockAdminRepository) SaveRefreshToken(admin *entities.Admin) error {
	args := m.Called(admin)
	if args.Get(0) == nil {
		return nil
	}
	return args.Error(0)
}

func (m *MockAdminRepository) GetUserByRefreshToken(token string) (*entities.Admin, error) {
	args := m.Called(token)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Admin), nil
}
