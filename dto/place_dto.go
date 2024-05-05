package dto

import (
	"evoting/entities"
	"github.com/google/uuid"
	"time"
)

type PlaceRequest struct {
	Name        string  `json:"name" validate:"required,min=5"`
	Description string  `json:"description" validate:"required"`
	Longitude   float64 `json:"longitude" validate:"numeric"`
	Latitude    float64 `json:"latitude" validate:"numeric"`
	MapImage    *string `json:"map_image"`
}

type PlaceResponse struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Longitude   float64      `json:"longitude"`
	Latitude    float64      `json:"latitude"`
	MapImage    *string      `json:"map_image"`
	Address     PlaceAddress `json:"address"`
}

type findAllPlacesResponse struct {
	Id          uuid.UUID    `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Longitude   float64      `json:"longitude"`
	Latitude    float64      `json:"latitude"`
	MapImage    *string      `json:"map_image"`
	Address     PlaceAddress `json:"address"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

func ToPlaceResponse(place *entities.Place, address *entities.PlaceAddress) *PlaceResponse {
	return &PlaceResponse{
		Name:        place.Name,
		Description: place.Description,
		Longitude:   place.Longitude,
		Latitude:    place.Latitude,
		MapImage:    place.MapImage,
		Address: PlaceAddress{
			Province:    address.Province,
			City:        address.City,
			SubDistrict: address.SubDistrict,
			StreetName:  address.StreetName,
			ZipCode:     address.ZipCode,
		},
	}
}

func ToPlaceByIdResponse(place *entities.Place) *PlaceResponse {
	return &PlaceResponse{
		Name:        place.Name,
		Description: place.Description,
		Longitude:   place.Longitude,
		Latitude:    place.Latitude,
		MapImage:    place.MapImage,
		Address: PlaceAddress{
			Province:    place.Address.Province,
			City:        place.Address.City,
			SubDistrict: place.Address.SubDistrict,
			StreetName:  place.Address.StreetName,
			ZipCode:     place.Address.ZipCode,
		},
	}
}

func ToFindAllPlacesResponse(places *[]entities.Place) *[]findAllPlacesResponse {
	responses := make([]findAllPlacesResponse, len(*places))
	for i, place := range *places {
		response := findAllPlacesResponse{
			Id:          place.Id,
			Name:        place.Name,
			Description: place.Description,
			Longitude:   place.Longitude,
			Latitude:    place.Latitude,
			MapImage:    place.MapImage,
			Address: PlaceAddress{
				Province:    place.Address.Province,
				City:        place.Address.City,
				SubDistrict: place.Address.SubDistrict,
				StreetName:  place.Address.StreetName,
				ZipCode:     place.Address.ZipCode,
			},
			CreatedAt: place.CreatedAt,
			UpdatedAt: place.UpdatedAt,
		}
		responses[i] = response
	}
	return &responses
}
