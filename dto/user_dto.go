package dto

type UserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Fullname string `json:"fullname" validate:"required,min=8"`
	Password string `json:"password" validate:"required,min=8"`
}
