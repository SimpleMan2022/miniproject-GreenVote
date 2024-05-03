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
	Province    string    `json:"province"`
	City        string    `json:"city"`
	SubDistrict string    `json:"sub_district"`
	StreetName  string    `json:"street_name"`
	ZipCode     string    `json:"zip_code"`
}

type CreateUserAddressResponse struct {
	Id          uuid.UUID `json:"id"`
	User        User      `json:"user"`
	Province    string    `json:"province"`
	City        string    `json:"city"`
	SubDistrict string    `json:"sub_district"`
	StreetName  string    `json:"street_name"`
	ZipCode     string    `json:"zip_code"`
}

func ToCreateUserAddressResponse(address *entities.UserAddress) *CreateUserAddressResponse {
	return &CreateUserAddressResponse{
		Id: address.Id,
		User: User{
			Email:    ,
			Fullname: ,
		},
		Province:    address.Province,
		City:        address.City,
		SubDistrict: address.SubDistrict,
		StreetName:  address.StreetName,
		ZipCode:     address.ZipCode,
	}
}
