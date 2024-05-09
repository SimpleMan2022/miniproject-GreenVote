package dto

import (
	"evoting/entities"
	"github.com/google/uuid"
)

type VoteRequest struct {
	PlaceId uuid.UUID `json:"place_id" validate:"required"`
}

type VoteResponse struct {
	UserId  uuid.UUID `json:"user_id"`
	PlaceId uuid.UUID `json:"place_id"`
}

type GetPlaceWithTotalVotes struct {
	PlaceId    uuid.UUID `json:"place_id"`
	PlaceName  string    `json:"place_name"`
	TotalVote  int       `json:"total_vote"`
	Percentage float64   `json:"percentage"`
}

type Detail struct {
	TotalVoters        int64 `json:"total_voters"`
	TotalVotesReceived int64 `json:"total_votes_received"`
}

type VoteData struct {
	Votes      *[]GetPlaceWithTotalVotes `json:"votes"`
	DetailVote Detail                    `json:"detail_vote"`
}

func ToVoteResponse(vote *entities.Vote) *VoteResponse {
	return &VoteResponse{
		UserId:  vote.UserId,
		PlaceId: vote.PlaceId,
	}
}
