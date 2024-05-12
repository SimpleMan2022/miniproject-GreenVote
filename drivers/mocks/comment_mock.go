package mocks

import (
	"evoting/dto"
	"evoting/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockCommentRepository struct {
	mock.Mock
}

func (m *MockCommentRepository) FindById(id uuid.UUID) (*entities.Comment, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Comment), nil
}

func (m *MockCommentRepository) GetDetailPlace(id uuid.UUID) (*dto.CommentDetail, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.CommentDetail), nil
}

func (m *MockCommentRepository) FindByPlaceId(id uuid.UUID) (*[]dto.CommentData, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]dto.CommentData), nil
}

func (m *MockCommentRepository) Create(comment *entities.Comment) (*entities.Comment, error) {
	args := m.Called(comment)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Comment), nil
}

func (m *MockCommentRepository) Update(comment *entities.Comment) (*entities.Comment, error) {
	args := m.Called(comment)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Comment), nil
}

func (m *MockCommentRepository) Delete(comment *entities.Comment) error {
	args := m.Called(comment)
	if args.Get(0) == nil {
		return nil
	}
	return args.Error(0)
}
