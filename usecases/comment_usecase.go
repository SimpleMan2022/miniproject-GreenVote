package usecases

import (
	"evoting/dto"
	"evoting/entities"
	"evoting/errorHandlers"
	"evoting/repositories"
	"github.com/google/uuid"
)

type CommentUsecase interface {
	Create(userId uuid.UUID, placeId uuid.UUID, request *dto.CommentRequest) (*entities.Comment, error)
}

type commentUsecase struct {
	repository repositories.CommentRepository
}

func NewCommentUsecase(repository repositories.CommentRepository) *commentUsecase {
	return &commentUsecase{repository}
}

func (uc *commentUsecase) Create(userId uuid.UUID, placeId uuid.UUID, request *dto.CommentRequest) (*entities.Comment, error) {
	response := &entities.Comment{
		Id:      uuid.New(),
		UserId:  userId,
		PlaceId: placeId,
		Body:    request.Body,
	}

	newComment, err := uc.repository.Create(response)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{err.Error()}
	}
	return newComment, nil
}
