package dto

import "mime/multipart"

type CreateRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Fullname string `json:"fullname" validate:"required,min=8"`
	Password string `json:"password" validate:"required,min=8"`
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
	AccessToken  string
	RefreshToken string
}
