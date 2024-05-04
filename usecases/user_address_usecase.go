package usecases

import (
	"evoting/dto"
	"evoting/entities"
	"evoting/errorHandlers"
	"evoting/repositories"
	"github.com/google/uuid"
)

type UserAddressUsecase interface {
	FindAll(page, limit int, sortBy, sortType string) (*[]entities.UserAddress, *int64, error)
	GetDetailUser(id uuid.UUID) (*entities.User, error)
	Create(address *dto.CreateUserAddressRequest) (*entities.UserAddress, error)
	FindByUserId(id uuid.UUID) (*entities.UserAddress, error)
	Update(address *dto.CreateUserAddressRequest) (*entities.UserAddress, error)
	Delete(id uuid.UUID) error
}

type userAddress struct {
	repositoryAddress repositories.UserAddressRepository
}

func NewUserAddressUsecase(address repositories.UserAddressRepository) *userAddress {
	return &userAddress{address}
}

func (uc *userAddress) FindAll(page, limit int, sortBy, sortType string) (*[]entities.UserAddress, *int64, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *userAddress) GetDetailUser(id uuid.UUID) (*entities.User, error) {
	user, err := uc.repositoryAddress.GetDetailUser(id)

	if err != nil {
		return nil, &errorHandlers.BadRequestError{Message: err.Error()}
	}
	return user, nil
}

func (uc *userAddress) Create(req *dto.CreateUserAddressRequest) (*entities.UserAddress, error) {
	existingUser, _ := uc.repositoryAddress.FindByUserId(req.UserId)
	if existingUser != nil {
		return nil, &errorHandlers.BadRequestError{Message: "Failed to create address, user only have one address"}
	}

	address := &entities.UserAddress{
		Id:          uuid.New(),
		UserId:      req.UserId,
		Province:    req.Province,
		City:        req.City,
		SubDistrict: req.SubDistrict,
		StreetName:  req.StreetName,
		ZipCode:     req.ZipCode,
	}

	newAddress, err := uc.repositoryAddress.Create(address)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{err.Error()}
	}
	return newAddress, nil
}

func (uc *userAddress) FindByUserId(id uuid.UUID) (*entities.UserAddress, error) {
	user, err := uc.repositoryAddress.FindByUserId(id)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{Message: err.Error()}
	}
	return user, nil
}

func (uc *userAddress) Update(address *dto.CreateUserAddressRequest) (*entities.UserAddress, error) {
	useradd, err := uc.repositoryAddress.FindByUserId(address.UserId)
	if err != nil {
		return nil, &errorHandlers.BadRequestError{err.Error()}
	}
	useradd.Province = address.Province
	useradd.City = address.City
	useradd.SubDistrict = address.SubDistrict
	useradd.StreetName = address.StreetName
	useradd.ZipCode = address.ZipCode

	updateUser, err := uc.repositoryAddress.Update(useradd)
	if err != nil {
		return nil, &errorHandlers.InternalServerError{err.Error()}
	}

	return updateUser, nil
}

func (uc *userAddress) Delete(id uuid.UUID) error {
	userAdd, err := uc.repositoryAddress.FindByUserId(id)
	if err != nil {
		return &errorHandlers.BadRequestError{err.Error()}
	}
	if err := uc.repositoryAddress.Delete(userAdd); err != nil {
		return &errorHandlers.InternalServerError{err.Error()}
	}
	return nil
}
