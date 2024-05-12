package tests

import (
	"errors"
	"evoting/drivers/mocks"
	"evoting/dto"
	"evoting/entities"
	"evoting/usecases"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestVoteUsecase(t *testing.T) {

	t.Run("Test GetPlaceWithTotalVotes success", func(t *testing.T) {
		mockRepo := new(mocks.MockVoteRepository)
		usecase := usecases.NewVoteUsecase(mockRepo)
		expectedVotes := &[]dto.GetPlaceWithTotalVotes{
			{
				PlaceId:   uuid.New(),
				PlaceName: "Place 1",
				TotalVote: 10,
			},
			{
				PlaceId:   uuid.New(),
				PlaceName: "Place 2",
				TotalVote: 15,
			},
		}
		expectedTotalVoters := int64(20)
		expectedTotalVotes := int64(25)
		mockRepo.On("GetTotalVotes").Return(expectedVotes, nil)
		mockRepo.On("TotalVoters").Return(expectedTotalVoters, nil)
		mockRepo.On("TotalVotesReceived").Return(expectedTotalVotes, nil)

		votes, total, err := usecase.GetPlaceWithTotalVotes()
		assert.NoError(t, err)
		assert.NotNil(t, votes)
		assert.Equal(t, expectedVotes, votes)
		assert.Equal(t, []int64{expectedTotalVoters, expectedTotalVotes}, total)
	})

	t.Run("Test GetPlaceWithTotalVotes error", func(t *testing.T) {
		mockRepo := new(mocks.MockVoteRepository)
		usecase := usecases.NewVoteUsecase(mockRepo)
		mockRepo.On("GetTotalVotes").Return(nil, errors.New("error getting total votes"))

		votes, total, err := usecase.GetPlaceWithTotalVotes()
		assert.Error(t, err)
		assert.Nil(t, votes)
		assert.Nil(t, total)
		assert.EqualError(t, err, "error getting total votes")
	})

	t.Run("Test Create success", func(t *testing.T) {
		mockRepo := new(mocks.MockVoteRepository)
		usecase := usecases.NewVoteUsecase(mockRepo)
		userId := uuid.New()
		request := &dto.VoteRequest{
			PlaceId: uuid.New(),
		}
		mockRepo.On("FindUserById", userId).Return(nil, nil)
		mockRepo.On("Create", mock.Anything).Return(&entities.Vote{}, nil)

		newVote, err := usecase.Create(userId, request)
		assert.NoError(t, err)
		assert.NotNil(t, newVote)
	})

	t.Run("Test Create error - Already voted", func(t *testing.T) {
		mockRepo := new(mocks.MockVoteRepository)
		usecase := usecases.NewVoteUsecase(mockRepo)
		userId := uuid.New()
		request := &dto.VoteRequest{
			PlaceId: uuid.New(),
		}
		mockRepo.On("FindUserById", userId).Return(&entities.Vote{}, nil)

		newVote, err := usecase.Create(userId, request)
		assert.Error(t, err)
		assert.Nil(t, newVote)
		assert.EqualError(t, err, "You have already voted or cannot vote more than once.")
	})

	t.Run("Test Create error - Repository error", func(t *testing.T) {
		mockRepo := new(mocks.MockVoteRepository)
		usecase := usecases.NewVoteUsecase(mockRepo)
		userId := uuid.New()
		request := &dto.VoteRequest{
			PlaceId: uuid.New(),
		}
		mockRepo.On("FindUserById", userId).Return(nil, errors.New("repository error"))
		mockRepo.On("Create", mock.Anything).Return(nil, errors.New("repository error"))

		newVote, err := usecase.Create(userId, request)
		assert.Error(t, err)
		assert.Nil(t, newVote)
		assert.EqualError(t, err, "repository error")
	})
}
