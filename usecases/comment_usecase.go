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
	GetAllCommentInPlace(placeId uuid.UUID) (*[]dto.CommentData, *dto.CommentDetail, error)
	Delete(commentId, userId, placeId uuid.UUID) error
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
func (uc *commentUsecase) GetAllCommentInPlace(placeId uuid.UUID) (*[]dto.CommentData, *dto.CommentDetail, error) {
	comments, err := uc.repository.FindByPlaceId(placeId)
	if err != nil {
		return nil, nil, &errorHandlers.BadRequestError{Message: err.Error()}
	}
	placeDetail, err := uc.repository.GetDetailPlace(placeId)
	if err != nil {
		return nil, nil, &errorHandlers.BadRequestError{Message: err.Error()}
	}

	return comments, placeDetail, nil
}

func (uc *commentUsecase) Delete(commentId, userId, placeId uuid.UUID) error {
	comment, err := uc.repository.FindById(commentId)
	if err != nil {
		return &errorHandlers.BadRequestError{Message: err.Error()}
	}
	if comment.PlaceId != placeId {
		return &errorHandlers.NotFoundError{Message: "comment does not belong to this post"}
	}

	if comment.UserId != userId {
		return &errorHandlers.UnAuthorizedError{Message: "You are not allowed to delete this comment"}
	}
	if err := uc.repository.Delete(comment); err != nil {
		return &errorHandlers.InternalServerError{Message: err.Error()}
	}
	return nil
}
