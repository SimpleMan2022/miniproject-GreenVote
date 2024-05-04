package repositories

import (
	"evoting/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserAddressRepository interface {
	FindAll(page, limit int, sortBy, sortType string) (*[]entities.UserAddress, *int64, error)
	GetDetailUser(id uuid.UUID) (*entities.User, error)
	Create(address *entities.UserAddress) (*entities.UserAddress, error)
	FindByUserId(id uuid.UUID) (*entities.UserAddress, error)
	Update(address *entities.UserAddress) (*entities.UserAddress, error)
	Delete(address *entities.UserAddress) error
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

func (r *userAddress) GetDetailUser(id uuid.UUID) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("id = ?", id).Preload("Address").First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
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

func (r *userAddress) Update(address *entities.UserAddress) (*entities.UserAddress, error) {
	if err := r.db.Save(&address).Error; err != nil {
		return nil, err
	}
	return address, nil
}

func (r *userAddress) Delete(address *entities.UserAddress) error {
	if err := r.db.Unscoped().Delete(&address).Error; err != nil {
		return err
	}
	return nil
}
