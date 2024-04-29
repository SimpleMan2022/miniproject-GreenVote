package dto

type UserRequest struct {
	Email    string `json:"email" validate:"required"`
	Fullname string `json:"fullname"`
	Password string `json:"password"`
}
