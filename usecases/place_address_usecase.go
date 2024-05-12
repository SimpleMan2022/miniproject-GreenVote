package usecases

import (
	"evoting/dto"
	"evoting/entities"
	"evoting/errorHandlers"
	"evoting/repositories"
	"github.com/google/uuid"
)

type PlaceAddressUsecase interface {
	Create(placeId uuid.UUID, request *dto.PlaceAddressRequest) (*entities.PlaceAddress, error)
	Update(id, placeId uuid.UUID, request *dto.PlaceAddressRequest) (*entities.PlaceAddress, error)
	Delete(id, placeId uuid.UUID) error
}

type placeAddressUsecase struct {
	repository repositories.PlaceAddressRepository
}

func NewPlaceAddressUsecase(repository repositories.PlaceAddressRepository) *placeAddressUsecase {
	return &placeAddressUsecase{repository}
}

func (uc *placeAddressUsecase) Create(placeId uuid.UUID, request *dto.PlaceAddressRequest) (*entities.PlaceAddress, error) {
	address := &entities.PlaceAddress{
		Id:          uuid.New(),
		PlaceId:     placeId,
		Province:    request.Province,
		City:        request.City,
		SubDistrict: request.SubDistrict,
		StreetName:  request.StreetName,
		ZipCode:     request.ZipCode,
	}

	newAddress, err := uc.repository.Create(address)
	if err != nil {
		return nil, err
	}
	return newAddress, nil
}

func (uc *placeAddressUsecase) Update(id, placeId uuid.UUID, request *dto.PlaceAddressRequest) (*entities.PlaceAddress, error) {
	address, err := uc.repository.FindById(id)
	if err != nil {
		return nil, &errorHandlers.BadRequestError{err.Error()}
	}
	if address.PlaceId != placeId {
		return nil, &errorHandlers.BadRequestError{err.Error()}
	}
	address.Province = request.Province
	address.City = request.City
	address.StreetName = request.StreetName
	address.SubDistrict = request.SubDistrict
	address.ZipCode = request.ZipCode

	updateAddress, err := uc.repository.Update(address)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{err.Error()}
	}
	return updateAddress, nil
}

func (uc *placeAddressUsecase) Delete(id, placeId uuid.UUID) error {
	address, err := uc.repository.FindById(id)
	if address == nil || address.PlaceId != placeId {
		return &errorHandlers.BadRequestError{Message: err.Error()}
	}
	if err != nil {
		return &errorHandlers.BadRequestError{Message: err.Error()}
	}
	if err := uc.repository.Delete(address); err != nil {
		return &errorHandlers.InternalServerError{Message: err.Error()}
	}
	return nil
}
