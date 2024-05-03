package repositories

import (
	"evoting/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserAddressRepository interface {
	FindAll(page, limit int, sortBy, sortType string) (*[]entities.UserAddress, *int64, error)
	FindById(id uuid.UUID) (*entities.UserAddress, error)
	Create(address *entities.UserAddress) (*entities.UserAddress, error)
	FindByUserId(id uuid.UUID) (*entities.UserAddress, error)
	Update(user *entities.UserAddress) (*entities.UserAddress, error)
	Delete(user *entities.UserAddress) error
}

type userAddress struct {
	db *gorm.DB
}

func NewUserAddressRepository(db *gorm.DB) *userAddress {
	return &userAddress{db}
}

func (r *userAddress) FindAll(page, limit int, sortBy, sortType string) (*[]entities.UserAddress, *int64, error) {
	//TODO implement me
	panic("implement me")
}

func (r *userAddress) FindById(id uuid.UUID) (*entities.UserAddress, error) {
	//TODO implement me
	panic("implement me")
}

func (r *userAddress) Create(address *entities.UserAddress) (*entities.UserAddress, error) {
	if err := r.db.Create(&address).Error; err != nil {
		return nil, err
	}
	return address, nil
}

func (r *userAddress) FindByUserId(id uuid.UUID) (*entities.UserAddress, error) {
	var address *entities.UserAddress
	if err := r.db.Where("user_id = ?", id).First(&address).Error; err != nil {
		return nil, err
	}
	return address, nil
}

func (r *userAddress) Update(user *entities.UserAddress) (*entities.UserAddress, error) {
	//TODO implement me
	panic("implement me")
}

func (r *userAddress) Delete(user *entities.UserAddress) error {
	//TODO implement me
	panic("implement me")
}
