package usecases

import (
	"evoting/dto"
	"evoting/entities"
	"evoting/errorHandlers"
	"evoting/helpers"
	"evoting/repositories"
	"github.com/google/uuid"
	"strings"
)

type PlaceUsecase interface {
	FindAll(page, limit int, sortBy, sortType string) (*[]entities.Place, *int64, error)
	FindById(id uuid.UUID) (*entities.Place, error)
	Create(request *dto.PlaceRequest) (*entities.Place, error)
	CreateAddress(place *entities.PlaceAddress) (*entities.PlaceAddress, error)
	Update(id uuid.UUID, request *dto.PlaceRequest) (*entities.Place, error)
	UpdateAddress(place *entities.PlaceAddress) (*entities.PlaceAddress, error)
	Delete(id uuid.UUID) error
}

type placeUsecase struct {
	repository repositories.PlaceRepository
}

func NewPlaceUsecase(uc repositories.PlaceRepository) *placeUsecase {
	return &placeUsecase{uc}
}

func (uc *placeUsecase) FindAll(page, limit int, sortBy, sortType string) (*[]entities.Place, *int64, error) {
	places, total, err := uc.repository.FindAll(page, limit, sortBy, sortType)
	if err != nil {
		return nil, nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	return places, total, nil
}

func (uc *placeUsecase) FindById(id uuid.UUID) (*entities.Place, error) {
	place, err := uc.repository.FindById(id)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	return place, nil
}

func (uc *placeUsecase) Create(request *dto.PlaceRequest) (*entities.Place, error) {
	exist, _ := uc.repository.FindByName(request.Name)

	if exist != nil && strings.ToLower(exist.Name) == strings.ToLower(request.Name) {
		return nil, &errorHandlers.BadRequestError{"Place is already created, please enter another place"}
	}
	location := helpers.GenerateImageLocation(request)
	place := &entities.Place{
		Id:          uuid.New(),
		Name:        request.Name,
		Description: request.Description,
		Longitude:   request.Longitude,
		Latitude:    request.Latitude,
		MapImage:    &location,
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
	}

	newAddress, err := uc.repository.CreateAddress(address)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{err.Error()}
	}
	return newAddress, nil
}

func (uc *placeUsecase) Update(id uuid.UUID, request *dto.PlaceRequest) (*entities.Place, error) {
	place, err := uc.repository.FindById(id)
	if err != nil {
		return nil, &errorHandlers.BadRequestError{err.Error()}
	}

	if strings.ToLower(place.Name) != strings.ToLower(request.Name) {
		exist, _ := uc.repository.FindByName(request.Name)

		if exist != nil && strings.ToLower(exist.Name) == strings.ToLower(request.Name) {
			return nil, &errorHandlers.BadRequestError{"Place is already created, please enter another place"}
		}
	}
	var location string
	if place.MapImage != nil {
		if err := helpers.DeleteImage("public/images/places", place.MapImage); err != nil {
			return nil, &errorHandlers.InternalServerError{err.Error()}
		}
		location = helpers.GenerateImageLocation(request)
	}
	place.Name = request.Name
	place.Description = request.Description
	place.Longitude = request.Longitude
	place.Latitude = request.Latitude
	place.MapImage = &location

	updatePlace, err := uc.repository.Update(place)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{err.Error()}
	}
	return updatePlace, nil
}

func (uc *placeUsecase) Delete(id uuid.UUID) error {
	place, err := uc.repository.FindById(id)
	if err != nil {
		return &errorHandlers.BadRequestError{err.Error()}
	}
	if place.MapImage != nil {
		if err := helpers.DeleteImage("public/images/places", place.MapImage); err != nil {
			return &errorHandlers.InternalServerError{err.Error()}
		}
	}
	if err := uc.repository.Delete(place); err != nil {
		return &errorHandlers.InternalServerError{err.Error()}
	}

	return nil
}

func (uc *placeUsecase) UpdateAddress(place *entities.PlaceAddress) (*entities.PlaceAddress, error) {
	address, err := uc.repository.FindAddress(place.PlaceId)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{err.Error()}
	}

	address.Province = place.Province
	address.City = place.City
	address.SubDistrict = place.SubDistrict
	address.StreetName = place.StreetName

	updateAddress, err := uc.repository.UpdateAddress(address)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{err.Error()}
	}
	return updateAddress, nil
}
