package mocks

import (
	"evoting/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByEmail(email string) (*entities.User, error) {
	args := m.Called(email)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserRepository) FindById(id uuid.UUID) (*entities.User, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserRepository) FindAll(page, limit int, sortBy, sortType, searchQuery string) (*[]entities.User, *int64, error) {
	args := m.Called(page, limit, sortBy, sortType, searchQuery)
	return args.Get(0).(*[]entities.User), args.Get(1).(*int64), args.Error(2)
}

func (m *MockUserRepository) Create(user *entities.User) (*entities.User, error) {
	args := m.Called(user)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserRepository) SaveRefreshToken(user *entities.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Update(user *entities.User) (*entities.User, error) {
	args := m.Called(user)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserRepository) Delete(user *entities.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByRefreshToken(token string) (*entities.User, error) {
	args := m.Called(token)
	return args.Get(0).(*entities.User), args.Error(1)
}
