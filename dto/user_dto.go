package dto

import (
	"evoting/entities"
	"github.com/google/uuid"
	"mime/multipart"
	"time"
)

type User struct {
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}

type CreateRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Fullname string `json:"fullname" validate:"required,min=8"`
	Password string `json:"password" validate:"required,min=8"`
}

type CreateResponse struct {
	Id        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Fullname  string    `json:"fullname"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateRequest struct {
	Email    string                `form:"email" validate:"required,email"`
	Fullname string                `form:"fullname" validate:"required,min=8"`
	Password string                `form:"password" validate:"required,min=8"`
	Image    *multipart.FileHeader `form:"image,maxFileSize"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	Id           uuid.UUID `json:"id"`
	Fullname     string    `json:"fullname"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token,omitempty"`
}

type findAllResponse struct {
	Id       uuid.UUID   `json:"id"`
	Email    string      `json:"email"`
	Fullname string      `json:"fullname"`
	Address  UserAddress `json:"address"`
	Image    *string     `json:"image"`
}

type ToNewTokenResponse struct {
	Token string
}

func ToCreateResponse(user *entities.User) *CreateResponse {
	return &CreateResponse{
		Id:       user.Id,
		Email:    user.Email,
		Fullname: user.Fullname,
	}
}

func ToLoginResponse(user *LoginResponse) *LoginResponse {
	return &LoginResponse{
		Id:          user.Id,
		Fullname:    user.Fullname,
		AccessToken: user.AccessToken,
	}
}
func ToFindAllResponse(users *[]entities.User) *[]findAllResponse {
	responses := make([]findAllResponse, len(*users))
	for i, user := range *users {
		response := findAllResponse{
			Id:       user.Id,
			Email:    user.Email,
			Fullname: user.Fullname,
			Address: UserAddress{
				Province:    user.Address.Province,
				City:        user.Address.City,
				SubDistrict: user.Address.SubDistrict,
				StreetName:  user.Address.StreetName,
				ZipCode:     user.Address.ZipCode,
			},
			Image: user.Image,
		}
		responses[i] = response
	}
	return &responses
}

func ToByIdResponse(user *entities.User) *findAllResponse {
	return &findAllResponse{
		Id:       user.Id,
		Email:    user.Email,
		Fullname: user.Fullname,
		Address: UserAddress{
			Province:    user.Address.Province,
			City:        user.Address.City,
			SubDistrict: user.Address.SubDistrict,
			StreetName:  user.Address.StreetName,
			ZipCode:     user.Address.ZipCode,
		},
		Image: user.Image,
	}
}
