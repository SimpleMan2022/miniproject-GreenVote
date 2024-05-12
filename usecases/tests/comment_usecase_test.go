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

func TestCommentUsecase(t *testing.T) {
	// Create a new instance of MockCommentRepository

	// Mock data
	userId := uuid.New()
	placeId := uuid.New()
	commentRequest := &dto.CommentRequest{Body: "Test comment body"}
	comment := &entities.Comment{
		Id:      uuid.New(),
		UserId:  userId,
		PlaceId: placeId,
		Body:    commentRequest.Body,
	}

	// Test Create method
	t.Run("Test Create success", func(t *testing.T) {
		mockRepo := new(mocks.MockCommentRepository)
		usecase := usecases.NewCommentUsecase(mockRepo)
		mockRepo.On("Create", mock.Anything).Return(comment, nil)
		newComment, err := usecase.Create(userId, placeId, commentRequest)
		assert.NoError(t, err)
		assert.NotNil(t, newComment)
		assert.Equal(t, comment.Id, newComment.Id)
	})

	t.Run("Test Create error", func(t *testing.T) {
		mockRepo := new(mocks.MockCommentRepository)
		usecase := usecases.NewCommentUsecase(mockRepo)
		mockRepo.On("Create", mock.Anything).Return(nil, errors.New("create error"))
		newComment, err := usecase.Create(userId, placeId, commentRequest)
		assert.Error(t, err)
		assert.Nil(t, newComment)
	})

	t.Run("Test GetAllCommentInPlace success", func(t *testing.T) {
		mockRepo := new(mocks.MockCommentRepository)
		usecase := usecases.NewCommentUsecase(mockRepo)
		mockRepo.On("FindByPlaceId", placeId).Return(&[]dto.CommentData{}, nil)
		mockRepo.On("GetDetailPlace", placeId).Return(&dto.CommentDetail{}, nil)

		comments, placeDetail, err := usecase.GetAllCommentInPlace(placeId)
		assert.NoError(t, err)
		assert.NotNil(t, comments)
		assert.NotNil(t, placeDetail)
	})

	t.Run("Test GetAllCommentInPlace error", func(t *testing.T) {
		mockRepo := new(mocks.MockCommentRepository)
		usecase := usecases.NewCommentUsecase(mockRepo)
		mockRepo.On("FindByPlaceId", placeId).Return(nil, nil, errors.New("find error"))
		mockRepo.On("GetDetailPlace", placeId).Return(nil, errors.New("find error"))
		comments, placeDetail, err := usecase.GetAllCommentInPlace(placeId)
		assert.Error(t, err)
		assert.Nil(t, comments)
		assert.Nil(t, placeDetail)
	})

	// Test Update method
	t.Run("Test Update success", func(t *testing.T) {
		mockRepo := new(mocks.MockCommentRepository)
		usecase := usecases.NewCommentUsecase(mockRepo)
		mockRepo.On("FindById", comment.Id).Return(comment, nil)
		mockRepo.On("Update", comment).Return(comment, nil)
		updatedComment, err := usecase.Update(comment.Id, userId, placeId, commentRequest)
		assert.NoError(t, err)
		assert.NotNil(t, updatedComment)
		assert.Equal(t, comment.Id, updatedComment.Id)
	})

	t.Run("Test Update error", func(t *testing.T) {
		mockRepo := new(mocks.MockCommentRepository)
		usecase := usecases.NewCommentUsecase(mockRepo)
		mockRepo.On("FindById", comment.Id).Return(nil, errors.New("find error"))
		updatedComment, err := usecase.Update(comment.Id, userId, placeId, commentRequest)
		assert.Error(t, err)
		assert.Nil(t, updatedComment)
	})

	// Test Delete method
	t.Run("Test Delete success", func(t *testing.T) {
		mockRepo := new(mocks.MockCommentRepository)
		usecase := usecases.NewCommentUsecase(mockRepo)
		mockRepo.On("FindById", comment.Id).Return(comment, nil)
		mockRepo.On("Delete", comment).Return(nil)
		err := usecase.Delete(comment.Id, userId, placeId)
		assert.NoError(t, err)
	})

	t.Run("Test Delete error", func(t *testing.T) {
		mockRepo := new(mocks.MockCommentRepository)
		usecase := usecases.NewCommentUsecase(mockRepo)
		mockRepo.On("FindById", comment.Id).Return(nil, errors.New("find error"))
		err := usecase.Delete(comment.Id, userId, placeId)
		assert.Error(t, err)
	})
}
