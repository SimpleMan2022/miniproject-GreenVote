package dto

import (
	"evoting/entities"
	"github.com/google/uuid"
)

type UserAddress struct {
	Province    string `json:"province"`
	City        string `json:"city"`
	SubDistrict string `json:"sub_district"`
	StreetName  string `json:"street_name"`
	ZipCode     string `json:"zip_code"`
}

type CreateUserAddressRequest struct {
	UserId      uuid.UUID `json:"user_id"`
	Province    string    `validate:"required" json:"province"`
	City        string    `validate:"required" json:"city"`
	SubDistrict string    `validate:"required" json:"sub_district"`
	StreetName  string    `validate:"required" json:"street_name"`
	ZipCode     string    `validate:"required,number" json:"zip_code"`
}

type CreateUserAddressResponse struct {
	Id          uuid.UUID `json:"id"`
	Province    string    `json:"province"`
	City        string    `json:"city"`
	SubDistrict string    `json:"sub_district"`
	StreetName  string    `json:"street_name"`
	ZipCode     string    `json:"zip_code"`
}

func ToUserAddressResponse(address *entities.UserAddress) *CreateUserAddressResponse {
	return &CreateUserAddressResponse{
		Id:          address.Id,
		Province:    address.Province,
		City:        address.City,
		SubDistrict: address.SubDistrict,
		StreetName:  address.StreetName,
		ZipCode:     address.ZipCode,
	}
}
