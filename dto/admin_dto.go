package dto

import "github.com/google/uuid"

type LoginAdminRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginAdminResponse struct {
	Id           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token,omitempty"`
}

func ToLoginAdminResponse(admin *LoginAdminResponse) *LoginAdminResponse {
	return &LoginAdminResponse{
		Id:          admin.Id,
		Username:    admin.Username,
		AccessToken: admin.AccessToken,
	}
}
