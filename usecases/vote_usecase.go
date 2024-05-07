package usecases

import (
	"evoting/dto"
	"evoting/entities"
	"evoting/errorHandlers"
	"evoting/repositories"
	"github.com/google/uuid"
)

type VoteUsecase interface {
	GetPlaceWithTotalVotes() (*[]dto.GetPlaceWithTotalVotes, []int64, error)
	Create(userId uuid.UUID, request *dto.VoteRequest) (*entities.Vote, error)
}

type voteUsecase struct {
	repository repositories.VoteRepository
}

func NewVoteUsecase(repository repositories.VoteRepository) *voteUsecase {
	return &voteUsecase{repository}
}

func (uc *voteUsecase) GetPlaceWithTotalVotes() (*[]dto.GetPlaceWithTotalVotes, []int64, error) {
	votes, err := uc.repository.GetTotalVotes()
	if err != nil {
		return nil, nil, &errorHandlers.InternalServerError{err.Error()}
	}
	totalVoters, err := uc.repository.TotalVoters()
	if err != nil {
		return nil, nil, &errorHandlers.InternalServerError{err.Error()}
	}
	totalVotes, err := uc.repository.TotalVotesReceived()
	if err != nil {
		return nil, nil, &errorHandlers.InternalServerError{err.Error()}
	}

	return votes, []int64{totalVoters, totalVotes}, nil
}

func (uc *voteUsecase) Create(userId uuid.UUID, request *dto.VoteRequest) (*entities.Vote, error) {
	user, _ := uc.repository.FindUserById(userId)
	if user != nil {
		return nil, &errorHandlers.BadRequestError{Message: "You have already voted or cannot vote more than once."}
	}
	vote := &entities.Vote{
		Id:      uuid.New(),
		UserId:  userId,
		PlaceId: request.PlaceId,
	}

	newVote, err := uc.repository.Create(vote)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{err.Error()}
	}
	return newVote, nil
}
