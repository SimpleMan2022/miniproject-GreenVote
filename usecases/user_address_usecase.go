package usecases

import (
	"evoting/entities"
	"evoting/errorHandlers"
	"evoting/repositories"
	"github.com/google/uuid"
)

type UserAddress interface {
	FindAll(page, limit int, sortBy, sortType string) (*[]entities.UserAddress, *int64, error)
	FindById(id uuid.UUID) (*entities.UserAddress, error)
	Create(address *entities.UserAddress) (*entities.UserAddress, error)
	FindByUserId(id uuid.UUID) (*entities.UserAddress, error)
	Update(user *entities.UserAddress) (*entities.UserAddress, error)
	Delete(user *entities.UserAddress) error
}

type userAddress struct {
	repository repositories.UserAddressRepository
}

func NewUserAddressUsecase(address repositories.UserAddressRepository) *userAddress {
	return &userAddress{address}
}

func (uc *userAddress) FindAll(page, limit int, sortBy, sortType string) (*[]entities.UserAddress, *int64, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *userAddress) FindById(id uuid.UUID) (*entities.UserAddress, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *userAddress) Create(address *entities.UserAddress) (*entities.UserAddress, error) {
	existingUser, _ := uc.repository.FindByUserId(address.UserId)
	if existingUser != nil {
		return nil, &errorHandlers.BadRequestError{Message: "Failed to create address, user only have one address"}
	}

	newAddress, err := uc.repository.Create(address)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{err.Error()}
	}
	return newAddress, nil
}

func (uc *userAddress) FindByUserId(id uuid.UUID) (*entities.UserAddress, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *userAddress) Update(user *entities.UserAddress) (*entities.UserAddress, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *userAddress) Delete(user *entities.UserAddress) error {
	//TODO implement me
	panic("implement me")
}
