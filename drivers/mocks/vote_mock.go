package mocks

import (
	"evoting/dto"
	"evoting/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockVoteRepository struct {
	mock.Mock
}

func (m *MockVoteRepository) TotalVoters() (int64, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return 0, args.Error(1)
	}
	return args.Get(0).(int64), nil
}

func (m *MockVoteRepository) TotalVotesReceived() (int64, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return 0, args.Error(1)
	}
	return args.Get(0).(int64), nil
}

func (m *MockVoteRepository) FindUserById(id uuid.UUID) (*entities.Vote, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Vote), nil
}

func (m *MockVoteRepository) GetTotalVotes() (*[]dto.GetPlaceWithTotalVotes, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]dto.GetPlaceWithTotalVotes), nil
}

func (m *MockVoteRepository) Create(vote *entities.Vote) (*entities.Vote, error) {
	args := m.Called(vote)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Vote), nil
}
