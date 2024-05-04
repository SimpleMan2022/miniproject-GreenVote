package usecases

import (
	"evoting/dto"
	"evoting/entities"
	"evoting/errorHandlers"
	"evoting/repositories"
	"github.com/google/uuid"
)

type PlaceUsecase interface {
	Create(request *dto.PlaceRequest) (*entities.Place, error)
	CreateAddress(place *entities.PlaceAddress) (*entities.PlaceAddress, error)
}

type placeUsecase struct {
	repository repositories.PlaceRepository
}

func NewPlaceUsecase(uc repositories.PlaceRepository) *placeUsecase {
	return &placeUsecase{uc}
}

func (uc *placeUsecase) Create(request *dto.PlaceRequest) (*entities.Place, error) {
	place := &entities.Place{
		Id:          uuid.New(),
		Name:        request.Name,
		Description: request.Description,
		Longitude:   request.Longnitude,
		Latitude:    request.Latitude,
	}
	newPlace, err := uc.repository.Create(place)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{err.Error()}
	}
	return newPlace, nil
}

func (uc *placeUsecase) CreateAddress(place *entities.PlaceAddress) (*entities.PlaceAddress, error) {
	address := &entities.PlaceAddress{
		Id:          uuid.New(),
		PlaceId:     place.PlaceId,
		Province:    place.Province,
		City:        place.City,
		SubDistrict: place.SubDistrict,
		StreetName:  place.StreetName,
		ZipCode:     place.ZipCode,
	}

	newAddress, err := uc.repository.CreateAddress(address)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{err.Error()}
	}
	return newAddress, nil
}
