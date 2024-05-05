package dto

import (
	"evoting/entities"
	"github.com/google/uuid"
)

type PlaceAddress struct {
	Id          uuid.UUID `json:"id"`
	Province    string    `json:"province"`
	City        string    `json:"city"`
	SubDistrict string    `json:"sub_district"`
	StreetName  string    `json:"street_name"`
	ZipCode     string    `json:"zip_code"`
}

type PlaceAddressRequest struct {
	PlaceId     uuid.UUID `json:"place_id"`
	Province    string    `json:"province" validate:"required,min=5"`
	City        string    `json:"city" validate:"required"`
	SubDistrict string    `json:"sub_district" validate:"required"`
	StreetName  string    `json:"street_name" validate:"required"`
	ZipCode     string    `json:"zip_code" validate:"required"`
}

type PlaceAddressResponse struct {
	Id          uuid.UUID `json:"id"`
	Province    string    `json:"province"`
	City        string    `json:"city"`
	SubDistrict string    `json:"sub_district"`
	StreetName  string    `json:"street_name"`
	ZipCode     string    `json:"zip_code"`
}

func ToPlaceAddressResponse(address *entities.PlaceAddress) *PlaceAddressResponse {
	return &PlaceAddressResponse{
		Id:          address.Id,
		Province:    address.Province,
		City:        address.City,
		SubDistrict: address.SubDistrict,
		StreetName:  address.StreetName,
		ZipCode:     address.ZipCode,
	}
}
