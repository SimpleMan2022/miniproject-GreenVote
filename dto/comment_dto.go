package dto

import (
	"evoting/entities"
	"github.com/google/uuid"
	"time"
)

type CommentRequest struct {
	Body string `json:"body" validate:"required"`
}

type CommentCreateResponse struct {
	Id      uuid.UUID `json:"id"`
	UserId  uuid.UUID `json:"user_id"`
	PlaceId uuid.UUID `json:"place_id"`
	Body    string    `json:"body"`
}

type CommentData struct {
	CommentId uuid.UUID `json:"comment_id"`
	UserId    uuid.UUID `json:"user_id"`
	Fullname  string    `json:"fullname"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentDetail struct {
	PlaceName   string `json:"place_name"`
	Province    string `json:"province"`
	City        string `json:"city"`
	SubDistrict string `json:"sub_district"`
	StreetName  string `json:"street_name"`
}

type CommentFindAllResponse struct {
	PlaceDetail CommentDetail  `json:"place_detail"`
	Comments    *[]CommentData `json:"comments"`
}

func ToCommentResponse(comment *entities.Comment) *CommentCreateResponse {
	return &CommentCreateResponse{
		Id:      comment.Id,
		UserId:  comment.UserId,
		PlaceId: comment.PlaceId,
		Body:    comment.Body,
	}
}
