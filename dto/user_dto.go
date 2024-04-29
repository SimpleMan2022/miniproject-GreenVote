package dto

type UserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Fullname string `json:"fullname" validate:"required,min=8"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
}
