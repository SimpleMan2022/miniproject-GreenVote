package dto

import (
	"evoting/entities"
	"github.com/google/uuid"
)

type CommentRequest struct {
	Body string `json:"body"`
}

type CommentCreateResponse struct {
	Id      uuid.UUID `json:"id"`
	UserId  uuid.UUID `json:"user_id"`
	PlaceId uuid.UUID `json:"place_id"`
	Body    string    `json:"body"`
}

func ToCommentCreateResponse(comment *entities.Comment) *CommentCreateResponse {
	return &CommentCreateResponse{
		Id:      comment.Id,
		UserId:  comment.UserId,
		PlaceId: comment.PlaceId,
		Body:    comment.Body,
	}
}
