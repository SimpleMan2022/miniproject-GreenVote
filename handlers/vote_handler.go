package handlers

import (
	"evoting/dto"
	"evoting/errorHandlers"
	"evoting/helpers"
	"evoting/usecases"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type voteHandler struct {
	usecase usecases.VoteUsecase
}

func NewVoteHandler(usecase usecases.VoteUsecase) *voteHandler {
	return &voteHandler{usecase}
}

func (h *voteHandler) GetPlaceWithTotalVotes(ctx echo.Context) error {
	votes, slice, err := h.usecase.GetPlaceWithTotalVotes()
	if err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}

	voteData := dto.VoteData{
		Votes: votes,
		DetailVote: dto.Detail{
			TotalVoters:        slice[0],
			TotalVotesReceived: slice[1],
		},
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "success",
		Data:       voteData,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *voteHandler) CreateVote(ctx echo.Context) error {
	var request dto.VoteRequest
	if err := ctx.Bind(&request); err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.InternalServerError{err.Error()})
	}
	if err := helpers.ValidateRequest(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to create vote. please ensure your input correctly",
			Data:       err,
		})
	}
	id := ctx.Get("userId")
	userId := id.(*uuid.UUID)
	newVote, err := h.usecase.Create(*userId, &request)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	voteResponse := dto.ToVoteResponse(newVote)
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Your vote has been recorded successfully. Thank you!",
		Data:       voteResponse,
	})
	return ctx.JSON(http.StatusCreated, response)
}
